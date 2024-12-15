package model

import (
	"database/sql"
	"fmt"

	"github.com/alf4ridzi/lapaksiswa-golang/database"
)

type Example struct {
}

type ExampleModel struct {
	DB      *sql.DB
	table   string
	columns string
}

func NewExampleModel() *ProdukModel {
	// allowed columns
	var columnsAllowed = []string{}

	columns := fmt.Sprintf("%s", columnsAllowed[0])
	for _, col := range columnsAllowed[1:] {
		columns = fmt.Sprintf("%s, %s", columns, col)
	}

	return &ProdukModel{
		DB:      database.InitDatabase(),
		table:   "example",
		columns: columns,
	}
}

func (p *ProdukModel) GetExample() ([]Produk, error) {
	return nil, nil
}
