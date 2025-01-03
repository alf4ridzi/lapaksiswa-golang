package model

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/alf4ridzi/lapaksiswa-golang/database"
)

type Tambah struct {
	Nama      string
	Deskripsi string
	Kategori  string
	Varian    string
	Unit      string
	Kondisi   string
	Stok      int64
	Harga     int64
}

type EditProduct struct {
	ProductID string
	Nama      string
	Deskripsi string
	Kategori  string
	Varian    string
	Unit      string
	Kondisi   string
	Stok      int64
	Harga     int64
}

type Produk struct {
	ProdukID    string
	Nama        string
	Domain      string
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
	Toko        string
	HargaFormat string
	CreatedAt   string
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
		"domain",
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

	db, err := database.InitDatabase()
	if err != nil {
		panic(fmt.Sprintf("Kesalahan database : %v", err))
	}

	const tableName string = "produk"

	for i, col := range columnsAllowed {
		columnsAllowed[i] = fmt.Sprintf("%s.%s", tableName, col)
	}

	columns := strings.Join(columnsAllowed, ", ")
	return &ProdukModel{
		DB:      db,
		table:   tableName,
		columns: columns,
	}
}

func (p *ProdukModel) DeleteProduct(Domain string, ProductID string) (bool, error) {
	Produk, err := p.GetProductByID(ProductID, Domain)
	if err != nil {
		return false, err
	}

	FotoProduct := Produk.Foto

	query := fmt.Sprintf("DELETE FROM %s WHERE domain = ? AND produk_id = ?", p.table)
	result, err := p.DB.Exec(query, Domain, ProductID)
	if err != nil {
		return false, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return false, err
	}

	return rowsAffected > 0, nil
}

func (p *ProdukModel) HandleEditProduct(Domain string, Product EditProduct) error {
	query := fmt.Sprintf(`UPDATE %s 
	SET nama = ?,
	deskripsi = ?,
	kategori = ?,
	varian = ?,
	unit = ?,
	kondisi = ?,
	stok = ?,
	harga = ?
	WHERE produk_id = ? AND domain = ?`, p.table)
	if _, err := p.DB.Exec(query,
		Product.Nama,
		Product.Deskripsi,
		Product.Kategori,
		Product.Varian,
		Product.Unit,
		Product.Kondisi,
		Product.Stok,
		Product.Harga,
		Product.ProductID,
		Domain); err != nil {
		return err
	}

	return nil
}

func (p *ProdukModel) GetProductByID(ProductID any, Domain string) (*Produk, error) {
	query := fmt.Sprintf("SELECT %s FROM %s WHERE produk_id = ? AND domain = ?", p.columns, p.table)
	row := p.DB.QueryRow(query, ProductID, Domain)

	var product Produk

	if err := row.Scan(
		&product.ProdukID,
		&product.Nama,
		&product.Domain,
		&product.Slug,
		&product.Terjual,
		&product.Kategori,
		&product.Rating,
		&product.Harga,
		&product.Stok,
		&product.Deskripsi,
		&product.Varian,
		&product.Diskon,
		&product.Status,
		&product.Unit,
		&product.Foto,
		&product.Kondisi,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	return &product, nil

}
func (p *ProdukModel) ChangeStatusProduct(ProdukID string, Domain string, Status string) (bool, error) {
	query := fmt.Sprintf("UPDATE %s SET status = ? WHERE produk_id = ? AND domain = ?", p.table)
	result, err := p.DB.Exec(query, Status, ProdukID, Domain)
	if err != nil {
		return false, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return false, err
	}

	return rowsAffected > 0, nil
}

func (p *ProdukModel) ValidasiProduk(ProdukID string, Domain string) (bool, error) {
	query := fmt.Sprintf("SELECT id FROM %s WHERE domain = ? AND produk_id = ?", p.table)
	row := p.DB.QueryRow(query, Domain, ProdukID)

	var id int64
	if err := row.Scan(&id); err != nil {
		return false, nil
	}

	return id > 0, nil
}

func (p *ProdukModel) TambahProduk(ProdukID string, Domain string, Slug string, Produk Tambah) error {
	query := fmt.Sprintf("INSERT INTO %s (produk_id, stok, nama, domain, slug, deskripsi, kategori, varian, unit, kondisi, harga) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", p.table)
	if _, err := p.DB.Exec(query, ProdukID, Produk.Stok, Produk.Nama, Domain, Slug, Produk.Deskripsi, Produk.Kategori, Produk.Varian, Produk.Unit, Produk.Kondisi, Produk.Harga); err != nil {
		return err
	}

	return nil
}

func (p *ProdukModel) IsProdukID(ProdukID string) (bool, error) {
	query := fmt.Sprintf("SELECT id FROM %s WHERE produk_id = ?", p.table)
	row := p.DB.QueryRow(query, ProdukID)

	var id int64

	if err := row.Scan(&id); err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

func (p *ProdukModel) UpdateFotoProduk(ProdukID string, Foto string) error {
	query := fmt.Sprintf("UPDATE %s SET foto = ? WHERE produk_id = ?", p.table)
	if _, err := p.DB.Exec(query, Foto, ProdukID); err != nil {
		return err
	}

	return nil
}

func (p *ProdukModel) GetProductToko(domain string) ([]Produk, error) {
	query := fmt.Sprintf("SELECT %s, created_at FROM %s WHERE domain = ?", p.columns, p.table)
	rows, err := p.DB.Query(query, domain)

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
		if err = rows.Scan(
			&prod.ProdukID,
			&prod.Nama,
			&prod.Domain,
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
			&prod.CreatedAt,
		); err != nil {
			return nil, err
		}

		produk = append(produk, prod)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return produk, nil
}

func (p *ProdukModel) Filter(Keyword string, KategoriFilter string, MinPrice int, MaxPrice int, KondisiFilter string, Urutan string) ([]Produk, error) {
	query := []string{fmt.Sprintf("SELECT %s, toko.nama FROM produk INNER JOIN toko ON produk.domain = toko.domain WHERE 1=1", p.columns)}
	args := []interface{}{}

	if Keyword != "" {
		query = append(query, "(produk.nama LIKE ? OR produk.kategori LIKE ?)")
		args = append(args, "%"+Keyword+"%", "%"+Keyword+"%")
	}

	if KategoriFilter != "" {
		query = append(query, "produk.kategori LIKE ?")
		args = append(args, "%"+KategoriFilter+"%")
	}

	if MinPrice > 0 && MaxPrice > MinPrice {
		query = append(query, "produk.harga BETWEEN ? AND ?")
		args = append(args, MinPrice, MaxPrice)
	}

	if KondisiFilter == "baru" || KondisiFilter == "bekas" {
		query = append(query, "produk.kondisi = ?")
		args = append(args, KondisiFilter)
	}

	orderBy := ""
	if Urutan == "terlama" {
		orderBy = "ORDER BY produk.created_at ASC"
	} else if Urutan == "terbaru" {
		orderBy = "ORDER BY produk.created_at DESC"
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
			&p.Domain,
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
			&p.Kondisi,
			&p.Toko); err != nil {
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
			&prod.Domain,
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
			&prod.Domain,
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
		&produk.Domain,
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

	query := fmt.Sprintf("SELECT %s, toko.nama FROM %s INNER JOIN toko ON produk.domain = toko.domain ORDER BY terjual DESC LIMIT %d", p.columns, p.table, limit)

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
			&produk.Domain,
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
			&produk.Toko,
		)

		if err != nil {
			return nil, err
		}

		produk.HargaFormat = FormatToIDR(produk.Harga)
		produkS = append(produkS, produk)
	}

	return produkS, nil
}
