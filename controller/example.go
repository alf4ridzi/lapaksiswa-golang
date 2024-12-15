package controller

import (
	"database/sql"
	"net/http"
)

func Example(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
