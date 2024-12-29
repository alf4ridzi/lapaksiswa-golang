package autentikasi

import (
	"encoding/json"
	"net/http"
	"path/filepath"
	"text/template"

	"github.com/alf4ridzi/lapaksiswa-golang/config/cookie"
	"github.com/alf4ridzi/lapaksiswa-golang/controller/dashboard/api"
	"github.com/alf4ridzi/lapaksiswa-golang/model"
)

func HandlePostDaftarSeller(w http.ResponseWriter, r *http.Request) {
	var responseJson []byte
	response := map[string]any{
		"status":  false,
		"message": nil,
	}

	w.Header().Set("Content-Type", "application/json")

	isLogin, err := r.Cookie("isLogin")
	if err != nil {
		response["message"] = "Login Dulu!"
		responseJson, err = api.ConvertMapToJson(response)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusBadRequest)
		w.Write(responseJson)
		return
	}

	if isLogin.Value != "true" {
		response["message"] = "Login Dulu!"
		responseJson, err = api.ConvertMapToJson(response)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusBadRequest)
		w.Write(responseJson)
		return
	}

	Username, err := r.Cookie("username")
	if err != nil {
		response["message"] = "Login Dulu!"
		responseJson, err = api.ConvertMapToJson(response)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusBadRequest)
		w.Write(responseJson)
		return
	}

	if Username.Value == "" {
		response["message"] = "Login Dulu!"
		responseJson, err = api.ConvertMapToJson(response)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusBadRequest)
		w.Write(responseJson)
		return
	}

	userModel := model.NewUserModel()
	isSeller, err := userModel.IsRole(Username.Value, "seller")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if isSeller {
		http.Redirect(w, r, "/seller", http.StatusPermanentRedirect)
		return
	}

	var seller model.Seller

	if err = json.NewDecoder(r.Body).Decode(&seller); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tokoModel := model.NewTokoModel()
	if err = tokoModel.DaftarToko(Username.Value, seller); err != nil {
		response["message"] = err.Error()
		responseJson, err = api.ConvertMapToJson(response)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusBadRequest)
		w.Write(responseJson)
		return
	}

	// ganti role user
	if err = userModel.ChangeRoleUser(Username.Value, "seller"); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response["status"] = true
	responseJson, err = api.ConvertMapToJson(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	cookieNew := map[string]string{
		"role": "seller",
	}

	cookie.SetCookie(w, cookieNew)
	w.WriteHeader(http.StatusOK)
	w.Write(responseJson)
}

func HandleGetDaftarSeller(w http.ResponseWriter, r *http.Request) {
	isLogin, err := r.Cookie("isLogin")
	if err != nil {
		cookie.SetFlashCookie(w, "error", "Login terlebih dahulu")
		http.Redirect(w, r, "/login", http.StatusPermanentRedirect)
		return
	}

	if isLogin.Value != "true" {
		cookie.SetFlashCookie(w, "error", "Login terlebih dahulu")
		http.Redirect(w, r, "/login", http.StatusPermanentRedirect)
		return
	}

	Username, err := r.Cookie("username")
	if err != nil {
		cookie.SetFlashCookie(w, "error", "Login terlebih dahulu")
		http.Redirect(w, r, "/login", http.StatusPermanentRedirect)
		return
	}

	if Username.Value == "" {
		cookie.SetFlashCookie(w, "error", "Login terlebih dahulu")
		http.Redirect(w, r, "/login", http.StatusPermanentRedirect)
		return
	}

	userModel := model.NewUserModel()
	isSeller, err := userModel.IsRole(Username.Value, "seller")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if isSeller {
		http.Redirect(w, r, "/seller", http.StatusPermanentRedirect)
		return
	}

	data := make(map[string]any)

	websiteModel := model.NewWebsiteModel()
	settings, err := websiteModel.GetSettings()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data["Website"] = settings
	data["Cookies"] = cookie.GetAllCookies(w, r)

	templates := []string{
		filepath.Join("views", "autentikasi", "daftar-seller.html"),
		filepath.Join("views", "templates.html"),
	}

	tmpl, err := template.ParseFiles(templates...)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	t := []string{
		"header", "content", "footer",
	}

	for _, tl := range t {
		if err = tmpl.ExecuteTemplate(w, tl, data); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

}

func DaftarSeller() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			HandlePostDaftarSeller(w, r)
		} else if r.Method == "GET" {
			HandleGetDaftarSeller(w, r)
		}
	}
}
