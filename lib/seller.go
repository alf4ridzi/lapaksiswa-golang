package lib

import (
	"errors"
	"math/rand"
	"net/http"
	"time"

	"github.com/alf4ridzi/lapaksiswa-golang/config/cookie"
	"github.com/alf4ridzi/lapaksiswa-golang/model"
)

func ValidateUserCookies(r *http.Request, roleExpected string) (bool, error) {
	if isLogin, err := cookie.GetCookieValue(r, "isLogin"); err != nil || isLogin != "true" {
		if err != nil {
			return false, err
		}

		return false, nil
	}

	if Role, err := cookie.GetCookieValue(r, "role"); err != nil || Role == "" {
		if err != nil {
			return false, err
		}

		return false, nil
	}

	Username, err := cookie.GetCookieValue(r, "username")
	if err != nil {
		return false, err
	}

	if Username == "" {
		return false, nil
	}

	userModel := model.NewUserModel()
	defer userModel.DB.Close()

	if isRole, err := userModel.IsRole(Username, roleExpected); err != nil || !isRole {
		if err != nil {
			return false, err
		}

		return false, nil
	}

	return true, nil
}

func ValidateProduct(Domain string, ProductID string) bool {
	productModel := model.NewProdukModel()
	defer productModel.DB.Close()

	if product, err := productModel.GetProductByID(ProductID, Domain); err != nil || product == nil {
		return false
	}

	return true
}

func GetUserDomain(Username string) (string, error) {
	tokoModel := model.NewTokoModel()
	defer tokoModel.DB.Close()

	Toko, err := tokoModel.GetTokoByUsername(Username)

	if err != nil {
		return "", err
	}

	return Toko.Domain, nil
}

func DeleteProduct(Domain string, ProductID string) error {
	if !ValidateProduct(Domain, ProductID) {
		return errors.New("invalid Shop")
	}

	productModel := model.NewProdukModel()
	defer productModel.DB.Close()

	isDelete, err := productModel.DeleteProduct(Domain, ProductID)
	if err != nil {
		return err
	}

	if !isDelete {
		return errors.New("gagal menghapus product")
	}

	return nil
}

func UpdateToko(Domain string, Toko model.EditToko) error {
	tokoModel := model.NewTokoModel()
	defer tokoModel.DB.Close()

	if err := tokoModel.UpdateToko(Domain, Toko); err != nil {
		return err
	}

	return nil
}

func GenerateRandomStr(length int) string {
	charset := "ABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	var seed = rand.New(rand.NewSource(time.Now().UnixNano()))

	randomStr := make([]byte, length)
	for i := range randomStr {
		randomStr[i] = charset[seed.Intn(len(charset))]
	}

	return string(randomStr)
}

func UpdateLogoToko(Domain string, Filename string) error {
	tokoModel := model.NewTokoModel()
	defer tokoModel.DB.Close()

	if err := tokoModel.UpdateLogoToko(Domain, Filename); err != nil {
		return err
	}

	return nil
}

func IsValidProductID(ProductID string) (bool, error) {
	productModel := model.NewProdukModel()
	defer productModel.DB.Close()

	isValid, err := productModel.IsProdukID(ProductID)
	if err != nil {
		return false, err
	}

	return isValid, err
}

func InsertIntoCart(Username string, ProductID string) error {
	cartModel := model.NewCartModel()
	defer cartModel.DB.Close()

	if err := cartModel.AddToCart(Username, ProductID); err != nil {
		return err
	}

	return nil
}
