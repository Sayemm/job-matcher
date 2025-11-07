package config

import "os"

type Config struct {
	DBHost       string
	DBPort       string
	DBName       string
	DBUser       string
	DBPassword   string
	ServerPort   string
	MLServiceURL string
}

var configurations *Config

func loadConfig() *Config {
	return &Config{
		DBHost:       getEnv("DB_HOST", "localhost"),
		DBPort:       getEnv("DB_PORT", "5432"),
		DBName:       getEnv("DB_NAME", "jobmatcher"),
		DBUser:       getEnv("DB_USER", "postgres"),
		DBPassword:   getEnv("DB_PASSWORD", ""),
		ServerPort:   getEnv("SERVER_PORT", "8080"),
		MLServiceURL: getEnv("ML_SERVICE_URL", "http://ml-matcher:5000"),
	}

}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func GetConfig() *Config {
	if configurations == nil {
		configurations = loadConfig()
	}
	return configurations
}
