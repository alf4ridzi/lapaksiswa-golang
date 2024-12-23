package model

import (
	"database/sql"
	"fmt"

	"github.com/alf4ridzi/lapaksiswa-golang/database"
)

type Order struct {
}

type OrderModel struct {
	DB      *sql.DB
	table   string
	columns string
}

func NewOrderModel() *OrderModel {
	// allowed columns
	var columnsAllowed = []string{}

	columns := fmt.Sprintf("%s", columnsAllowed[0])
	for _, col := range columnsAllowed[1:] {
		columns = fmt.Sprintf("%s, %s", columns, col)
	}

	return &OrderModel{
		DB:      database.InitDatabase(),
		table:   "order",
		columns: columns,
	}
}

func (o *OrderModel) Order() ([]Order, error) {
	return nil, nil
}
