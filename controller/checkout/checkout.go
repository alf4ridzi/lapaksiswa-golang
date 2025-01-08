package checkout

import (
	"net/http"
	"path/filepath"

	"github.com/alf4ridzi/lapaksiswa-golang/lib"
)

func PageCheckOut() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		templates := []string{
			filepath.Join("views", "checkout", "checkout.html"),
		}

		t := []string{
			"checkout",
		}

		if err := lib.RenderTemplate(w, templates, t, nil); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
