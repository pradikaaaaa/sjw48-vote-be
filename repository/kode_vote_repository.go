package repository

import (
	"be-awards/config"
	"be-awards/model"

	"gorm.io/gorm"
)

func CreateVote(tx *gorm.DB, kodeVote model.KodeVote) (response model.KodeVote, err error) {
	return response, tx.Create(&kodeVote).Error
}

func CheckKodeExist(tx *gorm.DB, kode string) (bool, error) {
	var count int64
	err := tx.Model(&model.KodeVote{}).
		Where("kode = ?", kode).
		Count(&count).Error

	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func GetJumlahVote(kodeVote string) (response model.KodeVote, err error) {
	db := config.DB

	query := db.Table("kode_vote").
		Select(`id "id", kode "kode", jumlah "jumlah"`).
		Where("kode = ?", kodeVote)

	err = query.First(&response).Error
	return response, err
}

func CheckKodeCount(id int) (response model.KodeVote, err error) {
	db := config.DB

	query := db.Table("kode_vote").
		Select(`id "id", kode "kode", jumlah "jumlah"`).
		Where("id = ?", id)

	err = query.First(&response).Error
	return response, err
}
