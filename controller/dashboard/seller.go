package dashboard

import (
	"fmt"
	"net/http"
	"path/filepath"
	"text/template"

	"github.com/alf4ridzi/lapaksiswa-golang/config/cookie"
	"github.com/alf4ridzi/lapaksiswa-golang/controller"
	"github.com/alf4ridzi/lapaksiswa-golang/lib"
	"github.com/alf4ridzi/lapaksiswa-golang/model"
	"github.com/gorilla/mux"
)

func Seller() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/seller" {
			controller.NotFoundHandler(w, r)
			return
		}

		isLogin, err := r.Cookie("isLogin")
		if err != nil {
			cookie.SetFlashCookie(w, "error", "Silahkan login terlebih dahulu!")
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}

		if isLogin.Value != "true" {
			cookie.SetFlashCookie(w, "error", "Silahkan login terlebih dahulu!")
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}

		Username, err := r.Cookie("username")
		if err != nil {
			cookie.SetFlashCookie(w, "error", "Silahkan login terlebih dahulu!")
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}

		if Username.Value == "" {
			cookie.SetFlashCookie(w, "error", "Silahkan login terlebih dahulu!")
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}

		tokoModel := model.NewTokoModel()
		defer tokoModel.DB.Close()
		Toko, err := tokoModel.GetTokoByUsername(Username.Value)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		if Toko.Domain == "" {
			cookie.SetFlashCookie(w, "error", "Area khusus seller. Silahkan login sebagai seller")
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}

		templates := []string{
			filepath.Join("views", "dashboard", "seller", "templates.html"),
			filepath.Join("views", "dashboard", "seller", "index.html"),
		}

		tmpl, err := template.ParseFiles(templates...)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		data := make(map[string]any)

		websiteModel := model.NewWebsiteModel()
		defer websiteModel.DB.Close()

		settings, err := websiteModel.GetSettings()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		Saldo, err := tokoModel.GetSaldoToko(Toko.Domain)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		orderModel := model.NewOrderModel()
		defer orderModel.DB.Close()

		TrxHariIni, err := orderModel.GetTransaksiHariIni(Toko.Domain)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		TrxKemarin, err := orderModel.GetTransaksiKemarin(Toko.Domain)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		Omset, err := orderModel.GetTotalOmset(Toko.Domain)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		OmsetBulanan, err := orderModel.GetOmsetBulanan(Toko.Domain)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		TotalTrx, err := orderModel.GetTotalTransaksi(Toko.Domain)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		kategoriModel := model.NewKategoriModel()
		defer kategoriModel.DB.Close()

		Kategori, err := kategoriModel.GetKategori()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		data["Saldo"] = Saldo
		data["TotalTrx"] = TotalTrx
		data["OmsetBulanan"] = OmsetBulanan
		data["Omset"] = Omset
		data["Kemarin"] = TrxKemarin
		data["HariIni"] = TrxHariIni
		data["Website"] = settings
		data["Toko"] = Toko
		data["Kategori"] = Kategori

		t := []string{
			"header", "navbar", "content", "end",
		}

		for _, tl := range t {
			if err = tmpl.ExecuteTemplate(w, tl, data); err != nil {
				fmt.Println(err.Error())
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}
	}
}

