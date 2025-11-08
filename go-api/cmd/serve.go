package cmd

import (
	"fmt"
	"os"

	"github.com/Sayemm/job-matcher/go-api/config"
	"github.com/Sayemm/job-matcher/go-api/internal/application/service/jobService"
	"github.com/Sayemm/job-matcher/go-api/internal/application/service/resumeService"
	"github.com/Sayemm/job-matcher/go-api/internal/infrastructure/database"
	"github.com/Sayemm/job-matcher/go-api/internal/infrastructure/http"
	"github.com/Sayemm/job-matcher/go-api/internal/infrastructure/http/handlers/jobHandler"
	"github.com/Sayemm/job-matcher/go-api/internal/infrastructure/http/handlers/resumeHandler"
)

func Serve() {
	fmt.Println("Job Matcher API Starting...")

	cnf := config.GetConfig()

	// DATABASE CONNECTIONS
	dbCon, err := database.NewConnection(cnf)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer dbCon.Close()
	fmt.Println("Database connected")

	// REPO
	jobRepo := database.NewJobRepo(dbCon)

	// SERVICE
	jobService := jobService.NewJobService(jobRepo)
	resumeService := resumeService.NewResumeService(jobRepo, cnf.MLServiceURL)

	// HANDLER
	jobHandler := jobHandler.NewJobHandler(jobService)
	resumehandler := resumeHandler.NewResumeHandler(resumeService)

	// START SERVER
	server := http.NewServer(jobHandler, resumehandler, cnf.ServerPort)

	// Start server
	fmt.Printf("Server starting on http://localhost%s\n", cnf.ServerPort)
	fmt.Println("Endpoints:")
	fmt.Println("   GET /api/jobs              - Get all jobs")
	fmt.Println("   GET /api/jobs/:id          - Get job by ID")
	fmt.Println("   GET /api/jobs/cluster/:id  - Get jobs by cluster")

	server.Start()

}
