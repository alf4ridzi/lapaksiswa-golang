package produk

import (
	"net/http"
	"path/filepath"

	"github.com/alf4ridzi/lapaksiswa-golang/config/cookie"
	"github.com/alf4ridzi/lapaksiswa-golang/controller/dashboard/api"
	"github.com/alf4ridzi/lapaksiswa-golang/lib"
)

func PageCart() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		isLogin, err := cookie.GetCookieValue(r, "isLogin")
		if err != nil {
			cookie.SetFlashCookie(w, "error", "Login terlebih dahulu!")
			http.Redirect(w, r, "/login", http.StatusPermanentRedirect)
			return
		}

		if isLogin != "true" {
			cookie.SetFlashCookie(w, "error", "Login terlebih dahulu!")
			http.Redirect(w, r, "/login", http.StatusPermanentRedirect)
			return
		}

		username, err := cookie.GetCookieValue(r, "username")
		if err != nil {
			cookie.SetFlashCookie(w, "error", "Login terlebih dahulu!")
			http.Redirect(w, r, "/login", http.StatusPermanentRedirect)
			return
		}

		if username == "" {
			cookie.SetFlashCookie(w, "error", "Login terlebih dahulu!")
			http.Redirect(w, r, "/login", http.StatusPermanentRedirect)
			return
		}

		data := make(map[string]any)

		settings, err := lib.GetSettingsWebsite()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		data["Website"] = settings
		data["Cookies"] = cookie.GetAllCookies(w, r)

		templates := []string{
			filepath.Join("views", "templates.html"),
			filepath.Join("views", "checkout", "keranjang.html"),
		}

		t := []string{
			"header", "cart", "footer",
		}

		if err := lib.RenderTemplate(w, templates, t, data); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func TambahCart() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		data := map[string]any{
			"result": nil,
			"status": false,
		}

		isLogin, err := cookie.GetCookieValue(r, "isLogin")
		if err != nil {
			data["result"] = err.Error()
			api.HandleResponseJson(w, data, http.StatusInternalServerError)
			return
		}

		if isLogin != "true" {
			data["result"] = "Login terlebih dahulu!"
			api.HandleResponseJson(w, data, http.StatusOK)
			return
		}

		username, err := cookie.GetCookieValue(r, "username")
		if err != nil {
			data["result"] = err.Error()
			api.HandleResponseJson(w, data, http.StatusInternalServerError)
			return
		}

		if username == "" {
			data["result"] = "Login terlebih dahulu!"
			api.HandleResponseJson(w, data, http.StatusOK)
			return
		}

		ProductID := r.FormValue("productid")

		if ProductID == "" {
			data["result"] = "Product ID tidak boleh kosong"
			api.HandleResponseJson(w, data, http.StatusOK)
			return
		}

		isValid, err := lib.IsValidProductID(ProductID)
		if err != nil {
			data["result"] = err.Error()
			api.HandleResponseJson(w, data, http.StatusInternalServerError)
			return
		}

		if !isValid {
			data["result"] = "Product tidak valid"
			api.HandleResponseJson(w, data, http.StatusOK)
			return
		}

		if err := lib.InsertIntoCart(username, ProductID); err != nil {
			data["result"] = err.Error()
			api.HandleResponseJson(w, data, http.StatusOK)
			return
		}

		data["status"] = true
		api.HandleResponseJson(w, data, http.StatusOK)
	}
}
