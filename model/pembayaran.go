package model

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/alf4ridzi/lapaksiswa-golang/database"
)

type Pembayaran struct {
	Nama string
	Logo string
}

type PembayaranModel struct {
	DB      *sql.DB
	table   string
	columns string
}

func NewPembayaranModel() *PembayaranModel {
	var columnsAllowed = []string{
		"nama",
		"logo",
	}

	db, err := database.InitDatabase()
	if err != nil {
		panic(fmt.Sprintf("Kesalahan database : %v", err))
	}

	columns := strings.Join(columnsAllowed, ", ")

	return &PembayaranModel{
		DB:      db,
		table:   "pembayaran",
		columns: columns,
	}
}

func (p *PembayaranModel) GetPembayaran() ([]Pembayaran, error) {
	query := fmt.Sprintf("SELECT %s FROM %s", p.columns, p.table)

	rows, err := p.DB.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var pembayaranS []Pembayaran
	for rows.Next() {
		var pembayaran Pembayaran
		err = rows.Scan(&pembayaran.Nama, &pembayaran.Logo)
		if err != nil {
			return nil, err
		}
		pembayaranS = append(pembayaranS, pembayaran)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return pembayaranS, nil
}
