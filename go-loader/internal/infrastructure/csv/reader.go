package csv

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/Sayemm/job-matcher/go-loader/internal/domain"
	"github.com/Sayemm/job-matcher/go-loader/internal/service"
)

type Reader interface {
	service.CSVReader
}

type reader struct {
	filePath string
}

func NewReader(filePath string) Reader {
	return &reader{
		filePath: filePath,
	}
}

func (r *reader) ReadInBatches(batchSize int, callback func([]*domain.Job) error) error {
	// Open The File
	fmt.Println("Hello from Reader...")
	file, err := os.Open(r.filePath)
	if err != nil {
		return fmt.Errorf("failed to open CSV file: %w", err)
	}
	defer file.Close()

	// CSV Reader and Reading the Header and mapping
	csvReader := csv.NewReader(file)

	header, err := csvReader.Read()
	if err != nil {
		return fmt.Errorf("failed to read CSV header: %w", err)
	}

	columnIndex := make(map[string]int)
	for i, col := range header {
		columnIndex[strings.ToLower(strings.TrimSpace(col))] = i
	}

	// batch processing
	batch := make([]*domain.Job, 0, batchSize)
	rowNum := 1

	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			if len(batch) > 0 {
				if err := callback(batch); err != nil {
					return fmt.Errorf("failed to process final batch: %w", err)
				}
			}
			break
		}

		if err != nil {
			fmt.Printf("Warning: skipping row %d due to error: %v\n", rowNum, err)
			rowNum++
			continue
		}

		// converting csv to job domain
		job := r.parseJob(record, columnIndex, rowNum)
		if job != nil {
			batch = append(batch, job)
		}

		if len(batch) >= batchSize {
			if err := callback(batch); err != nil {
				return fmt.Errorf("failed to process batch at row %d: %w", rowNum, err)
			}
			batch = batch[:0]
		}
		rowNum++
	}
	fmt.Println("End from reader")
	return nil
}

func (r *reader) parseJob(record []string, columnIndex map[string]int, rowNum int) *domain.Job {
	getColumn := func(name string) string {
		if idx, exists := columnIndex[name]; exists && idx < len(record) {
			return strings.TrimSpace(record[idx])
		}
		return ""
	}

	jobID := getColumn("job_id")
	if jobID == "" {
		fmt.Printf("Warning: skipping row %d - missing job_id\n", rowNum)
		return nil
	}

	job := &domain.Job{
		JobID:           jobID,
		CompanyName:     getColumn("company_name"),
		Title:           getColumn("title"),
		Description:     getColumn("description"),
		Location:        getColumn("location"),
		ExperienceLevel: getColumn("formatted_experience_level"),
	}
	remoteStr := getColumn("remote_allowed")
	job.RemoteAllowed = remoteStr == "1" || remoteStr == "true" || remoteStr == "True"

	if minSalStr := getColumn("min_salary"); minSalStr != "" {
		if val, err := strconv.ParseFloat(minSalStr, 64); err == nil {
			job.MinSalary = &val
		}
	}

	if maxSalStr := getColumn("max_salary"); maxSalStr != "" {
		if val, err := strconv.ParseFloat(maxSalStr, 64); err == nil {
			job.MaxSalary = &val
		}
	}

	return job
}
