package kategori

import (
	"database/sql"
	"net/http"
	"path/filepath"
	"text/template"

	"github.com/alf4ridzi/lapaksiswa-golang/config/cookie"
	"github.com/alf4ridzi/lapaksiswa-golang/model"
	"github.com/gorilla/mux"
)

func Kategori(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		kategori := vars["kategori"]

		websiteModel := model.NewWebsiteModel()
		settings, err := websiteModel.GetSettings()
		if err != nil {
			w.Write([]byte(err.Error()))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		produkModel := model.NewProdukModel()
		Produk, err := produkModel.GetProdukByKategori(kategori)

		if err != nil {
			w.Write([]byte(err.Error()))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		data := make(map[string]any)
		data["Website"] = settings
		data["Produk"] = Produk
		data["Cookies"] = cookie.GetAllCookies(w, r)
		data["Kategori"] = kategori

		templates := []string{
			filepath.Join("views", "templates.html"),
			filepath.Join("views", "kategori", "kategori.html"),
		}

		tmpl, err := template.ParseFiles(templates...)
		if err != nil {
			w.Write([]byte(err.Error()))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		tl := []string{
			"header",
			"kategori",
			"footer",
		}

		for _, t := range tl {
			if err = tmpl.ExecuteTemplate(w, t, data); err != nil {
				w.Write([]byte(err.Error()))
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		}

	}
}
