package routes

import (
	"database/sql"
	"net/http"

	"github.com/alf4ridzi/lapaksiswa-golang/config"
	"github.com/alf4ridzi/lapaksiswa-golang/controller"
	"github.com/alf4ridzi/lapaksiswa-golang/controller/autentikasi"
	"github.com/alf4ridzi/lapaksiswa-golang/controller/dashboard"
	"github.com/alf4ridzi/lapaksiswa-golang/controller/produk"
	"github.com/gorilla/mux"
)

func MapRoutes(server *mux.Router, db *sql.DB) {
	// index
	server.HandleFunc("/", controller.Index(db)).Methods("GET")

	// produk
	server.HandleFunc("/produk/{slug}", produk.Produk(db)).Methods("GET")

	// dashboard
	server.HandleFunc("/user", dashboard.User(db)).Methods("GET")

	// autentikasi
	server.HandleFunc("/login", autentikasi.Login(db)).Methods("GET", "POST")
	server.HandleFunc("/register", autentikasi.Register(db)).Methods("GET", "POST")
	server.HandleFunc("/keluar", autentikasi.Logout()).Methods("GET")

	// 404 handler
	server.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		controller.NotFoundHandler(w, r)
	})

	server.Handle("/assets/", http.StripPrefix(
		"/assets/",
		http.FileServer(http.Dir("./"+config.GetPublicFolder()+"/assets")),
	))

	server.Handle("/js/", http.StripPrefix(
		"/js/",
		http.FileServer(http.Dir("./"+config.GetPublicFolder()+"/js")),
	))

	server.Handle("/css/", http.StripPrefix(
		"/css/",
		http.FileServer(http.Dir("./"+config.GetPublicFolder()+"/css")),
	))
}
