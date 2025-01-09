package lib

import "github.com/alf4ridzi/lapaksiswa-golang/model"

func CreateCheckout(ProductID string, Checkout string, Total int64, Username string) error {
	checkoutModel := model.NewCheckoutModel()
	defer checkoutModel.DB.Close()

	if err := checkoutModel.InsertCheckout(ProductID, Checkout, Total, Username); err != nil {
		return err
	}

	return nil
}
