package search

import (
	"database/sql"
	"net/http"
	"path/filepath"
	"strconv"
	"text/template"

	"github.com/alf4ridzi/lapaksiswa-golang/model"
)

func Search(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		Keyword := r.URL.Query().Get("keyword")
		KategoriFilter := r.URL.Query().Get("kategori")
		HargaMin := r.URL.Query().Get("harga_min")
		HargaMax := r.URL.Query().Get("harga_max")
		KondisiFilter := r.URL.Query().Get("kondisi")
		Urutan := r.URL.Query().Get("urutan")

		MinPrice, _ := strconv.Atoi(HargaMin)
		MaxPrice, _ := strconv.Atoi(HargaMax)

		WebsiteModel := model.NewWebsiteModel()
		settings, err := WebsiteModel.GetSettings()
		if err != nil {
			w.Write([]byte(err.Error()))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		ProdukModel := model.NewProdukModel()
		Produk, err := ProdukModel.Filter(Keyword, KategoriFilter, MinPrice, MaxPrice, KondisiFilter, Urutan)
		if err != nil {
			w.Write([]byte(err.Error()))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		data := make(map[string]any)
		data["Website"] = settings
		data["Produk"] = Produk
		data["Keyword"] = Keyword

		templates := []string{
			filepath.Join("views", "templates.html"),
			filepath.Join("views", "search", "search.html"),
		}

		tmpl, err := template.ParseFiles(templates...)
		if err != nil {
			w.Write([]byte(err.Error()))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		t := []string{
			"header",
			"search",
			"footer",
		}

		for _, tl := range t {
			if err = tmpl.ExecuteTemplate(w, tl, data); err != nil {
				w.Write([]byte(err.Error()))
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		}
	}
}
