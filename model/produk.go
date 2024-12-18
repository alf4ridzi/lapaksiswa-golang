package model

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/alf4ridzi/lapaksiswa-golang/database"
)

type Produk struct {
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

func (p *ProdukModel) Filter(Keyword string, KategoriFilter string, MinPrice int, MaxPrice int, KondisiFilter string, Urutan string) ([]Produk, error) {
	query := []string{fmt.Sprintf("SELECT %s FROM produk WHERE 1=1", p.columns)}
	args := []interface{}{}

	if Keyword != "" {
		query = append(query, "(nama LIKE ? OR kategori LIKE ?)")
		args = append(args, "%"+Keyword+"%", "%"+Keyword+"%")
	}

	if KategoriFilter != "" {
		query = append(query, "kategori LIKE ?")
		args = append(args, "%"+KategoriFilter+"%")
	}

	if MinPrice > 0 && MaxPrice > MinPrice {
		query = append(query, "harga BETWEEN ? AND ?")
		args = append(args, MinPrice, MaxPrice)
	}

	if KondisiFilter == "baru" || KondisiFilter == "bekas" {
		query = append(query, "kondisi = ?")
		args = append(args, KondisiFilter)
	}

	orderBy := ""
	if Urutan == "terlama" {
		orderBy = "ORDER BY created_at ASC"
	} else if Urutan == "terbaru" {
		orderBy = "ORDER BY created_at DESC"
	}

	finalQuery := strings.Join(query, " AND ") + " " + orderBy
	rows, err := p.DB.Query(finalQuery, args...)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var produk []Produk
	for rows.Next() {
		var p Produk
		if err := rows.Scan(
			&p.ProdukID,
			&p.Nama,
			&p.Username,
			&p.Slug,
			&p.Terjual,
			&p.Kategori,
			&p.Rating,
			&p.Harga,
			&p.Stok,
			&p.Deskripsi,
			&p.Varian,
			&p.Diskon,
			&p.Status,
			&p.Unit,
			&p.Foto,
			&p.Kondisi); err != nil {
			return nil, err
		}

		p.HargaFormat = FormatToIDR(p.Harga)
		produk = append(produk, p)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return produk, nil
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

func (p *ProdukModel) GetProdukByKategori(kategori string) ([]Produk, error) {
	query := fmt.Sprintf("SELECT %s FROM %s WHERE kategori LIKE ?", p.columns, p.table)
	rows, err := p.DB.Query(query, kategori)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	defer rows.Close()

	var produk []Produk
	for rows.Next() {
		var prod Produk
		if err := rows.Scan(
			&prod.ProdukID,
			&prod.Nama,
			&prod.Username,
			&prod.Slug,
			&prod.Terjual,
			&prod.Kategori,
			&prod.Rating,
			&prod.Harga,
			&prod.Stok,
			&prod.Deskripsi,
			&prod.Varian,
			&prod.Diskon,
			&prod.Status,
			&prod.Unit,
			&prod.Foto,
			&prod.Kondisi,
		); err != nil {
			return nil, err
		}

		prod.HargaFormat = FormatToIDR(prod.Harga)
		produk = append(produk, prod)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return produk, nil

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

	defer rows.Close()

	var related []Produk
	for rows.Next() {
		var prod Produk
		if err := rows.Scan(
			&prod.ProdukID,
			&prod.Nama,
			&prod.Username,
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

	defer rows.Close()

	var produkS []Produk
	for rows.Next() {
		var produk Produk
		err = rows.Scan(
			&produk.ProdukID,
			&produk.Nama,
			&produk.Username,
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
