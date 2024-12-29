package database

import (
	"database/sql"
	"fmt"

	"github.com/alf4ridzi/lapaksiswa-golang/config"
	_ "github.com/go-sql-driver/mysql"
)

type DBConfig struct {
	Host     string
	Database string
	Username string
	Password string
	Port     string
}

func InitDatabase() (*sql.DB, error) {
	database := config.GetEnvDatabase()

	dsn := fmt.Sprintf("%s@tcp(%s:%s)/%s", database.Username, database.Host, database.Port, database.Database)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	defer db.Close()

	return db, nil
}
