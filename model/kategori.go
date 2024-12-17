package model

import (
	"database/sql"
	"fmt"

	"github.com/alf4ridzi/lapaksiswa-golang/database"
)

type Kategori struct {
	Nama      string
	Deskripsi string
	Slug      string
}

type KategoriModel struct {
	DB      *sql.DB
	table   string
	columns string
}

func NewKategoriModel() *KategoriModel {
	// allowed columns
	var columnsAllowed = []string{
		"kategori",
		"deskripsi",
		"slug",
	}

	columns := fmt.Sprintf("%s", columnsAllowed[0])
	for _, col := range columnsAllowed[1:] {
		columns = fmt.Sprintf("%s, %s", columns, col)
	}

	return &KategoriModel{
		DB:      database.InitDatabase(),
		table:   "kategori",
		columns: columns,
	}
}

func (k *KategoriModel) GetKategori() ([]Kategori, error) {

	query := fmt.Sprintf("SELECT %s FROM %s", k.columns, k.table)
	rows, err := k.DB.Query(query)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	defer rows.Close()

	var kategoriS []Kategori
	for rows.Next() {
		var kategori Kategori
		err = rows.Scan(
			&kategori.Nama,
			&kategori.Deskripsi,
			&kategori.Slug,
		)

		if err != nil {
			return nil, err
		}

		kategoriS = append(kategoriS, kategori)
	}

	return kategoriS, nil
}

func (p *ProdukModel) GetKategoriLike(kategori string) ([]Kategori, error) {
	query := fmt.Sprintf("SELECT %s FROM %s WHERE kategori LIKE ?", p.columns, p.table)
	rows, err := p.DB.Query(query, kategori)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var kategoriS []Kategori

	for rows.Next() {
		var kat Kategori

		if err = rows.Scan(
			&kat.Nama,
			&kat.Deskripsi,
			&kat.Slug,
		); err != nil {
			return nil, err
		}

		kategoriS = append(kategoriS, kat)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return kategoriS, nil
}
