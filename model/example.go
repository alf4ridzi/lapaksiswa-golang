package model

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/alf4ridzi/lapaksiswa-golang/database"
)

type Example struct {
}

type ExampleModel struct {
	DB      *sql.DB
	table   string
	columns string
}

func NewExampleModel() *ExampleModel {
	// allowed columns
	var columnsAllowed = []string{}

	db, err := database.InitDatabase()
	if err != nil {
		panic(fmt.Sprintf("Kesalahan database : %v", err))
	}

	columns := strings.Join(columnsAllowed, ", ")

	return &ExampleModel{
		DB:      db,
		table:   "example",
		columns: columns,
	}
}

func (e *Example) GetExample() ([]Example, error) {
	return nil, nil
}
