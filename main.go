package main

import (
	"fmt"
	"net/http"

	"github.com/alf4ridzi/lapaksiswa-golang/database"
	"github.com/alf4ridzi/lapaksiswa-golang/routes"
	"github.com/gorilla/mux"
)

func main() {
	db, err := database.InitDatabase()
	if err != nil {
		fmt.Printf("Kesalahan database : %s", err.Error())
		return
	}

	db.Close()
	server := mux.NewRouter()
	routes.MapRoutes(server)
	http.ListenAndServe(":8080", server)
}
