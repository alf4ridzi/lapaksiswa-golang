package model

import (
	"database/sql"
	"fmt"
	"strings"

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

	db, err := database.InitDatabase()
	if err != nil {
		panic(fmt.Sprintf("Kesalahan database : %v", err))
	}

	columns := strings.Join(columnsAllowed, ", ")

	return &KategoriModel{
		DB:      db,
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

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return kategoriS, nil
}

func (k *KategoriModel) GetKategoriLike(kategori string) ([]Kategori, error) {
	query := fmt.Sprintf("SELECT %s FROM %s WHERE kategori LIKE ?", k.columns, k.table)
	rows, err := k.DB.Query(query, kategori)

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

func (k *KategoriModel) GetProdukKategori(kategori string) ([]Kategori, error) {
	query := fmt.Sprintf("SELECT %s FROM %s WHERE kategori LIKE ?", k.columns, k.table)
	rows, err := k.DB.Query(query, kategori)

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
