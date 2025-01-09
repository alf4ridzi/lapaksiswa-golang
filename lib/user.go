package lib

import "github.com/alf4ridzi/lapaksiswa-golang/model"

func GetUserSpec(Username string, Col string) (string, error) {
	userModel := model.NewUserModel()
	defer userModel.DB.Close()

	Value, err := userModel.GetUserSpec(Username, Col)
	if err != nil {
		return "", nil
	}

	return Value, err
}