func HandleGetEditProduct(w http.ResponseWriter, r *http.Request) {
	isLogin, err := r.Cookie("isLogin")
	if err != nil {
		cookie.SetFlashCookie(w, "error", "Login terlebih dahulu.")
		http.Redirect(w, r, "/login", http.StatusPermanentRedirect)
		return
	}

	if isLogin.Value != "true" {
		cookie.SetFlashCookie(w, "error", "Login terlebih dahulu.")
		http.Redirect(w, r, "/login", http.StatusPermanentRedirect)
		return
	}

	role, err := r.Cookie("role")
	if err != nil {
		cookie.SetFlashCookie(w, "error", "Login terlebih dahulu.")
		http.Redirect(w, r, "/login", http.StatusPermanentRedirect)
		return
	}

	if role.Value != "seller" {
		http.Redirect(w, r, "/", http.StatusPermanentRedirect)
		return
	}

	Username, err := r.Cookie("username")
	if err != nil {
		cookie.SetFlashCookie(w, "error", "Login terlebih dahulu.")
		http.Redirect(w, r, "/login", http.StatusPermanentRedirect)
		return
	}

	if Username.Value == "" {
		cookie.SetFlashCookie(w, "error", "Login terlebih dahulu.")
		http.Redirect(w, r, "/login", http.StatusPermanentRedirect)
		return
	}

	tokoModel := model.NewTokoModel()
	defer tokoModel.DB.Close()

	Toko, err := tokoModel.GetTokoByUsername(Username.Value)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if Toko.Domain == "" {
		cookie.SetFlashCookie(w, "error", "Silahkan login menggunakan akun seller.")
		http.Redirect(w, r, "/login", http.StatusPermanentRedirect)
		return
	}

	vars := mux.Vars(r)
	ProductID := vars["id"]

	produkModel := model.NewProdukModel()
	defer produkModel.DB.Close()

	if isValid, err := produkModel.ValidasiProduk(ProductID, Toko.Domain); err != nil || !isValid {
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		controller.NotFoundHandler(w, r)
		return
	}

	Product, err := produkModel.GetProductByID(ProductID, Toko.Domain)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	kategoriModel := model.NewKategoriModel()
	defer kategoriModel.DB.Close()

	Kategori, err := kategoriModel.GetKategori()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := make(map[string]any)

	data["Product"] = Product
	data["Kategori"] = Kategori

	templates := []string{
		filepath.Join("views", "dashboard", "seller", "edit.html"),
		filepath.Join("views", "dashboard", "seller", "index.html"),
		filepath.Join("views", "dashboard", "seller", "templates.html"),
	}

	tmpl, err := template.ParseFiles(templates...)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tl := []string{
		"header", "navbar", "edit", "end",
	}

	for _, t := range tl {
		if err = tmpl.ExecuteTemplate(w, t, data); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

}

func HandlePostEditProduct(w http.ResponseWriter, r *http.Request) {

}

func EditProduct() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			HandlePostEditProduct(w, r)
		} else if r.Method == "GET" || r.Method == "" {
			HandleGetEditProduct(w, r)
		}

	}
}

func ListProductPage() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		isLogin, err := r.Cookie("isLogin")
		if err != nil {
			cookie.SetFlashCookie(w, "error", "Login terlebih dahulu.")
			http.Redirect(w, r, "/login", http.StatusPermanentRedirect)
			return
		}

		if isLogin.Value != "true" {
			cookie.SetFlashCookie(w, "error", "Login terlebih dahulu.")
			http.Redirect(w, r, "/login", http.StatusPermanentRedirect)
			return
		}

		role, err := r.Cookie("role")
		if err != nil {
			cookie.SetFlashCookie(w, "error", "Login terlebih dahulu.")
			http.Redirect(w, r, "/login", http.StatusPermanentRedirect)
			return
		}

		if role.Value != "seller" {
			http.Redirect(w, r, "/", http.StatusPermanentRedirect)
			return
		}

		Username, err := r.Cookie("username")
		if err != nil {
			cookie.SetFlashCookie(w, "error", "Login terlebih dahulu.")
			http.Redirect(w, r, "/login", http.StatusPermanentRedirect)
			return
		}

		if Username.Value == "" {
			cookie.SetFlashCookie(w, "error", "Login terlebih dahulu.")
			http.Redirect(w, r, "/login", http.StatusPermanentRedirect)
			return
		}

		tokoModel := model.NewTokoModel()
		defer tokoModel.DB.Close()

		Toko, err := tokoModel.GetTokoByUsername(Username.Value)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if Toko.Domain == "" {
			cookie.SetFlashCookie(w, "error", "Silahkan login menggunakan akun seller.")
			http.Redirect(w, r, "/login", http.StatusPermanentRedirect)
			return
		}

		kategoriModel := model.NewKategoriModel()
		defer kategoriModel.DB.Close()

		Kategori, err := kategoriModel.GetKategori()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		data := make(map[string]any)
		data["Toko"] = Toko
		data["Kategori"] = Kategori

		templates := []string{
			filepath.Join("views", "dashboard", "seller", "list-product.html"),
			filepath.Join("views", "dashboard", "seller", "templates.html"),
		}

		tmpl, err := template.ParseFiles(templates...)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		t := []string{
			"header", "navbar", "content", "end",
		}

		for _, tl := range t {
			if err = tmpl.ExecuteTemplate(w, tl, data); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}
	}
}

