package dashboard

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"text/template"
	"time"

	"github.com/alf4ridzi/lapaksiswa-golang/config"
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

		username, err := r.Cookie("username")
		if err != nil || username.Value == "" {
			cookie.SetFlashCookie(w, "error", "Login terlebih dahulu!")
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}

		templates := []string{
			filepath.Join("views", "templates.html"),
			filepath.Join("views", "dashboard", "user", "user.html"),
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
			"Cookies": cookie.GetAllCookies(w, r),
			"User":    user,
			"Error":   cookie.GetFlashCookies(w, r, "error"),
			"Sukses":  cookie.GetFlashCookies(w, r, "sukses"),
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
				filepath.Join("views", "dashboard", "user", "edit.html"),
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
			data["Cookies"] = cookie.GetAllCookies(w, r)
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

		err = json.NewDecoder(r.Body).Decode(&profile)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		userModel := model.NewUserModel()
		w.Header().Set("Content-Type", "application/json")
		// update profile
		if err = userModel.UpdateProfileUser(Username.Value, profile.Nama, profile.TanggalLahir, profile.JenisKelamin, profile.Email, profile.NoHP); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Server Error"))
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Update Profile Berhasil!."))
	}
}

func GenerateRandomStr() string {
	rand.Seed(time.Now().UnixNano())
	length := 10

	ran_str := make([]byte, length)

	for i := 0; i < length; i++ {
		ran_str[i] = byte(65 + rand.Intn(26))
	}

	str := string(ran_str)
	return str
}

func UpdateProfile() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/update-profile" || r.Method != "POST" {
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

		if err := r.ParseMultipartForm(1024); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Gambar terlalu besar!"))
			return
		}

		file, handler, err := r.FormFile("profile_picture")
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Server Error : " + err.Error()))
			return
		}

		defer file.Close()

		dir, err := os.Getwd()
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Server Error : " + err.Error()))
			return
		}

		// get filename
		filename := handler.Filename
		ext := filepath.Ext(filename)

		extAllowed := []string{
			".jpg", ".png", ".jpeg", ".webp",
		}

		Allowed := false

		for _, extAl := range extAllowed {
			if ext == extAl {
				Allowed = true
			}
		}

		if !Allowed {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Ext bukan gambar!"))
			return
		}

		fileName := GenerateRandomStr()
		filename = fmt.Sprintf("%s%s", fileName, ext)

		fileLocation := filepath.Join(dir, "public/assets/picture/user", filename)
		targetFile, err := os.OpenFile(fileLocation, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Server Error : " + err.Error()))
			return
		}

		defer targetFile.Close()

		if _, err := io.Copy(targetFile, file); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Server Error : " + err.Error()))
			return
		}
		modelUser := model.NewUserModel()

		// hapus piccture lama
		query := "SELECT foto FROM user WHERE username = ?"
		oldPicture := modelUser.DB.QueryRow(query, Username.Value)
		var old string

		if err = oldPicture.Scan(
			&old,
		); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Server Error : " + err.Error()))
			return
		}

		if old != "default.png" {
			fileLama := filepath.Join(dir, "public/assets/picture/user", old)
			if err = os.Remove(fileLama); err != nil {
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte("Server Error : " + err.Error()))
				return
			}
		}

		// update picture baru
		if err = modelUser.UpdatePicture(Username.Value, filename); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Server Error : " + err.Error()))
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(config.GetUrl() + "/public/assets/picture/user/" + filename))
	}
}
