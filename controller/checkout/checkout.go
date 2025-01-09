package checkout

import (
	"encoding/json"
	"html/template"
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/alf4ridzi/lapaksiswa-golang/config/cookie"
	"github.com/alf4ridzi/lapaksiswa-golang/controller"
	"github.com/alf4ridzi/lapaksiswa-golang/controller/dashboard/api"
	"github.com/alf4ridzi/lapaksiswa-golang/lib"
	"github.com/alf4ridzi/lapaksiswa-golang/model"
)

func PageCheckOut() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			controller.NotFoundHandler(w, r)
			return
		}

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

		var product model.TambahTransaksi

		if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		templates := filepath.Join("views", "checkout", "checkout.html")
		tmpl, err := template.ParseFiles(templates)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if err := tmpl.Execute(w, nil); err != nil {
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
				data["result"] = err.Error()
				api.HandleResponseJson(w, data, http.StatusInternalServerError)
				return
			}

			data["result"] = "Product ID tidak valid"
			api.HandleResponseJson(w, data, http.StatusBadRequest)
			return
		}

		Product, err := lib.GetProduct(ProductID)
		if err != nil {
			data["result"] = err.Error()
			api.HandleResponseJson(w, data, http.StatusInternalServerError)
			return
		}

		TotalHarga := Product.Harga * Qty
		RandCheckout := lib.GenerateRandomStr(16)

		if err := lib.CreateCheckout(ProductID, RandCheckout, TotalHarga, username); err != nil {
			data["result"] = err.Error()
			api.HandleResponseJson(w, data, http.StatusInternalServerError)
			return
		}

		DataCookie := map[string]string{
			"checkout": RandCheckout,
		}

		cookie.SetCookie(w, DataCookie, cookie.CookieMaxAge)
		data["status"] = true
		api.HandleResponseJson(w, data, http.StatusOK)
		return
	}
}
