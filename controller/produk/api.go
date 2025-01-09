package produk

import (
	"net/http"

	"github.com/alf4ridzi/lapaksiswa-golang/config/cookie"
	"github.com/alf4ridzi/lapaksiswa-golang/controller/dashboard/api"
	"github.com/alf4ridzi/lapaksiswa-golang/lib"
)

func GetCart() func(w http.ResponseWriter, r *http.Request) {
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
			data["result"] = "Login terlebih dahulu!"
			api.HandleResponseJson(w, data, http.StatusBadRequest)
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
			api.HandleResponseJson(w, data, http.StatusBadRequest)
			return
		}

		listProduct, err := lib.GetUserCart(username)
		if err != nil {
			data["result"] = err.Error()
			api.HandleResponseJson(w, data, http.StatusInternalServerError)
			return
		}

		data["status"] = true
		data["result"] = listProduct
		api.HandleResponseJson(w, data, http.StatusOK)
	}
}
