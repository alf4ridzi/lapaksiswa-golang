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

type EditToko struct {
	Nama      string `json:"nama"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Deskripsi string `json:"deskripsi"`
	Alamat    string `json:"alamat"`
	Status    string `json:"status"`
	Kategori  string `json:"kategori"`
}

type Seller struct {
	Nama      string `json:"nama"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Alamat    string `json:"alamat"`
	Deskripsi string `json:"deskripsi"`
	NoHP      string `json:"nohp"`
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

func (t *TokoModel) UpdateToko(Domain string, Toko EditToko) error {
	query := fmt.Sprintf("UPDATE %s SET nama = ?, email = ?, no_hp = ?, deskripsi = ?, alamat = ?, status = ?, kategori = ? WHERE domain = ?", t.table)

	if _, err := t.DB.Exec(query, Toko.Nama, Toko.Email, Toko.Phone, Toko.Deskripsi, Toko.Alamat, Toko.Status, Toko.Kategori, Domain); err != nil {
		return err
	}

	return nil
}

func (t *TokoModel) GetToko(domain string) (*Toko, error) {
	query := fmt.Sprintf("SELECT %s FROM %s WHERE domain = ?", t.columns, t.table)
	row := t.DB.QueryRow(query, domain)

	var toko Toko
	var kategori sql.NullString

	if err := row.Scan(
		&toko.Username,
		&toko.Domain,
		&toko.Nama,
		&kategori,
		&toko.Logo,
		&toko.Deskripsi,
		&toko.Email,
		&toko.NoHP,
		&toko.Alamat,
		&toko.Rating,
		&toko.Status,
	); err != nil {
		if err == sql.ErrNoRows {
			return &Toko{}, nil
		}
		return nil, err
	}

	if kategori.Valid {
		toko.Kategori = kategori.String
	} else {
		toko.Kategori = ""
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
	var kategori sql.NullString

	if err := row.Scan(
		&toko.Username,
		&toko.Domain,
		&toko.Nama,
		&kategori,
		&toko.Logo,
		&toko.Deskripsi,
		&toko.Email,
		&toko.NoHP,
		&toko.Alamat,
		&toko.Rating,
		&toko.Status,
	); err != nil {
		if err == sql.ErrNoRows {
			return &Toko{}, nil
		}
		return nil, err
	}

	if kategori.Valid {
		toko.Kategori = kategori.String
	} else {
		toko.Kategori = ""
	}

	return &toko, nil
}

func (t *TokoModel) isExist(column string, value string) (bool, error) {
	query := fmt.Sprintf("SELECT 1 FROM %s WHERE %s = ? LIMIT 1", t.table, column)
	row := t.DB.QueryRow(query, value)

	var exist int
	if err := row.Scan(&exist); err != nil {
		return false, err
	}

	return exist == 1, nil
}

func (t *TokoModel) DaftarToko(Username string, seller Seller) error {
	// cek apakah username sudah ada
	if isExist, err := t.isExist("username", Username); err != nil {
		if err == sql.ErrNoRows {

		} else {
			return err
		}

	} else if isExist {
		return fmt.Errorf("username sudah memiliki toko")
	}

	// cek apakah domain sudah ada
	if isExist, err := t.isExist("domain", seller.Username); err != nil {
		if err == sql.ErrNoRows {

		} else {
			return err
		}
	} else if isExist {
		return fmt.Errorf("username toko sudah digunakan")
	}

	// cek apakah email sudah ada
	if isExist, err := t.isExist("email", seller.Email); err != nil {
		if err == sql.ErrNoRows {

		} else {
			return err
		}
	} else if isExist {
		return fmt.Errorf("email sudah digunakan")
	}

	query := fmt.Sprintf("INSERT INTO %s (domain, username, nama, deskripsi, email, no_hp, alamat) VALUES (?, ?, ?, ?, ?, ?, ?)", t.table)
	if _, err := t.DB.Exec(query, seller.Username, Username, seller.Nama, seller.Deskripsi, seller.Email, seller.NoHP, seller.Alamat); err != nil {
		return err
	}
	return nil
}
