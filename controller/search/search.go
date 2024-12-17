package search

import (
	"database/sql"
	"net/http"
)

func Search(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		Keyword := r.URL.Query().Get("keyword")

	}
}
