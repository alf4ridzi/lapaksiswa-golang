package model

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/alf4ridzi/lapaksiswa-golang/database"
)

type Toko struct {
	Username  string
	Domain    string
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
		"domain",
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

	db, err := database.InitDatabase()
	if err != nil {
		panic(fmt.Sprintf("Kesalahan database : %v", err))
	}

	columns := strings.Join(columnsAllowed, ", ")

	return &TokoModel{
		DB:      db,
		table:   "toko",
		columns: columns,
	}
}

func (t *TokoModel) GetToko(domain string) (*Toko, error) {
	query := fmt.Sprintf("SELECT %s FROM %s WHERE domain = ?", t.columns, t.table)
	row := t.DB.QueryRow(query, domain)

	var toko Toko
	if err := row.Scan(
		&toko.Username,
		&toko.Domain,
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

func (t *TokoModel) GetSaldoToko(domain string) (string, error) {
	query := fmt.Sprintf("SELECT saldo FROM %s WHERE domain = ?", t.table)
	row := t.DB.QueryRow(query, domain)

	var saldo int64

	if err := row.Scan(&saldo); err != nil {
		if err == sql.ErrNoRows {
			return "", nil
		}

		return "", err
	}

	TotalSaldo := FormatToIDR(saldo)
	return TotalSaldo, nil
}

func (t *TokoModel) GetTokoByUsername(username string) (*Toko, error) {
	query := fmt.Sprintf("SELECT %s FROM %s WHERE username = ?", t.columns, t.table)
	row := t.DB.QueryRow(query, username)

	var toko Toko
	if err := row.Scan(
		&toko.Username,
		&toko.Domain,
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
