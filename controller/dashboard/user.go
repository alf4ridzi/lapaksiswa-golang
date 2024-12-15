package dashboard

import (
	"database/sql"
	"net/http"
	"path/filepath"
	"text/template"

	"github.com/alf4ridzi/lapaksiswa-golang/config/cookie"
	"github.com/alf4ridzi/lapaksiswa-golang/controller"
	"github.com/alf4ridzi/lapaksiswa-golang/model"
)

func User(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/user" {
			controller.ErrorHandler(w, r, http.StatusPermanentRedirect)
			return
		}

		isLogin, err := r.Cookie("isLogin")
		if err != nil {
			cookie.SetFlashCookie(w, "error", "Login terlebih dahulu!")
			http.Redirect(w, r, "/login", http.StatusPermanentRedirect)
			return
		}

		if isLogin.Value != "true" {
			cookie.SetFlashCookie(w, "error", "Login terlebih dahulu!")
			http.Redirect(w, r, "/login", http.StatusPermanentRedirect)
			return
		}

		Role, err := r.Cookie("role")
		if err != nil {
			cookie.SetFlashCookie(w, "error", "Login terlebih dahulu!")
			http.Redirect(w, r, "/login", http.StatusPermanentRedirect)
			return
		}

		if Role.Value != "user" {
			http.Redirect(w, r, "/", http.StatusPermanentRedirect)
			return
		}

		UserName, err := r.Cookie("username")
		if err != nil {
			cookie.SetFlashCookie(w, "error", "Login terlebih dahulu!")
			http.Redirect(w, r, "/login", http.StatusPermanentRedirect)
			return
		}

		if UserName.Value == "" {
			http.Redirect(w, r, "/", http.StatusPermanentRedirect)
			return
		}

		templates := []string{
			filepath.Join("views", "templates.html"),
			filepath.Join("views", "dashboard", "user.html"),
		}

		tmpl, err := template.ParseFiles(templates...)
		if err != nil {
			w.Write([]byte(err.Error()))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		data := make(map[string]any)

		websiteModel := model.NewWebsiteModel()
		settings, err := websiteModel.GetSettings()

		if err != nil {
			w.Write([]byte(err.Error()))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		userModel := model.NewUserModel()
		User, err := userModel.GetUser(UserName.Value)

		if err != nil {
			w.Write([]byte(err.Error()))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		data["Website"] = settings
		data["Cookies"] = cookie.GetAllCookies(r)
		data["User"] = User

		err = tmpl.ExecuteTemplate(w, "header", data)
		if err != nil {
			w.Write([]byte(err.Error()))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		err = tmpl.ExecuteTemplate(w, "dashboard", data)
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
