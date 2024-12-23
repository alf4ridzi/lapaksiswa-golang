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
			filepath.Join("views", "dashboard", "seller", "templates.html"),
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

		tokoModel := model.NewTokoModel()
		Toko, err := tokoModel.GetTokoByUsername(Username.Value)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		orderModel := model.NewOrderModel()
		TrxHariIni, err := orderModel.GetTransaksiHariIni(Toko.Domain)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		TrxKemarin, err := orderModel.GetTransaksiKemarin(Toko.Domain)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		Omset, err := orderModel.GetTotalOmset(Toko.Domain)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		OmsetBulanan, err := orderModel.GetOmsetBulanan(Toko.Domain)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		TotalTrx, err := orderModel.GetTotalTransaksi(Toko.Domain)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		data["TotalTrx"] = TotalTrx
		data["OmsetBulanan"] = OmsetBulanan
		data["Omset"] = Omset
		data["Kemarin"] = TrxKemarin
		data["HariIni"] = TrxHariIni
		data["Website"] = settings
		data["Toko"] = Toko

		t := []string{
			"header", "navbar", "content", "end",
		}

		for _, tl := range t {
			if err = tmpl.ExecuteTemplate(w, tl, data); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(err.Error()))
				return
			}
		}
	}
}
