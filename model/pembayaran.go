package model

import (
	"database/sql"
	"fmt"

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

	columns := fmt.Sprintf("%s", columnsAllowed[0])
	for _, col := range columnsAllowed[1:] {
		columns = fmt.Sprintf("%s, %s", columns, col)
	}

	return &PembayaranModel{
		DB:      database.InitDatabase(),
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
