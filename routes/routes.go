package routes

import (
	"net/http"

	"github.com/alf4ridzi/lapaksiswa-golang/controller"
	"github.com/alf4ridzi/lapaksiswa-golang/controller/autentikasi"
	"github.com/alf4ridzi/lapaksiswa-golang/controller/checkout"
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

	// checkout
	server.HandleFunc("/checkout", checkout.)
	// dashboard user
	server.HandleFunc("/user", dashboard.User()).Methods("GET")
	server.HandleFunc("/user/edit", dashboard.Edit()).Methods("GET", "POST")

	// dashboard seller
	server.HandleFunc("/seller", dashboard.Seller()).Methods("GET")
	server.HandleFunc("/seller/list-product", dashboard.ListProductPage()).Methods("GET")
	server.HandleFunc("/seller/list-orders", dashboard.ListOrderPage()).Methods("GET")
	server.HandleFunc("/seller/edit-product/{id}", dashboard.EditProduct()).Methods("GET")
	server.HandleFunc("/seller/edit-product", api.EditProduct()).Methods("POST")
	server.HandleFunc("/seller/delete-product", api.DeleteProduct()).Methods("POST")
	server.HandleFunc("/seller/profile", dashboard.ProfilePage()).Methods("GET")
	server.HandleFunc("/seller/update-toko", api.UpdateToko()).Methods("POST")
	server.HandleFunc("/seller/update-picture", api.UpdateFotoToko()).Methods("POST")

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
	server.HandleFunc("/seller/update-status", api.UpdateStatusProduct()).Methods("POST")

	// produk
	server.HandleFunc("/{toko}/{slug}", produk.Produk()).Methods("GET")
	server.HandleFunc("/api/add-cart", produk.TambahCart()).Methods("POST")

	// public folder
	server.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("./public/"))))

	// 404 handler
	server.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		controller.NotFoundHandler(w, r)
	})

}