func ListOrderPage() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		isLogin, err := r.Cookie("isLogin")
		if err != nil {
			cookie.SetFlashCookie(w, "error", "Login terlebih dahulu.")
			http.Redirect(w, r, "/login", http.StatusPermanentRedirect)
			return
		}

		if isLogin.Value != "true" {
			cookie.SetFlashCookie(w, "error", "Login terlebih dahulu.")
			http.Redirect(w, r, "/login", http.StatusPermanentRedirect)
			return
		}

		role, err := r.Cookie("role")
		if err != nil {
			cookie.SetFlashCookie(w, "error", "Login terlebih dahulu.")
			http.Redirect(w, r, "/login", http.StatusPermanentRedirect)
			return
		}

		if role.Value != "seller" {
			http.Redirect(w, r, "/", http.StatusPermanentRedirect)
			return
		}

		Username, err := r.Cookie("username")
		if err != nil {
			cookie.SetFlashCookie(w, "error", "Login terlebih dahulu.")
			http.Redirect(w, r, "/login", http.StatusPermanentRedirect)
			return
		}

		if Username.Value == "" {
			cookie.SetFlashCookie(w, "error", "Login terlebih dahulu.")
			http.Redirect(w, r, "/login", http.StatusPermanentRedirect)
			return
		}

		tokoModel := model.NewTokoModel()
		defer tokoModel.DB.Close()

		Toko, err := tokoModel.GetTokoByUsername(Username.Value)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if Toko.Domain == "" {
			cookie.SetFlashCookie(w, "error", "Silahkan login menggunakan akun seller.")
			http.Redirect(w, r, "/login", http.StatusPermanentRedirect)
			return
		}

		data := make(map[string]any)

		trxModel := model.NewTransaksiModel()
		defer trxModel.DB.Close()

		trx, err := trxModel.GetAllTransaksi(Toko.Domain)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		data["trx"] = trx
		data["Toko"] = Toko

		templates := []string{
			filepath.Join("views", "dashboard", "seller", "list-order.html"),
			filepath.Join("views", "dashboard", "seller", "templates.html"),
		}

		tmpl, err := template.ParseFiles(templates...)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		t := []string{
			"header", "navbar", "content", "end",
		}

		for _, tl := range t {
			if err = tmpl.ExecuteTemplate(w, tl, data); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}
	}
}

func ProfilePage() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if isCookie, err := lib.ValidateUserCookies(r, "seller"); err != nil || !isCookie {
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			cookie.SetFlashCookie(w, "error", "Login terlebih dahulu!")
			http.Redirect(w, r, "/login", http.StatusForbidden)
			return
		}

		Username, err := cookie.GetCookieValue(r, "username")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if Username == "" {
			cookie.SetFlashCookie(w, "error", "Login terlebih dahulu!")
			http.Redirect(w, r, "/login", http.StatusForbidden)
			return
		}

		data := make(map[string]any)

		tokoModel := model.NewTokoModel()
		defer tokoModel.DB.Close()

		Toko, err := tokoModel.GetTokoByUsername(Username)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if Toko == nil {
			http.Error(w, "Toko tidak ditemukan", http.StatusBadRequest)
			return
		}

		kategoriModel := model.NewKategoriModel()
		defer kategoriModel.DB.Close()

		Kategori, err := kategoriModel.GetKategori()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		data["Kategori"] = Kategori
		data["Toko"] = Toko

		templates := []string{
			filepath.Join("views", "dashboard", "seller", "templates.html"),
			filepath.Join("views", "dashboard", "seller", "profile.html"),
		}

		t := []string{
			"header", "navbar", "content", "end",
		}

		if err := lib.RenderTemplate(w, templates, t, data); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
