package checkout

import (
	"html/template"
	"net/http"
	"path/filepath"
)

func PageCheckOut() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		templates := filepath.Join("views", "checkout", "checkout.html")
		tmpl, err := template.ParseFiles(templates)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if err := tmpl.Execute(w, nil); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
