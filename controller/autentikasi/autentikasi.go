package autentikasi

import (
	"database/sql"
	"net/http"
	"path/filepath"
	"regexp"
	"text/template"

	"github.com/alf4ridzi/lapaksiswa-golang/config/cookie"
	"github.com/alf4ridzi/lapaksiswa-golang/controller"
	"github.com/alf4ridzi/lapaksiswa-golang/model"
)

func isValidEmail(email string) bool {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return emailRegex.MatchString(email)
}

func Login(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/login" {
			controller.ErrorHandler(w, r, http.StatusNotFound)
			return
		}

		if r.Method == "POST" {
			checkLogin, err := r.Cookie("isLogin")
			if err == nil && checkLogin.Value == "true" {
				role, err := r.Cookie("role")
				if err != nil {
					http.Redirect(w, r, "/", http.StatusUnauthorized)
					return
				}

				switch role.Value {
				case "admin":
					http.Redirect(w, r, "/admin", http.StatusFound)
				case "seller":
					http.Redirect(w, r, "/seller", http.StatusFound)
				default:
					http.Redirect(w, r, "/", http.StatusFound)
				}
				return
			}

			Email := r.FormValue("email")
			Password := r.FormValue("password")

			if len(Email) == 0 || len(Password) == 0 {
				http.Redirect(w, r, "/login", http.StatusFound)
				return
			}

			userModel := model.NewUserModel()
			isLogin, err := userModel.ValidasiLogin(w, Email, Password)

			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			if !isLogin {
				http.Redirect(w, r, "/login", http.StatusFound)
				return
			}

			role, err := r.Cookie("role")
			if err != nil {
				http.Redirect(w, r, "/", http.StatusFound)
				return
			}

			switch role.Value {
			case "admin":
				http.Redirect(w, r, "/admin", http.StatusFound)
			case "seller":
				http.Redirect(w, r, "/seller", http.StatusFound)
			default:
				http.Redirect(w, r, "/", http.StatusFound)
			}

		} else if r.Method == "GET" || r.Method == "" {
			websiteModel := model.NewWebsiteModel()
			settings, err := websiteModel.GetSettings()

			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			templates := []string{
				filepath.Join("views", "templates.html"),
				filepath.Join("views", "autentikasi", "login.html"),
			}

			data := map[string]any{
				"Website": settings,
				"Cookies": cookie.GetAllCookies(r),
			}

			tmpl, err := template.ParseFiles(templates...)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			err = tmpl.ExecuteTemplate(w, "header", data)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			err = tmpl.ExecuteTemplate(w, "login", data)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			err = tmpl.ExecuteTemplate(w, "footer", data)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}
	}
}

func Register(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/register" {
			controller.ErrorHandler(w, r, http.StatusNotFound)
			return
		}
		if r.Method == "POST" {

			Username := r.FormValue("username")
			Nama := r.FormValue("nama")
			Email := r.FormValue("email")
			NoHP := r.FormValue("no_hp")
			Password := r.FormValue("password")
			ConfirmPassword := r.FormValue("confirmpassword")

			if Username == "" || Nama == "" || Email == "" || NoHP == "" || Password == "" || ConfirmPassword == "" {
				cookie.SetFlashCookie(w, "error", "Semua kolom wajib di isi!")
				http.Redirect(w, r, "register", http.StatusFound)
				return
			}

			if !isValidEmail(Email) {
				cookie.SetFlashCookie(w, "error", "Email tidak valid!")
				http.Redirect(w, r, "register", http.StatusFound)
				return
			}

			if Password != ConfirmPassword {
				cookie.SetFlashCookie(w, "error", "Password tidak sesuai")
				http.Redirect(w, r, "register", http.StatusFound)
				return
			}

			userModel := model.NewUserModel()
			_, msg := userModel.ValidasiRegister(Username, Nama, Email, NoHP, Password)
			if msg != "" {
				cookie.SetFlashCookie(w, "error", msg)
				http.Redirect(w, r, "register", http.StatusFound)
				return
			}

			cookie.SetFlashCookie(w, "sukses", "Pendaftaran berhasil! silahkan login")
			http.Redirect(w, r, "login", http.StatusFound)

		} else {

			templates := []string{
				filepath.Join("views", "templates.html"),
				filepath.Join("views", "autentikasi", "register.html"),
			}

			tmpl, err := template.ParseFiles(templates...)
			if err != nil {
				w.Write([]byte(err.Error()))
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			data := make(map[string]interface{})

			websiteModel := model.NewWebsiteModel()
			settings, err := websiteModel.GetSettings()

			if err != nil {
				w.Write([]byte(err.Error()))
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			data["Website"] = settings
			data["Cookies"] = cookie.GetAllCookies(r)

			err = tmpl.ExecuteTemplate(w, "header", data)
			if err != nil {
				w.Write([]byte(err.Error()))
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			err = tmpl.ExecuteTemplate(w, "register", data)
			if err != nil {
				w.Write([]byte(err.Error()))
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			err = tmpl.ExecuteTemplate(w, "footer", data)
			if err != nil {
				w.Write([]byte(err.Error()))
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

		}
	}
}

func Logout() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		CookiesName := []string{
			"isLogin",
			"username",
			"role",
		}

		cookie.DeleteBulkCookies(w, CookiesName)

		http.Redirect(w, r, "/", http.StatusFound)
	}
}
