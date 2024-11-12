package repository

import (
	"be-awards/config"
	"be-awards/model"
)

func GetAllNominasi() (response []model.Nominasi, err error) {
	db := config.DB

	query := db.Table(`nominasi`).
		Select(`id "id", nama "nama", logo "logo"`).
		Where(`view = 1`).Order(`nama asc`)

	err = query.Scan(&response).Error
	if err != nil {
		return response, err
	}

	return response, err
}
