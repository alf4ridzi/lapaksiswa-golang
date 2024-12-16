package model

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/alf4ridzi/lapaksiswa-golang/database"
)

type Produk struct {
	Id          int64
	ProdukID    string
	Nama        string
	Username    string
	Toko        string
	Slug        string
	Terjual     int64
	Kategori    string
	Rating      float32
	Harga       int64
	Stok        int64
	Deskripsi   string
	Varian      string
	Diskon      float32
	Status      string
	Unit        string
	Foto        string
	Kondisi     string
	HargaFormat string
}

type ProdukModel struct {
	DB      *sql.DB
	table   string
	columns string
}

func NewProdukModel() *ProdukModel {
	// allowed columns
	var columnsAllowed = []string{
		"produk_id",
		"nama",
		"username",
		"nama_toko",
		"slug",
		"terjual",
		"kategori",
		"rating",
		"harga",
		"stok",
		"deskripsi",
		"varian",
		"diskon",
		"status",
		"unit",
		"foto",
		"kondisi",
	}

	columns := fmt.Sprintf("%s", columnsAllowed[0])
	for _, col := range columnsAllowed[1:] {
		columns = fmt.Sprintf("%s, %s", columns, col)
	}

	return &ProdukModel{
		DB:      database.InitDatabase(),
		table:   "produk",
		columns: columns,
	}
}

func FormatToIDR(price int64) string {
	idr := fmt.Sprintf("%d", price)
	n := len(idr)

	if n <= 3 {
		return idr
	}

	var result strings.Builder
	r := n % 3

	if r != 0 {
		result.WriteString(idr[:r])
		result.WriteString(".")
	}

	for i := r; i < n; i += 3 {
		result.WriteString(idr[i : i+3])
		if i+3 < n {
			result.WriteString(".")
		}
	}

	return result.String()
}

func (p *ProdukModel) GetRelated(produk *Produk) ([]Produk, error) {
	query := fmt.Sprintf("SELECT %s FROM %s WHERE kategori LIKE ? AND nama != ? AND nama LIKE ? AND stok > 0 AND status = 'tersedia' ORDER BY rating DESC, terjual DESC LIMIT 10", p.columns, p.table)
	rows, err := p.DB.Query(query, produk.Kategori, produk.Nama, produk.Nama)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	var related []Produk
	for rows.Next() {
		var prod Produk
		if err := rows.Scan(
			&prod.ProdukID,
			&prod.Nama,
			&prod.Username,
			&prod.Toko,
			&prod.Slug,
			&prod.Terjual,
			&prod.Kategori,
			&prod.Rating,
			&prod.Harga,
			&prod.Stok,
			&prod.Deskripsi,
			&prod.Varian,
			&prod.Diskon,
			&prod.Unit,
			&prod.Foto,
			&prod.Kondisi,
		); err != nil {
			return nil, err
		}

		related = append(related, prod)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return related, nil
}

func (p *ProdukModel) GetProduk(Slug string) (*Produk, error) {
	query := fmt.Sprintf("SELECT %s FROM %s WHERE slug = ?", p.columns, p.table)
	row := p.DB.QueryRow(query, Slug)

	var produk Produk

	err := row.Scan(
		&produk.ProdukID,
		&produk.Nama,
		&produk.Username,
		&produk.Toko,
		&produk.Slug,
		&produk.Terjual,
		&produk.Kategori,
		&produk.Rating,
		&produk.Harga,
		&produk.Stok,
		&produk.Deskripsi,
		&produk.Varian,
		&produk.Diskon,
		&produk.Status,
		&produk.Unit,
		&produk.Foto,
		&produk.Kondisi,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	produk.HargaFormat = FormatToIDR(produk.Harga)

	return &produk, nil
}

func (p *ProdukModel) GetTerlaris() ([]Produk, error) {
	const limit int = 8

	query := fmt.Sprintf("SELECT %s FROM %s ORDER BY terjual DESC LIMIT %d", p.columns, p.table, limit)
	rows, err := p.DB.Query(query)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	var produkS []Produk
	for rows.Next() {
		var produk Produk
		err = rows.Scan(
			&produk.ProdukID,
			&produk.Nama,
			&produk.Username,
			&produk.Toko,
			&produk.Slug,
			&produk.Terjual,
			&produk.Kategori,
			&produk.Rating,
			&produk.Harga,
			&produk.Stok,
			&produk.Deskripsi,
			&produk.Varian,
			&produk.Diskon,
			&produk.Status,
			&produk.Unit,
			&produk.Foto,
			&produk.Kondisi,
		)

		if err != nil {
			return nil, err
		}

		produk.HargaFormat = FormatToIDR(produk.Harga)
		produkS = append(produkS, produk)
	}

	return produkS, nil
}
