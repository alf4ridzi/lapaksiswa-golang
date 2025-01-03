package cookie

import (
	"net/http"
)

// expired dalam 1 hari dalam hitungan detik.
var CookieMaxAge = 86400

// expired dalam 30 hari
const CookieMaxAgeMonth = 30 * 24 * 60 * 60

type Cookie struct {
	Cookies map[string]string
}

func SetCookie(w http.ResponseWriter, data map[string]string, expired int) {
	for name, value := range data {
		Cookie := &http.Cookie{
			Name:     name,
			Value:    value,
			Path:     "/",
			HttpOnly: false,
			Secure:   false,
			MaxAge:   expired,
		}

		http.SetCookie(w, Cookie)
	}
}

func SetFlashCookie(w http.ResponseWriter, name string, value string) {

	Cookie := &http.Cookie{
		Name:     name,
		Value:    value,
		Path:     "/",
		HttpOnly: false,
		Secure:   false,
		MaxAge:   3,
	}

	http.SetCookie(w, Cookie)
}

func GetAllAndClearFlashCookie(w http.ResponseWriter, r *http.Request, names []string) map[string]string {
	cookies := map[string]string{}

	for _, name := range names {
		cookie, err := r.Cookie(name)
		if err == nil {
			cookies[name] = cookie.Value
			DeleteCookie(w, name)
		}
	}

	return cookies
}

func GetFlashCookies(w http.ResponseWriter, r *http.Request, name string) string {
	cookie, err := r.Cookie(name)
	if err != nil {
		return ""
	}

	cookieVal := cookie.Value
	DeleteCookie(w, name)

	return cookieVal
}

func DeleteCookie(w http.ResponseWriter, name string) {
	Cookie := &http.Cookie{
		Name:     name,
		Value:    "",
		Path:     "/",
		HttpOnly: false,
		Secure:   false,
		MaxAge:   -1,
	}

	http.SetCookie(w, Cookie)
}

func GetAllCookies(w http.ResponseWriter, r *http.Request) map[string]string {
	cookies := make(map[string]string)
	for _, c := range r.Cookies() {
		cookies[c.Name] = c.Value
	}

	return cookies
}

func DeleteBulkCookies(w http.ResponseWriter, name []string) {
	for _, cookie := range name {
		cookies := &http.Cookie{
			Name:     cookie,
			Value:    "",
			Path:     "/",
			HttpOnly: false,
			Secure:   false,
			MaxAge:   -1,
		}

		http.SetCookie(w, cookies)
	}
}
