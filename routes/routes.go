package routes

import (
	"net/http"

	"github.com/alf4ridzi/lapaksiswa-golang/controller"
	"github.com/alf4ridzi/lapaksiswa-golang/controller/autentikasi"
	"github.com/alf4ridzi/lapaksiswa-golang/controller/dashboard"
	"github.com/alf4ridzi/lapaksiswa-golang/controller/dashboard/api"
	"github.com/alf4ridzi/lapaksiswa-golang/controller/kategori"
	"github.com/alf4ridzi/lapaksiswa-golang/controller/produk"
	"github.com/alf4ridzi/lapaksiswa-golang/controller/search"
	"github.com/gorilla/mux"
)

func MapRoutes(server *mux.Router) {
	// index
	server.HandleFunc("/", controller.Index()).Methods("GET")

	// produk
	server.HandleFunc("/produk/{slug}", produk.Produk()).Methods("GET")

	// dashboard user
	server.HandleFunc("/user", dashboard.User()).Methods("GET")
	server.HandleFunc("/user/edit", dashboard.Edit()).Methods("GET", "POST")

	// dashboard seller
	server.HandleFunc("/seller", dashboard.Seller()).Methods("GET")

	// kategori
	server.HandleFunc("/kategori/{kategori}", kategori.Kategori()).Methods("GET")

	// search
	server.HandleFunc("/search", search.Search()).Methods("GET")

	// update profile
	server.HandleFunc("/api/update-user", dashboard.UpdateData()).Methods("POST")
	server.HandleFunc("/api/update-profile", dashboard.UpdateProfile()).Methods("POST")

	// autentikasi
	server.HandleFunc("/login", autentikasi.Login()).Methods("GET", "POST")
	server.HandleFunc("/register", autentikasi.Register()).Methods("GET", "POST")
	server.HandleFunc("/keluar", autentikasi.Logout()).Methods("GET")
	server.HandleFunc("/reset-password", autentikasi.ResetPassword()).Methods("GET", "POST")
	server.HandleFunc("/daftar-seller", autentikasi.DaftarSeller()).Methods("GET", "POST")
	// seller api
	server.HandleFunc("/seller/get-product", api.GetProductToko()).Methods("GET")
	server.HandleFunc("/seller/tambah", api.TambahProduct()).Methods("POST")

	server.PathPrefix("/").Handler(http.FileServer(http.Dir("./public/")))

	// 404 handler
	server.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		controller.NotFoundHandler(w, r)
	})

}
