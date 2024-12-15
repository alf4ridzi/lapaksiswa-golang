package cookie

import "net/http"

// expired dalam 1 hari dalam hitungan detik.
var CookieMaxAge = 86400

type Cookie struct {
	Cookies map[string]string
}

func SetCookie(w http.ResponseWriter, data map[string]string) {
	for name, value := range data {
		Cookie := &http.Cookie{
			Name:     name,
			Value:    value,
			Path:     "/",
			HttpOnly: false,
			Secure:   false,
			MaxAge:   CookieMaxAge,
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
		MaxAge:   2,
	}

	http.SetCookie(w, Cookie)
}

func GetAllCookies(r *http.Request) map[string]string {
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
