package main

import (
	"net/http"

	"github.com/alf4ridzi/lapaksiswa-golang/database"
	"github.com/alf4ridzi/lapaksiswa-golang/routes"
	"github.com/gorilla/mux"
)

func main() {
	db, err := database.InitDatabase()
	if err != nil {
		panic("Error Database : " + err.Error())
	}

	db.Close()
	server := mux.NewRouter()
	routes.MapRoutes(server)
	http.ListenAndServe(":8080", server)
}
