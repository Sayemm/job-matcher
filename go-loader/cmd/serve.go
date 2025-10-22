package cmd

import (
	"fmt"
	"os"

	"github.com/Sayemm/job-matcher/go-loader/config"
	"github.com/Sayemm/job-matcher/go-loader/internal/infrastructure/database"
)

func Serve() {
	cnf := config.GetConfig()

	// DATABASE CONNECTIONS & MIGRATIONS
	dbCon, err := database.NewConnection(cnf.DB)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = database.MigrateDB(dbCon, "./migrations")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
