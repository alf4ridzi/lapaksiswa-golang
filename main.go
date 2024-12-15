package main

import (
	"net/http"

	"github.com/alf4ridzi/lapaksiswa-golang/database"
	"github.com/alf4ridzi/lapaksiswa-golang/routes"
	"github.com/gorilla/mux"
)

func main() {
	db := database.InitDatabase()
	server := mux.NewRouter()
	routes.MapRoutes(server, db)
	http.ListenAndServe(":8080", server)

}
