package produk

import (
	"net/http"
	"path/filepath"
	"text/template"

	"github.com/alf4ridzi/lapaksiswa-golang/config/cookie"
	"github.com/alf4ridzi/lapaksiswa-golang/controller"
	"github.com/alf4ridzi/lapaksiswa-golang/model"
	"github.com/gorilla/mux"
)

func Produk() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		slug := vars["slug"]

		modelProduk := model.NewProdukModel()
		defer modelProduk.DB.Close()

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
		defer modelWebsite.DB.Close()

		Website, err := modelWebsite.GetSettings()
		if err != nil {
			w.Write([]byte(err.Error()))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		Related, err := modelProduk.GetRelated(Produk)
		if err != nil {
			w.Write([]byte(err.Error()))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		modelToko := model.NewTokoModel()
		defer modelToko.DB.Close()

		Toko, err := modelToko.GetToko(Produk.Domain)
		if err != nil {
			w.Write([]byte(err.Error()))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		modelTransaksi := model.NewTransaksiModel()
		ProdukTerjual, err := modelTransaksi.GetTotalProdukTerjual(Toko.Domain)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		data := make(map[string]any)

		data["Produk"] = Produk
		data["Website"] = Website
		data["Cookies"] = cookie.GetAllCookies(w, r)
		data["Related"] = Related
		data["Toko"] = Toko
		data["Terjual"] = ProdukTerjual

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
			if err = tmpl.ExecuteTemplate(w, tl, data); err != nil {
				w.Write([]byte(err.Error()))
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		}
	}
}
