package api

import (
	"encoding/json"
	"net/http"

	"github.com/alf4ridzi/lapaksiswa-golang/model"
)

func ConvertMapToJson(data map[string]any) ([]byte, error) {
	Json, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	return Json, nil
}

func GetProductToko() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// if r.Method != "GET" || r.Method != "" {
		// 	http.Error(w, "GET DOANG OI", http.StatusBadRequest)
		// 	return
		// }

		data := map[string]any{
			"status": false,
			"result": "",
		}

		isLogin, err := r.Cookie("isLogin")
		if err != nil {
			data["result"] = "Login terlebih dahulu!"
			w.WriteHeader(http.StatusBadRequest)

			responseJson, err := ConvertMapToJson(data)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			w.Write(responseJson)
			return
		}

		if isLogin.Value != "true" {
			data["result"] = "Login terlebih dahulu!"
			w.WriteHeader(http.StatusBadRequest)

			responseJson, err := ConvertMapToJson(data)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			w.Write(responseJson)
			return
		}

		Username, err := r.Cookie("username")
		if err != nil {
			data["result"] = "Login terlebih dahulu!"
			w.WriteHeader(http.StatusBadRequest)

			responseJson, err := ConvertMapToJson(data)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			w.Write(responseJson)
			return
		}

		if Username.Value == "" {
			data["result"] = "Login terlebih dahulu!"
			w.WriteHeader(http.StatusBadRequest)

			responseJson, err := ConvertMapToJson(data)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			w.Write(responseJson)
			return
		}

		userModel := model.NewUserModel()
		isSeller, err := userModel.IsRole(Username.Value, "seller")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if !isSeller {
			data["result"] = "Khusus seller"
			w.WriteHeader(http.StatusBadRequest)

			responseJson, err := ConvertMapToJson(data)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			w.Write(responseJson)
			return
		}

		tokoModel := model.NewTokoModel()
		Toko, err := tokoModel.GetTokoByUsername(Username.Value)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		produkModel := model.NewProdukModel()
		Produk, err := produkModel.GetProductToko(Toko.Domain)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		data["status"] = true
		data["result"] = Produk
		responseJson, err := ConvertMapToJson(data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(responseJson)
	}
}
