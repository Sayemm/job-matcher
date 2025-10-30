package database

import (
	"fmt"

	"github.com/Sayemm/job-matcher/go-api/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func GetConnectionString(cnf *config.Config) string {
	connString := fmt.Sprintf(
		"user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
		cnf.DBUser, cnf.DBPassword, cnf.DBHost, cnf.DBPort, cnf.DBName,
	)
	return connString
}

func NewConnection(cnf *config.Config) (*sqlx.DB, error) {
	dbSource := GetConnectionString(cnf)
	dbCon, err := sqlx.Connect("postgres", dbSource)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return dbCon, nil
}
