package config

import (
	"os"

	"github.com/joho/godotenv"
)

type DBConfig struct {
	Host     string
	Database string
	Username string
	Password string
	Port     string
}

const PublicFolder = "public"

func GetPublicFolder() string {
	return PublicFolder
}
func GetEnvDatabase() *DBConfig {
	var dbConfig DBConfig

	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	dbConfig = DBConfig{
		Host:     os.Getenv("DB_HOST"),
		Username: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		Database: os.Getenv("DB_DATABASE"),
		Port:     "3306",
	}

	return &dbConfig
}

func GetUrl() string {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	return os.Getenv("URL")
}
