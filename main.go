package main

import (
	"net/http"

	"github.com/alf4ridzi/lapaksiswa-golang/routes"
	"github.com/gorilla/mux"
)

func main() {
	server := mux.NewRouter()
	routes.MapRoutes(server)
	http.ListenAndServe(":8080", server)

}
