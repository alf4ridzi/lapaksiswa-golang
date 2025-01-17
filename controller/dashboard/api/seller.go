package api

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/alf4ridzi/lapaksiswa-golang/config/cookie"
	"github.com/alf4ridzi/lapaksiswa-golang/controller/dashboard"
	"github.com/alf4ridzi/lapaksiswa-golang/lib"
	"github.com/alf4ridzi/lapaksiswa-golang/model"
)

func ConvertMapToJson(data map[string]any) ([]byte, error) {
	Json, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	return Json, nil
}

func GenerateProdukID() string {
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	randomNumber := rng.Intn(1000000)
	return fmt.Sprintf("P%06d", randomNumber)
}

func GenerateSlug(ProductName string, ProductID string) string {
	lower := strings.ToLower(ProductName)
	slug := strings.ReplaceAll(lower, " ", "-")
	return fmt.Sprintf("%s-%s", slug, ProductID)
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

		tokoModel := model.NewTokoModel()
		defer tokoModel.DB.Close()

		Toko, err := tokoModel.GetTokoByUsername(Username.Value)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if Toko.Domain == "" {
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

		produkModel := model.NewProdukModel()
		defer produkModel.DB.Close()

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

func HandleImageUpload(ProdukID string, w http.ResponseWriter, r *http.Request) error {
	uploadedFile, handler, err := r.FormFile("gambar")

	if err != nil {
		return err
	}

	defer uploadedFile.Close()

	dir, err := os.Getwd()

	if err != nil {
		return err
	}

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
		return fmt.Errorf("ext yang di perbolehkan hanya gambar")
	}

	fileName := dashboard.GenerateRandomStr()
	filename = fmt.Sprintf("%s%s", fileName, ext)

	fileLocation := filepath.Join(dir, "public/img/product", filename)
	targetFile, err := os.OpenFile(fileLocation, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return err
	}

	defer targetFile.Close()

	if _, err := io.Copy(targetFile, uploadedFile); err != nil {
		return err
	}

	produkModel := model.NewProdukModel()
	defer produkModel.DB.Close()

	if err = produkModel.UpdateFotoProduk(ProdukID, filename); err != nil {
		return err
	}

	return nil
}

func TambahProduct() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		if r.Method != "POST" {
			w.WriteHeader(http.StatusForbidden)
			w.Write([]byte("Post doang oi"))
			return
		}

		data := map[string]any{
			"status": false,
			"result": nil,
		}

		w.Header().Set("Content-Type", "application/json")

		isLogin, err := r.Cookie("isLogin")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if err := r.ParseMultipartForm(1024); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if isLogin.Value != "true" {
			data["result"] = "Login dlu woilah"
			responseJson, err := ConvertMapToJson(data)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			w.WriteHeader(http.StatusInternalServerError)
			w.Write(responseJson)
			return
		}

		Username, err := r.Cookie("username")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if Username.Value == "" {
			data["result"] = "Login dlu woilah"
			responseJson, err := ConvertMapToJson(data)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			w.WriteHeader(http.StatusInternalServerError)
			w.Write(responseJson)
			return
		}

		tokoModel := model.NewTokoModel()
		defer tokoModel.DB.Close()

		Toko, err := tokoModel.GetTokoByUsername(Username.Value)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if Toko.Domain == "" {
			data["result"] = "Toko Tidak Ditemukan"
			responseJson, err := ConvertMapToJson(data)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			w.WriteHeader(http.StatusBadRequest)
			w.Write(responseJson)
			return
		}

		Harga, err := strconv.ParseInt(r.FormValue("harga"), 10, 64)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		Stok, err := strconv.ParseInt(r.FormValue("stok"), 10, 64)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		Product := model.Tambah{
			Nama:      r.FormValue("nama"),
			Deskripsi: r.FormValue("deskripsi"),
			Kategori:  r.FormValue("kategori"),
			Varian:    r.FormValue("varian"),
			Unit:      r.FormValue("unit"),
			Kondisi:   r.FormValue("kondisi"),
			Stok:      Stok,
			Harga:     Harga,
		}

		var ProdukID string
		produkModel := model.NewProdukModel()
		defer produkModel.DB.Close()

		for {
			ProdukID = GenerateProdukID()
			isProdukID, err := produkModel.IsProdukID(ProdukID)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			if !isProdukID {
				break
			}
		}

		Slug := GenerateSlug(Product.Nama, ProdukID)
		if err = produkModel.TambahProduk(ProdukID, Toko.Domain, Slug, Product); err != nil {
			data["result"] = "Gagal tambah produk. Internal Server Error " + err.Error()
			responseJson, err := ConvertMapToJson(data)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			w.WriteHeader(http.StatusInternalServerError)
			w.Write(responseJson)
			return
		}

		if err := HandleImageUpload(ProdukID, w, r); err != nil {
			data["result"] = "Gagal upload gambar"
			responseJson, err := ConvertMapToJson(data)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			w.WriteHeader(http.StatusInternalServerError)
			w.Write(responseJson)
			return
		}

		data["status"] = true
		responseJson, err := ConvertMapToJson(data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(responseJson)
	}
}

