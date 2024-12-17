package routes

import (
	"database/sql"
	"net/http"

	"github.com/alf4ridzi/lapaksiswa-golang/controller"
	"github.com/alf4ridzi/lapaksiswa-golang/controller/autentikasi"
	"github.com/alf4ridzi/lapaksiswa-golang/controller/dashboard"
	"github.com/alf4ridzi/lapaksiswa-golang/controller/kategori"
	"github.com/alf4ridzi/lapaksiswa-golang/controller/produk"
	"github.com/alf4ridzi/lapaksiswa-golang/controller/search"
	"github.com/gorilla/mux"
)

func MapRoutes(server *mux.Router, db *sql.DB) {
	// index
	server.HandleFunc("/", controller.Index(db)).Methods("GET")

	// produk
	server.HandleFunc("/produk/{slug}", produk.Produk(db)).Methods("GET")

	// dashboard
	server.HandleFunc("/user", dashboard.User(db)).Methods("GET")

	// kategori
	server.HandleFunc("/kategori/{kategori}", kategori.Kategori(db)).Methods("GET")

	// search
	server.HandleFunc("/search", search.Search(db)).Methods("GET")

	// autentikasi
	server.HandleFunc("/login", autentikasi.Login(db)).Methods("GET", "POST")
	server.HandleFunc("/register", autentikasi.Register(db)).Methods("GET", "POST")
	server.HandleFunc("/keluar", autentikasi.Logout()).Methods("GET")

	server.PathPrefix("/").Handler(http.FileServer(http.Dir("./public/")))

	// 404 handler
	server.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		controller.NotFoundHandler(w, r)
	})

}
