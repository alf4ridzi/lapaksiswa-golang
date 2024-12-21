package dashboard

import (
	"net/http"
	"path/filepath"
	"text/template"

	"github.com/alf4ridzi/lapaksiswa-golang/config/cookie"
	"github.com/alf4ridzi/lapaksiswa-golang/controller"
	"github.com/alf4ridzi/lapaksiswa-golang/model"
)

func Seller() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/seller" {
			controller.NotFoundHandler(w, r)
			return
		}

		isLogin, err := r.Cookie("isLogin")
		if err != nil {
			cookie.SetFlashCookie(w, "error", "Silahkan login terlebih dahulu!")
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}

		if isLogin.Value != "true" {
			cookie.SetFlashCookie(w, "error", "Silahkan login terlebih dahulu!")
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}

		Username, err := r.Cookie("username")
		if err != nil {
			cookie.SetFlashCookie(w, "error", "Silahkan login terlebih dahulu!")
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}

		if Username.Value == "" {
			cookie.SetFlashCookie(w, "error", "Silahkan login terlebih dahulu!")
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}

		// check role
		userModel := model.NewUserModel()
		isRole, err := userModel.IsRole(Username.Value, "seller")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		if !isRole {
			cookie.SetFlashCookie(w, "error", "Area khusus seller. Silahkan login sebagai seller")
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}

		templates := []string{
			filepath.Join("views", "dashboard", "seller", "index.html"),
		}

		tmpl, err := template.ParseFiles(templates...)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		data := make(map[string]any)

		websiteModel := model.NewWebsiteModel()
		settings, err := websiteModel.GetSettings()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		data["Website"] = settings

		if err = tmpl.Execute(w, data); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
	}
}
