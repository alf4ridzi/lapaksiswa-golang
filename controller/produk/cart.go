package produk

import (
	"net/http"

	"github.com/alf4ridzi/lapaksiswa-golang/config/cookie"
	"github.com/alf4ridzi/lapaksiswa-golang/controller/dashboard/api"
)

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

		ProductID := r.FormValue("productid")

		data["status"] = true
		api.HandleResponseJson(w, data, http.StatusOK)
	}
}
