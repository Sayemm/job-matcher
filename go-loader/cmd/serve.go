package cmd

import (
	"fmt"
	"os"

	"github.com/Sayemm/job-matcher/go-loader/config"
	"github.com/Sayemm/job-matcher/go-loader/internal/infrastructure/csv"
	"github.com/Sayemm/job-matcher/go-loader/internal/infrastructure/database"
	"github.com/Sayemm/job-matcher/go-loader/internal/infrastructure/repository"
	"github.com/Sayemm/job-matcher/go-loader/internal/service"
)

func Serve() {
	cnf := config.GetConfig()

	// DATABASE CONNECTIONS & MIGRATIONS
	dbCon, err := database.NewConnection(cnf.DB)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer dbCon.Close()

	err = database.MigrateDB(dbCon, "./migrations")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// REPO
	jobRepo := repository.NewJobRepository(dbCon)
	csvReader := csv.NewReader(cnf.CSVPath)

	// SERVICE
	service := service.NewJobLoaderService(
		jobRepo,
		csvReader,
		cnf.BatchSize,
	)
	service.LoadJobs()
}
