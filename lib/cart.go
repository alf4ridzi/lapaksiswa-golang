package lib

import "github.com/alf4ridzi/lapaksiswa-golang/model"

func GetUserCart(username string) ([]model.Produk, error) {
	productModel := model.NewProdukModel()
	defer productModel.DB.Close()

	listProduct, err := productModel.GetUserCart(username)
	if err != nil {
		return nil, err
	}

	return listProduct, nil
}
