package produk

import (
	"database/sql"
	"net/http"
	"path/filepath"
	"text/template"

	"github.com/alf4ridzi/lapaksiswa-golang/config/cookie"
	"github.com/alf4ridzi/lapaksiswa-golang/controller"
	"github.com/alf4ridzi/lapaksiswa-golang/model"
	"github.com/gorilla/mux"
)

func Produk(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		slug := vars["slug"]

		modelProduk := model.NewProdukModel()
		Produk, err := modelProduk.GetProduk(slug)
		if err != nil {
			w.Write([]byte(err.Error()))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if Produk == nil {
			controller.NotFoundHandler(w, r)
			return
		}

		modelWebsite := model.NewWebsiteModel()
		Website, err := modelWebsite.GetSettings()
		if err != nil {
			w.Write([]byte(err.Error()))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		data := make(map[string]any)

		data["Produk"] = Produk
		data["Website"] = Website
		data["Cookies"] = cookie.GetAllCookies(r)

		templates := []string{
			filepath.Join("views", "templates.html"),
			filepath.Join("views", "produk", "produk.html"),
		}

		tmpl, err := template.ParseFiles(templates...)
		if err != nil {
			w.Write([]byte(err.Error()))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		t := []string{
			"header", "produk", "footer",
		}

		for _, tl := range t {
			err = tmpl.ExecuteTemplate(w, tl, data)
			if err != nil {
				w.Write([]byte(err.Error()))
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		}
	}
}
