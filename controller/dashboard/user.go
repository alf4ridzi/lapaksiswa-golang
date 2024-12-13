package dashboard

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"path/filepath"
	"text/template"

	"github.com/alf4ridzi/lapaksiswa-golang/config/cookie"
	"github.com/alf4ridzi/lapaksiswa-golang/controller"
	"github.com/alf4ridzi/lapaksiswa-golang/model"
)

type Profile struct {
	Nama         string `json:"nama"`
	TanggalLahir string `json:"tanggalLahir"`
	JenisKelamin string `json:"jenisKelamin"`
	Email        string `json:"email"`
	NoHP         string `json:"noHP"`
}

func User(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/user" {
			controller.ErrorHandler(w, r, http.StatusPermanentRedirect)
			return
		}

		isLogin, err := r.Cookie("isLogin")
		if err != nil || isLogin.Value != "true" {
			cookie.SetFlashCookie(w, "error", "Login terlebih dahulu!")
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}

		role, err := r.Cookie("role")
		if err != nil || role.Value != "user" {
			http.Redirect(w, r, "/", http.StatusFound)
			return
		}

		username, err := r.Cookie("username")
		if err != nil || username.Value == "" {
			cookie.SetFlashCookie(w, "error", "Login terlebih dahulu!")
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}

		templates := []string{
			filepath.Join("views", "templates.html"),
			filepath.Join("views", "dashboard", "user.html"),
		}

		tmpl, err := template.ParseFiles(templates...)
		if err != nil {
			http.Error(w, "Internal Server Error: "+err.Error(), http.StatusInternalServerError)
			return
		}

		websiteModel := model.NewWebsiteModel()
		settings, err := websiteModel.GetSettings()
		if err != nil {
			http.Error(w, "Internal Server Error: "+err.Error(), http.StatusInternalServerError)
			return
		}

		userModel := model.NewUserModel()
		user, err := userModel.GetUser(username.Value)
		if err != nil {
			http.Error(w, "Internal Server Error: "+err.Error(), http.StatusInternalServerError)
			return
		}

		data := map[string]any{
			"Website": settings,
			"Cookies": cookie.GetAllCookies(r),
			"User":    user,
		}

		for _, section := range []string{"header", "dashboard", "footer"} {
			err = tmpl.ExecuteTemplate(w, section, data)
			if err != nil {
				http.Error(w, "Internal Server Error: "+err.Error(), http.StatusInternalServerError)
				return
			}
		}
	}
}

func Edit(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/user/edit" {
			controller.NotFoundHandler(w, r)
			return
		}

		isLogin, err := r.Cookie("isLogin")
		if err != nil || isLogin.Value != "true" {
			cookie.SetFlashCookie(w, "error", "Login terlebih dahulu!")
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}

		role, err := r.Cookie("role")
		if err != nil || role.Value != "user" {
			http.Redirect(w, r, "/", http.StatusFound)
			return
		}

		username, err := r.Cookie("username")
		if err != nil || username.Value == "" {
			cookie.SetFlashCookie(w, "error", "Login terlebih dahulu!")
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}

		if r.Method == "POST" {

		} else if r.Method == "GET" || r.Method == "" {
			templates := []string{
				filepath.Join("views", "templates.html"),
				filepath.Join("views", "dashboard", "edit.html"),
			}

			tmpl, err := template.ParseFiles(templates...)
			if err != nil {
				w.Write([]byte(err.Error()))
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			modelWebsite := model.NewWebsiteModel()
			settings, err := modelWebsite.GetSettings()
			if err != nil {
				w.Write([]byte(err.Error()))
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			modelUser := model.NewUserModel()
			User, err := modelUser.GetUser(username.Value)

			if err != nil {
				w.Write([]byte(err.Error()))
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			data := make(map[string]any)

			data["Website"] = settings
			data["Cookies"] = cookie.GetAllCookies(r)
			data["User"] = User

			tName := []string{
				"header",
				"edit",
				"footer",
			}

			for _, t := range tName {
				if err = tmpl.ExecuteTemplate(w, t, data); err != nil {
					w.Write([]byte(err.Error()))
					w.WriteHeader(http.StatusInternalServerError)
					return
				}
			}
		}
	}
}

func UpdateData() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/update-user" || r.Method != "POST" {
			controller.NotFoundHandler(w, r)
			return
		}

		isLogin, err := r.Cookie("isLogin")
		if err != nil {
			cookie.SetFlashCookie(w, "error", "Silahkan login terlebih dahulu!")
			http.Redirect(w, r, "/login", http.StatusPermanentRedirect)
			return
		}

		if isLogin.Value != "true" {
			cookie.SetFlashCookie(w, "error", "Silahkan login terlebih dahulu!")
			http.Redirect(w, r, "/login", http.StatusPermanentRedirect)
			return
		}

		Username, err := r.Cookie("username")
		if err != nil {
			cookie.SetFlashCookie(w, "error", "Silahkan login terlebih dahulu!")
			http.Redirect(w, r, "/login", http.StatusPermanentRedirect)
			return
		}

		if Username.Value == "" {
			cookie.SetFlashCookie(w, "error", "Silahkan login terlebih dahulu!")
			http.Redirect(w, r, "/login", http.StatusPermanentRedirect)
			return
		}

		var profile Profile

		data := make(map[string]any)

		err = json.NewDecoder(r.Body).Decode(&profile)
		if err != nil {
			w.Write([]byte(err.Error()))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		userModel := model.NewUserModel()
		w.Header().Set("Content-Type", "application/json")
		// update profile
		if err = userModel.UpdateProfileUser(Username.Value, profile.Nama, profile.TanggalLahir, profile.JenisKelamin, profile.Email, profile.NoHP); err != nil {
			data["success"] = false
			data["msg"] = "Error " + err.Error()

			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(data)
			return
		}

		data["success"] = true
		data["msg"] = "Berhasil update profile"
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(data)
	}
}
