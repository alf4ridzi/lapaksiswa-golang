package lib

import "github.com/alf4ridzi/lapaksiswa-golang/model"

func GetMetodePembayaran() ([]model.Pembayaran, error) {
	pemModel := model.NewPembayaranModel()
	defer pemModel.DB.Close()

	Pembayaran, err := pemModel.GetPembayaran()
	if err != nil {
		return nil, err
	}

	return Pembayaran, nil
}
