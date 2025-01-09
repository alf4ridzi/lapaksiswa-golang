package lib

import "github.com/alf4ridzi/lapaksiswa-golang/model"

func CreateCheckout(ProductID string, Checkout string, Total int64, Username string, Qty int64) error {
	checkoutModel := model.NewCheckoutModel()
	defer checkoutModel.DB.Close()

	if err := checkoutModel.InsertCheckout(ProductID, Checkout, Total, Username, Qty); err != nil {
		return err
	}

	return nil
}

func IsValidCheckout(CheckoutID string, Username string) (bool, error) {
	checkoutModel := model.NewCheckoutModel()
	defer checkoutModel.DB.Close()

	IsValid, err := checkoutModel.IsValidCheckout(CheckoutID, Username)
	if err != nil {
		return false, err
	}

	return IsValid, nil
}

func GetDetailCheckout(CheckoutID string, Username string) (*model.Checkout, error) {
	checkoutModel := model.NewCheckoutModel()
	defer checkoutModel.DB.Close()

	Detail, err := checkoutModel.GetDetailCheckout(CheckoutID, Username)
	if err != nil {
		return nil, err
	}

	return Detail, nil
}