func HandleResponseJson(w http.ResponseWriter, data map[string]any, code int) {
	responseJson, err := ConvertMapToJson(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(code)
	w.Write(responseJson)
}

func UpdateStatusProduct() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		data := map[string]any{
			"result":  false,
			"message": nil,
		}

		w.Header().Set("Content-Type", "application/json")

		isLogin, err := r.Cookie("isLogin")
		if err != nil {
			data["message"] = "Login dulu cak"
			HandleResponseJson(w, data, http.StatusBadRequest)
			return
		}

		if isLogin.Value != "true" {
			data["message"] = "Login dulu cak"
			HandleResponseJson(w, data, http.StatusBadRequest)
			return
		}

		role, err := r.Cookie("role")
		if err != nil {
			data["message"] = "Internal server error"
			HandleResponseJson(w, data, http.StatusBadRequest)
			return
		}

		if role.Value != "seller" {
			data["message"] = "Khusus seller"
			HandleResponseJson(w, data, http.StatusBadRequest)
			return
		}

		Username, err := r.Cookie("username")
		if err != nil {
			data["message"] = "Username tidak ditemukan"
			HandleResponseJson(w, data, http.StatusBadRequest)
			return
		}

		if Username.Value == "" {
			data["message"] = "Username tidak ditemukan"
			HandleResponseJson(w, data, http.StatusBadRequest)
			return
		}

		ProductID := strings.TrimSpace(r.FormValue("product_id"))
		Status := strings.TrimSpace(r.FormValue("status"))
		if Status != "tersedia" && Status != "nonaktif" {
			data["message"] = "Invalid status"
			HandleResponseJson(w, data, http.StatusBadRequest)
			return
		}

		tokoModel := model.NewTokoModel()
		defer tokoModel.DB.Close()
		Toko, err := tokoModel.GetTokoByUsername(Username.Value)
		if err != nil {
			data["message"] = err.Error()
			HandleResponseJson(w, data, http.StatusInternalServerError)
			return
		}

		if Toko.Domain == "" {
			data["message"] = "Toko tidak ditemukan"
			HandleResponseJson(w, data, http.StatusInternalServerError)
			return
		}

		produkModel := model.NewProdukModel()
		defer produkModel.DB.Close()

		if isProduct, err := produkModel.ValidasiProduk(ProductID, Toko.Domain); err != nil || !isProduct {
			data["message"] = "Produk ID tidak valid / tidak ditemukan"
			if err != nil {
				data["message"] = err.Error()
			}
			HandleResponseJson(w, data, http.StatusInternalServerError)
			return
		}

		if changeStatus, err := produkModel.ChangeStatusProduct(ProductID, Toko.Domain, Status); err != nil || !changeStatus {
			data["message"] = "Gagal Update Status Produk"
			if err != nil {
				data["message"] = err.Error()
			}
			HandleResponseJson(w, data, http.StatusInternalServerError)
			return
		}

		data["status"] = true
		HandleResponseJson(w, data, http.StatusOK)
	}
}

