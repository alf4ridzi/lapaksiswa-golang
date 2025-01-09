package model

import (
	"database/sql"
	"fmt"
	"strconv"
	"strings"

	"github.com/alf4ridzi/lapaksiswa-golang/database"
)

type Transaksi struct {
	OrderID   string
	ProductID string
	Domain    string
	Username  string
	Alamat    string
	Qty       int64
	Status    string
	Harga     int64
	Metode    string
	Note      string
	NoHP      int64
}

type TransaksiModel struct {
	DB      *sql.DB
	table   string
	columns string
}

type TambahTransaksi struct {
	ProductID string `json:"productid"`
	Domain    string `json:"domain"`
	Username  string `json:"username"`
	Alamat    string `json:"alamat"`
	Qty       string `json:"qty"`
	Harga     string `json:"harga"`
	NoHP      string `json:"nohp"`
}

func NewTransaksiModel() *TransaksiModel {
	// allowed columns
	var columnsAllowed = []string{
		"order_id",
		"produk_id",
		"domain",
		"username",
		"alamat",
		"qty",
		"status",
		"harga",
		"metode",
		"note",
		"no_hp",
	}

	db, err := database.InitDatabase()
	if err != nil {
		panic(fmt.Sprintf("Kesalahan database : %v", err))
	}

	columns := strings.Join(columnsAllowed, ", ")

	return &TransaksiModel{
		DB:      db,
		table:   "transaksi",
		columns: columns,
	}
}

func (t *TransaksiModel) GetTotalProdukTerjual(Domain string) (int64, error) {
	query := fmt.Sprintf("SELECT SUM(id) FROM %s WHERE domain = ? AND status = 'sukses'", t.table)
	row := t.DB.QueryRow(query, Domain)

	var total int64
	var tmp sql.NullString

	if err := row.Scan(&tmp); err != nil {
		if err == sql.ErrNoRows {
			return 0, nil
		}

		return 0, err
	}

	if tmp.Valid {
		totalParse, err := strconv.ParseInt(tmp.String, 10, 64)
		if err != nil {
			return 0, err
		}

		total = totalParse
	} else {
		total = 0
	}

	return total, nil
}

func (t *TransaksiModel) GetAllTransaksi(Domain string) ([]Transaksi, error) {
	query := fmt.Sprintf("SELECT %s FROM %s WHERE domain = ?", t.columns, t.table)
	rows, err := t.DB.Query(query)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	defer rows.Close()

	var transaksiS []Transaksi
	for rows.Next() {
		var trx Transaksi
		if err := rows.Scan(
			&trx.OrderID,
			&trx.ProductID,
			&trx.Domain,
			&trx.Username,
			&trx.Alamat,
			&trx.Qty,
			&trx.Status,
			&trx.Harga,
			&trx.Metode,
			&trx.Note,
			&trx.NoHP); err != nil {
			return nil, err
		}

		transaksiS = append(transaksiS, trx)
	}

	if rows.Err() != nil {
		return nil, err
	}

	return transaksiS, nil
}
