package controller

import (
	"net/http"
	"path/filepath"
	"text/template"

	"github.com/alf4ridzi/lapaksiswa-golang/config/cookie"
	"github.com/alf4ridzi/lapaksiswa-golang/model"
)

func Index() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			ErrorHandler(w, r, http.StatusNotFound)
			return
		}

		websiteModel := model.NewWebsiteModel()
		settings, err := websiteModel.GetSettings()

		if err != nil {
			w.Write([]byte(err.Error()))
			w.WriteHeader(http.StatusInternalServerError)
		}

		produkModel := model.NewProdukModel()
		terlaris, err := produkModel.GetTerlaris()

		if err != nil {
			w.Write([]byte(err.Error()))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		templates := []string{
			filepath.Join("views", "templates.html"),
			filepath.Join("views", "index.html"),
		}

		// ambil cookies
		data := map[string]interface{}{
			"Terlaris": terlaris,
			"Website":  settings,
			"Cookies":  cookie.GetAllCookies(w, r),
		}

		tmpl, err := template.ParseFiles(templates...)
		if err != nil {
			w.Write([]byte(err.Error()))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		t := []string{
			"header",
			"content",
			"footer",
		}

		for _, val := range t {
			err = tmpl.ExecuteTemplate(w, val, data)
			if err != nil {
				w.Write([]byte(err.Error()))
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		}

	}
}