func EditProduct() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		data := map[string]any{
			"status": false,
			"result": nil,
		}

		w.Header().Set("Content-Type", "application/json")

		isLogin, err := r.Cookie("isLogin")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if err := r.ParseMultipartForm(1024); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if isLogin.Value != "true" {
			data["result"] = "Login dlu woilah"
			responseJson, err := ConvertMapToJson(data)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			w.WriteHeader(http.StatusInternalServerError)
			w.Write(responseJson)
			return
		}

		Username, err := r.Cookie("username")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if Username.Value == "" {
			data["result"] = "Login dlu woilah"
			responseJson, err := ConvertMapToJson(data)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			w.WriteHeader(http.StatusInternalServerError)
			w.Write(responseJson)
			return
		}

		tokoModel := model.NewTokoModel()
		defer tokoModel.DB.Close()

		Toko, err := tokoModel.GetTokoByUsername(Username.Value)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if Toko.Domain == "" {
			data["result"] = "Toko Tidak Ditemukan"
			responseJson, err := ConvertMapToJson(data)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			w.WriteHeader(http.StatusBadRequest)
			w.Write(responseJson)
			return
		}

		Harga, err := strconv.ParseInt(r.FormValue("harga"), 10, 64)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		Stok, err := strconv.ParseInt(r.FormValue("stok"), 10, 64)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		ProductID := r.FormValue("productid")
		Product := model.EditProduct{
			ProductID: ProductID,
			Nama:      r.FormValue("nama"),
			Deskripsi: r.FormValue("deskripsi"),
			Kategori:  r.FormValue("kategori"),
			Varian:    r.FormValue("varian"),
			Unit:      r.FormValue("unit"),
			Kondisi:   r.FormValue("kondisi"),
			Stok:      Stok,
			Harga:     Harga,
		}

		produkModel := model.NewProdukModel()
		defer produkModel.DB.Close()

		if product, err := produkModel.GetProductByID(ProductID, Toko.Domain); err != nil || product == nil {
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			if product == nil {
				data["result"] = "Product tidak ditemukan / tidak valid"
				HandleResponseJson(w, data, http.StatusBadRequest)
				return
			}
		}

		uploadImage := true
		if _, _, err := r.FormFile("gambar"); err != nil {
			if err != http.ErrMissingFile {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			if err == http.ErrMissingFile {
				uploadImage = false
			}
		}

		if uploadImage {
			if err := HandleImageUpload(ProductID, w, r); err != nil {
				data["result"] = "Gagal upload gambar"
				HandleResponseJson(w, data, http.StatusInternalServerError)
				return
			}
		}

		if err := produkModel.HandleEditProduct(Toko.Domain, Product); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		data["status"] = true
		HandleResponseJson(w, data, http.StatusOK)
	}
}

func DeleteProduct() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		data := map[string]any{
			"status": false,
			"result": nil,
		}

		ProductID := r.FormValue("productid")

		w.Header().Set("Content-Type", "application/json")

		if isCookie, err := lib.ValidateUserCookies(r, "seller"); err != nil || !isCookie {
			data["result"] = "Cookie tidak valid"
			if err != nil {
				data["result"] = err.Error()
			}
			HandleResponseJson(w, data, http.StatusInternalServerError)
			return
		}

		Username, err := cookie.GetCookieValue(r, "username")
		if err != nil {
			data["result"] = err.Error()
			HandleResponseJson(w, data, http.StatusInternalServerError)
			return
		}
		// get user domain
		TokoDomain, err := lib.GetUserDomain(Username)
		if err != nil {
			data["result"] = err.Error()
			HandleResponseJson(w, data, http.StatusInternalServerError)
			return
		}

		if err = lib.DeleteProduct(TokoDomain, ProductID); err != nil {
			data["result"] = err.Error()
			HandleResponseJson(w, data, http.StatusBadRequest)
			return
		}

		data["status"] = true
		HandleResponseJson(w, data, http.StatusOK)
	}
}

