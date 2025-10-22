package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type DBConfig struct {
	Host          string
	Port          int
	Name          string
	User          string
	Password      string
	EnableSSLMode bool
}

type Config struct {
	DB        *DBConfig
	CSVPath   string
	BatchSize int
}

var configurations *Config

func loadConfig() {
	err := godotenv.Load("../.env")
	if err != nil {
		fmt.Println(err)
		fmt.Println("Failed to load the Env Variables")
		os.Exit(1)
	}

	dbHost := os.Getenv("DB_HOST")
	if dbHost == "" {
		fmt.Println("Host is required")
		os.Exit(1)
	}

	dbPort := os.Getenv("DB_PORT")
	dbPrt, err := strconv.ParseInt(dbPort, 10, 64)
	if err != nil {
		fmt.Println("Port must be number")
		os.Exit(1)
	}

	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		fmt.Println("DB Name is required")
		os.Exit(1)
	}

	dbUser := os.Getenv("DB_USER")
	if dbUser == "" {
		fmt.Println("DB User is required")
		os.Exit(1)
	}

	dbPassword := os.Getenv("DB_PASSWORD")
	if dbPassword == "" {
		fmt.Println("DB Password is required")
		os.Exit(1)
	}

	dbSSLMode := os.Getenv("DB_ENABLE_SSL_MODE")
	enableSSLMode, err := strconv.ParseBool(dbSSLMode)
	if err != nil {
		fmt.Println("Please Provide Boolean Value")
		os.Exit(1)
	}

	dbConfig := &DBConfig{
		Host:          dbHost,
		Port:          int(dbPrt),
		Name:          dbName,
		User:          dbUser,
		Password:      dbPassword,
		EnableSSLMode: enableSSLMode,
	}

	csvPath := os.Getenv("CSV_PATH")
	if csvPath == "" {
		csvPath = "/data/postings.csv"
	}

	batchSize := 1000

	configurations = &Config{
		DB:        dbConfig,
		CSVPath:   csvPath,
		BatchSize: batchSize,
	}

}

func GetConfig() *Config {
	if configurations == nil {
		loadConfig()
	}
	return configurations
}
