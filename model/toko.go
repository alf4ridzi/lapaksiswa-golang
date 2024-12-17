package model

import (
	"database/sql"
	"fmt"

	"github.com/alf4ridzi/lapaksiswa-golang/database"
)

type Toko struct {
	Username  string
	Nama      string
	Kategori  string
	Logo      string
	Deskripsi string
	Email     string
	NoHP      int64
	Alamat    string
	Rating    int
	Status    string
}

type TokoModel struct {
	DB      *sql.DB
	table   string
	columns string
}

func NewTokoModel() *TokoModel {
	// allowed columns
	var columnsAllowed = []string{
		"username",
		"nama",
		"kategori",
		"logo",
		"deskripsi",
		"email",
		"no_hp",
		"alamat",
		"rating",
		"status",
	}

	columns := fmt.Sprintf("%s", columnsAllowed[0])
	for _, col := range columnsAllowed[1:] {
		columns = fmt.Sprintf("%s, %s", columns, col)
	}

	return &TokoModel{
		DB:      database.InitDatabase(),
		table:   "toko",
		columns: columns,
	}
}

func (t *TokoModel) GetToko(username string) (*Toko, error) {
	query := fmt.Sprintf("SELECT %s FROM %s WHERE username = ?", t.columns, t.table)
	row := t.DB.QueryRow(query, username)

	var toko Toko
	if err := row.Scan(
		&toko.Username,
		&toko.Nama,
		&toko.Kategori,
		&toko.Logo,
		&toko.Deskripsi,
		&toko.Email,
		&toko.NoHP,
		&toko.Alamat,
		&toko.Rating,
		&toko.Status,
	); err != nil {
		return nil, err
	}

	return &toko, nil
}
