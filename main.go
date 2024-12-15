package main

import (
	"net/http"

	"github.com/alf4ridzi/lapaksiswa-golang/database"
	"github.com/alf4ridzi/lapaksiswa-golang/routes"
)

func main() {
	db := database.InitDatabase()
	server := http.NewServeMux()
	routes.MapRoutes(server, db)
	http.ListenAndServe(":8080", server)

}
