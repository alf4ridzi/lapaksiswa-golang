package kategori

import (
	"database/sql"
	"net/http"
)

func Kategori(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// vars := mux.Vars(r)
		// kategori := vars["kategori"]

		// websiteModel := model.NewWebsiteModel()
		// settings, err := websiteModel.GetSettings()
		// if err != nil {
		// 	w.Write([]byte(err.Error()))
		// 	w.WriteHeader(http.StatusInternalServerError)
		// 	return
		// }

		// modelProduk := model.NewProdukModel()

	}
}
