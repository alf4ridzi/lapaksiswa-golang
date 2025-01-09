package model

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/alf4ridzi/lapaksiswa-golang/database"
)

type Cart struct {
	ProductID string
	Username  string
	CreatedAt string
}

type CartModel struct {
	DB      *sql.DB
	table   string
	columns string
}

func NewCartModel() *CartModel {
	var columnsAllowed = []string{
		"produk_id",
		"username",
	}

	db, err := database.InitDatabase()
	if err != nil {
		panic(fmt.Sprintf("Kesalahan database : %v", err))
	}

	columns := strings.Join(columnsAllowed, ", ")

	return &CartModel{
		DB:      db,
		table:   "cart",
		columns: columns,
	}
}

func (c *CartModel) AddToCart(Username string, ProductID string) error {
	// check if already in cart
	query := fmt.Sprintf("SELECT id FROM %s WHERE produk_id = ? AND username = ?", c.table)

	var id int

	row := c.DB.QueryRow(query, ProductID, Username)
	if err := row.Scan(&id); err != nil {
		if err != sql.ErrNoRows {
			return err
		}
	}

	if id > 0 {
		return fmt.Errorf("Barang sudah ada di keranjang!")
	}

	query = fmt.Sprintf("INSERT INTO %s (produk_id, username) VALUES (?, ?)", c.table)
	if _, err := c.DB.Exec(query, ProductID, Username); err != nil {
		return err
	}

	return nil
}
