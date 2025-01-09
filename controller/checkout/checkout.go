package checkout

import (
	"html/template"
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/alf4ridzi/lapaksiswa-golang/config/cookie"
	"github.com/alf4ridzi/lapaksiswa-golang/controller/dashboard/api"
	"github.com/alf4ridzi/lapaksiswa-golang/lib"
	"github.com/alf4ridzi/lapaksiswa-golang/model"
	"github.com/gorilla/mux"
)

func PageCheckOut() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		isLogin, err := cookie.GetCookieValue(r, "isLogin")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if isLogin != "true" {
			cookie.SetFlashCookie(w, "error", "Login terlebih dahulu sebelum checkout!")
			http.Redirect(w, r, "/login", http.StatusBadRequest)
			return
		}

		username, err := cookie.GetCookieValue(r, "username")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if username == "" {
			cookie.SetFlashCookie(w, "error", "Login terlebih dahulu sebelum checkout!")
			http.Redirect(w, r, "/login", http.StatusBadRequest)
			return
		}

		vars := mux.Vars(r)
		id := vars["checkout"]

		if isValid, err := lib.IsValidCheckout(id, username); err != nil || !isValid {
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			http.Error(w, "Checkout tidak valid", http.StatusBadRequest)
			return
		}

		data := make(map[string]any)

		Alamat, err := lib.GetUserSpec(username, "alamat")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		Nama, err := lib.GetUserSpec(username, "nama")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		Phone, err := lib.GetUserSpec(username, "no_hp")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		tokoModel := model.NewTokoModel()
		defer tokoModel.DB.Close()

		Toko, err := tokoModel.GetTokoByUsername(username)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		DetailCheckout, err := lib.GetDetailCheckout(id, username)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		Produk, err := lib.GetProduct(DetailCheckout.ProductID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		Pembayaran, err := lib.GetMetodePembayaran()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		data["Alamat"] = Alamat
		data["Nama"] = Nama
		data["Produk"] = Produk
		data["Detail"] = DetailCheckout
		data["Payment"] = Pembayaran
		data["Phone"] = Phone
		data["Toko"] = Toko
		data["Checkout"] = id

		templates := filepath.Join("views", "checkout", "checkout.html")
		tmpl, err := template.ParseFiles(templates)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if err := tmpl.Execute(w, data); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func CreateCheckOutURL() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		data := map[string]any{
			"status": false,
			"result": nil,
		}

		isLogin, err := cookie.GetCookieValue(r, "isLogin")
		if err != nil {
			data["result"] = err.Error()
			api.HandleResponseJson(w, data, http.StatusInternalServerError)
			return
		}

		if isLogin != "true" {
			cookie.SetFlashCookie(w, "error", "Login terlebih dahulu sebelum checkout!")
			http.Redirect(w, r, "/login", http.StatusBadRequest)
			return
		}

		username, err := cookie.GetCookieValue(r, "username")
		if err != nil {
			data["result"] = err.Error()
			api.HandleResponseJson(w, data, http.StatusInternalServerError)
			return
		}

		if username == "" {
			cookie.SetFlashCookie(w, "error", "Login terlebih dahulu sebelum checkout!")
			http.Redirect(w, r, "/login", http.StatusBadRequest)
			return
		}

		ProductID := r.FormValue("productid")
		QtyParse, err := strconv.ParseInt(r.FormValue("qty"), 10, 64)
		Qty := QtyParse

		if isValid, err := lib.IsValidProductID(ProductID); err != nil || !isValid {
			if err != nil {
				data["result"] = "1" + err.Error()
				api.HandleResponseJson(w, data, http.StatusInternalServerError)
				return
			}

			data["result"] = "Product ID tidak valid"
			api.HandleResponseJson(w, data, http.StatusBadRequest)
			return
		}

		Product, err := lib.GetProduct(ProductID)
		if err != nil {
			data["result"] = "2" + err.Error()
			api.HandleResponseJson(w, data, http.StatusInternalServerError)
			return
		}

		TotalHarga := Product.Harga * Qty
		RandCheckout := lib.GenerateRandomStr(16)

		if err := lib.CreateCheckout(ProductID, RandCheckout, TotalHarga, username, Qty); err != nil {
			data["result"] = "3" + err.Error()
			api.HandleResponseJson(w, data, http.StatusInternalServerError)
			return
		}

		DataCookie := map[string]string{
			"checkout": RandCheckout,
		}

		cookie.SetCookie(w, DataCookie, cookie.CookieMaxAge)
		data["status"] = true
		data["result"] = RandCheckout
		api.HandleResponseJson(w, data, http.StatusOK)
		return
	}
}
