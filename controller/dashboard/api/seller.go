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

	"github.com/alf4ridzi/lapaksiswa-golang/controller/dashboard"
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
		return fmt.Errorf("Ext yang di perbolehkan hanya gambar")
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

		w.Header().Set("Content-Type", "application/json")

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

		Product := model.Tambah{
			Nama:      r.FormValue("nama"),
			Deskripsi: r.FormValue("deskripsi"),
			Kategori:  r.FormValue("kategori"),
			Varian:    r.FormValue("varian"),
			Unit:      r.FormValue("unit"),
			Kondisi:   r.FormValue("kondisi"),
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
		tambah, err := produkModel.TambahProduk(ProdukID, Toko.Domain, Slug, Product)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if !tambah {
			data["result"] = "Gagal tambah gambar"
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
