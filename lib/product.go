package lib

import "github.com/alf4ridzi/lapaksiswa-golang/model"

func GetProduct(ProductID string) (*model.Produk, error) {
	productModel := model.NewProdukModel()
	defer productModel.DB.Close()

	product, err := productModel.GetProductByOnlyID(ProductID)
	if err != nil {
		return nil, err
	}

	return product, nil
}
