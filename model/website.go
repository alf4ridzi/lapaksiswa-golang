package model

import (
	"database/sql"
	"fmt"

	"github.com/alf4ridzi/lapaksiswa-golang/config"
	"github.com/alf4ridzi/lapaksiswa-golang/database"
)

type Website struct {
	Title       string
	Icon        string
	Logo        string
	Author      string
	Keywords    string
	Description string
	Canonical   string
	Pembayaran  []Pembayaran
	Kategori    []Kategori
	BaseUrl     string
}

type WebsiteModel struct {
	DB      *sql.DB
	table   string
	columns string
}

func NewWebsiteModel() *WebsiteModel {
	// allowed columns
	var columnsAllowed = []string{
		"web_title",
		"web_icon",
		"web_logo",
		"web_author",
		"web_keywords",
		"web_description",
	}

	columns := fmt.Sprintf("%s", columnsAllowed[0])
	for _, col := range columnsAllowed[1:] {
		columns = fmt.Sprintf("%s, %s", columns, col)
	}
	return &WebsiteModel{
		DB:      database.InitDatabase(),
		table:   "settings_web",
		columns: columns,
	}
}

func (w *WebsiteModel) GetSettings() (*Website, error) {
	query := fmt.Sprintf("SELECT %s FROM %s", w.columns, w.table)
	row := w.DB.QueryRow(query)
	var website Website

	err := row.Scan(&website.Title, &website.Icon, &website.Logo, &website.Author, &website.Keywords, &website.Description)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	pembayaranModel := NewPembayaranModel()
	pembayaran, err := pembayaranModel.GetPembayaran()

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	kategoriModel := NewKategoriModel()
	kategori, err := kategoriModel.GetKategori()

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	website.Kategori = kategori
	website.Pembayaran = pembayaran
	website.BaseUrl = config.GetUrl()
	website.Canonical = config.GetUrl()
	return &website, nil
}
