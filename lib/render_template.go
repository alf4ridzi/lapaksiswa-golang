package lib

import (
	"net/http"
	"text/template"
)

func RenderTemplate(w http.ResponseWriter, T []string, Nama []string, data map[string]any) error {
	tmpl, err := template.ParseFiles(T...)
	if err != nil {
		return err
	}

	for _, n := range Nama {
		if err = tmpl.ExecuteTemplate(w, n, data); err != nil {
			return err
		}
	}

	return nil
}
