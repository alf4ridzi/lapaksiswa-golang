package lib

import (
	"errors"
	"net/http"

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
