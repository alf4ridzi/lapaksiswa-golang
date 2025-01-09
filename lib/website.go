package lib

import "github.com/alf4ridzi/lapaksiswa-golang/model"

func GetSettingsWebsite() (*model.Website, error) {
	websiteModel := model.NewWebsiteModel()
	defer websiteModel.DB.Close()

	settings, err := websiteModel.GetSettings()

	if err != nil {
		return nil, err
	}

	return settings, nil
}
