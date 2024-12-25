package model

import (
	"database/sql"
	"fmt"
	"strings"

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

	columns := strings.Join(columnsAllowed, ", ")

	db, err := database.InitDatabase()
	if err != nil {
		panic(fmt.Sprintf("Kesalahan database : %v", err))
	}

	return &WebsiteModel{
		DB:      db,
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
