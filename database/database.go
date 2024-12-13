package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

type DBConfig struct {
	Host     string
	Database string
	Username string
	Password string
	Port     string
}

var dbConfig DBConfig

func GetEnvDatabase() {
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
}

func InitDatabase() *sql.DB {
	GetEnvDatabase()
	dsn := fmt.Sprintf("%s@tcp(%s:%s)/%s", dbConfig.Username, dbConfig.Host, dbConfig.Port, dbConfig.Database)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return db
}
