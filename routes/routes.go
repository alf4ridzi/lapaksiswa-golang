package routes

import (
	"database/sql"
	"net/http"

	"github.com/alf4ridzi/lapaksiswa-golang/config"
	"github.com/alf4ridzi/lapaksiswa-golang/controller"
	"github.com/alf4ridzi/lapaksiswa-golang/controller/autentikasi"
	"github.com/alf4ridzi/lapaksiswa-golang/controller/dashboard"
)

func MapRoutes(server *http.ServeMux, db *sql.DB) {
	// index
	server.HandleFunc("/", controller.Index(db))

	// dashboard
	server.HandleFunc("/user", dashboard.User(db))

	// autentikasi
	server.HandleFunc("/login", autentikasi.Login(db))
	server.HandleFunc("/register", autentikasi.Register(db))
	server.HandleFunc("/keluar", autentikasi.Logout())

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
