package model

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/alf4ridzi/lapaksiswa-golang/database"
)

type Checkout struct {
}

type CheckoutModel struct {
	DB      *sql.DB
	table   string
	columns string
}

func NewCheckoutModel() *CheckoutModel {
	// allowed columns
	var columnsAllowed = []string{}

	db, err := database.InitDatabase()
	if err != nil {
		panic(fmt.Sprintf("Kesalahan database : %v", err))
	}

	columns := strings.Join(columnsAllowed, ", ")

	return &CheckoutModel{
		DB:      db,
		table:   "checkout",
		columns: columns,
	}
}

func (c *CheckoutModel) InsertCheckout(ProductID string, checkout string, total int64, username string) error {
	query := fmt.Sprintf("INSERT INTO %s (produk_id, checkout, total, username) VALUES (?, ?, ?, ?)", c.table)
	if _, err := c.DB.Exec(query, ProductID, checkout, total, username); err != nil {
		return err
	}

	return nil
}