func UpdateToko() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		response := map[string]any{
			"status": false,
			"result": nil,
		}

		if isValid, err := lib.ValidateUserCookies(r, "seller"); err != nil || !isValid {
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			if !isValid {
				response["result"] = "Data user tidak valid"
				HandleResponseJson(w, response, http.StatusBadRequest)
				return
			}
		}

		Username, err := cookie.GetCookieValue(r, "username")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		Domain, err := lib.GetUserDomain(Username)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if Domain == "" {
			response["result"] = "Hayo kamu bukan seller ngapain ?"
			HandleResponseJson(w, response, http.StatusBadRequest)
			return
		}

		var Toko model.EditToko
		if err = json.NewDecoder(r.Body).Decode(&Toko); err != nil {
			response["result"] = err.Error()
			HandleResponseJson(w, response, http.StatusBadRequest)
			return
		}

		if err = lib.UpdateToko(Domain, Toko); err != nil {
			response["result"] = "Gagal update : " + err.Error()
			HandleResponseJson(w, response, http.StatusBadRequest)
			return
		}

		response["status"] = true
		HandleResponseJson(w, response, http.StatusOK)
	}
}

func UpdateFotoToko() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		response := map[string]any{
			"status": false,
			"result": nil,
		}

		if isValid, err := lib.ValidateUserCookies(r, "seller"); err != nil || !isValid {
			if err != nil {
				response["result"] = err.Error()
				HandleResponseJson(w, response, http.StatusInternalServerError)
				return
			}

			response["result"] = "Autentikasi tidak valid."
			HandleResponseJson(w, response, http.StatusBadRequest)
			return
		}

		Username, err := cookie.GetCookieValue(r, "username")
		if err != nil {
			response["result"] = err.Error()
			HandleResponseJson(w, response, http.StatusInternalServerError)
			return
		}

		Domain, err := lib.GetUserDomain(Username)
		if err != nil {
			response["result"] = err.Error()
			HandleResponseJson(w, response, http.StatusInternalServerError)
			return
		}

		if err := r.ParseMultipartForm(1024); err != nil {
			response["result"] = "Size terlalu besar"
			HandleResponseJson(w, response, http.StatusBadRequest)
			return
		}

		imageUpload, handler, err := r.FormFile("image")
		if err != nil {
			response["result"] = err.Error()
			HandleResponseJson(w, response, http.StatusInternalServerError)
			return
		}

		defer imageUpload.Close()

		dir, err := os.Getwd()
		if err != nil {
			response["result"] = err.Error()
			HandleResponseJson(w, response, http.StatusInternalServerError)
			return
		}

		ext := filepath.Ext(handler.Filename)

		if !lib.IsExtAllowed(ext) {
			response["result"] = "Ext hanya gambar!"
			HandleResponseJson(w, response, http.StatusInternalServerError)
			return
		}

		randFilename := lib.GenerateRandomStr(10)
		filename := fmt.Sprintf("%s%s", randFilename, ext)

		fileLocation := filepath.Join(dir, "public", "img", "market", filename)

		targetFile, err := os.OpenFile(fileLocation, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			response["result"] = err.Error()
			HandleResponseJson(w, response, http.StatusInternalServerError)
			return
		}

		defer targetFile.Close()

		if _, err := io.Copy(targetFile, imageUpload); err != nil {
			response["result"] = err.Error()
			HandleResponseJson(w, response, http.StatusInternalServerError)
			return
		}

		if err = lib.UpdateLogoToko(Domain, filename); err != nil {
			response["result"] = err.Error()
			HandleResponseJson(w, response, http.StatusInternalServerError)
			return
		}

		response["status"] = true
		HandleResponseJson(w, response, http.StatusOK)
	}
}
