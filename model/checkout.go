package model

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/alf4ridzi/lapaksiswa-golang/database"
)

type Checkout struct {
	ProductID  string
	CheckoutID string
	Total      string
	Username   string
	Qty        int64
}

type CheckoutModel struct {
	DB      *sql.DB
	table   string
	columns string
}

func NewCheckoutModel() *CheckoutModel {
	// allowed columns
	var columnsAllowed = []string{
		"produk_id",
		"checkout",
		"total",
		"username",
		"qty",
	}

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

func (c *CheckoutModel) IsValidCheckout(CheckoutID string, Username string) (bool, error) {
	query := fmt.Sprintf("SELECT id FROM %s WHERE checkout = ? AND username = ?", c.table)
	row := c.DB.QueryRow(query, CheckoutID, Username)

	var id int

	if err := row.Scan(&id); err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}

		return false, err
	}

	return id > 0, nil
}

func (c *CheckoutModel) InsertCheckout(ProductID string, checkout string, total int64, username string, Qty int64) error {
	query := fmt.Sprintf("INSERT INTO %s (produk_id, checkout, total, username, qty) VALUES (?, ?, ?, ?, ?)", c.table)
	if _, err := c.DB.Exec(query, ProductID, checkout, total, username, Qty); err != nil {
		return err
	}

	return nil
}

func (c *CheckoutModel) GetDetailCheckout(CheckoutID string, Username string) (*Checkout, error) {
	query := fmt.Sprintf("SELECT %s FROM %s WHERE checkout = ? AND username = ?", c.columns, c.table)
	row := c.DB.QueryRow(query, CheckoutID, Username)

	var detail Checkout
	var tmpT int64
	if err := row.Scan(&detail.ProductID,
		&detail.CheckoutID,
		&tmpT,
		&detail.Username,
		&detail.Qty); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	detail.Total = FormatToIDR(tmpT)
	return &detail, nil
}
