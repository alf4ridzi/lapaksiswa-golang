package model

import (
	"database/sql"
	"fmt"
	"strconv"
	"strings"

	"github.com/alf4ridzi/lapaksiswa-golang/database"
)

type Transaksi struct {
}

type TransaksiModel struct {
	DB      *sql.DB
	table   string
	columns string
}

func NewTransaksiModel() *TransaksiModel {
	// allowed columns
	var columnsAllowed = []string{}

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
