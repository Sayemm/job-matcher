package service

import (
	"fmt"
	"time"

	"github.com/Sayemm/job-matcher/go-loader/internal/domain"
)

type JobLoaderService struct {
	repo      JobRepository
	csvReader CSVReader
	batchSize int
}

func NewJobLoaderService(repo JobRepository, csvReader CSVReader, batchSize int) *JobLoaderService {
	return &JobLoaderService{
		repo:      repo,
		csvReader: csvReader,
		batchSize: batchSize,
	}
}

func (s *JobLoaderService) LoadJobs() error {
	fmt.Println("Starting to load...")
	startTime := time.Now()

	// Current Database Count
	initialCount, err := s.repo.Count()
	if err != nil {
		return fmt.Errorf("failed to get initial job count: %w", err)
	}
	fmt.Printf("Current jobs in database: %d\n", initialCount)

	totalProcessed := 0
	batchCount := 0

	// function for processign in Batch
	processBatch := func(jobs []*domain.Job) error {
		batchCount++
		if err := s.repo.SaveBatch(jobs); err != nil {
			return fmt.Errorf("failed to save batch %d: %w", batchCount, err)
		}
		totalProcessed += len(jobs)

		fmt.Printf("Processed batch %d: %d jobs (Total: %d)\n", batchCount, len(jobs), totalProcessed)
		return nil
	}

	// read csv and process in batches
	fmt.Println("Reading CSV File from service..")
	if err := s.csvReader.ReadInBatches(s.batchSize, processBatch); err != nil {
		return fmt.Errorf("failed to read CSV: %w", err)
	}

	// final Count after saving the batch and taken time
	finalCount, err := s.repo.Count()
	if err != nil {
		return fmt.Errorf("failed to get final job count: %w", err)
	}
	duration := time.Since(startTime)
	newJobs := finalCount - initialCount

	fmt.Printf("Job loading completed!\n")
	fmt.Printf("Dration: %v\n", duration.Round(time.Second))
	fmt.Printf("New jobs added: %d\n", newJobs)
	return nil
}
