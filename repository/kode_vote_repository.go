package repository

import (
	"be-awards/config"
	"be-awards/model"

	"gorm.io/gorm"
)

func CreateVote(tx *gorm.DB, kodeVote model.KodeVote) (response model.KodeVote, err error) {
	return response, tx.Create(&kodeVote).Error
}

func GetJumlahVote(kodeVote string) (response model.KodeVote, err error) {
	db := config.DB

	query := db.Table("kode_vote").
		Select(`id "id", kode "kode", jumlah "jumlah"`).
		Where("kode = ?", kodeVote)

	err = query.First(&response).Error
	return response, err
}
