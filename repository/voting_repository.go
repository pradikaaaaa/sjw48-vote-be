package repository

import (
	"be-awards/model"

	"gorm.io/gorm"
)

func UpdateKodeVote(tx *gorm.DB, id int) (status bool, err error) {
	query := "UPDATE kode_vote SET jumlah = 0 WHERE id = ?"

	err = tx.Exec(query, id).Error
	if err != nil {
		return false, err
	}

	return true, err
}

func UpdateJumlahVote(tx *gorm.DB, request model.NominasiVote) (status bool, err error) {
	query := "UPDATE nominasi SET total_vote = total_vote + ? WHERE id = ?"

	err = tx.Exec(query, request.JumlahVote, request.ID).Error
	if err != nil {
		return false, err
	}

	return true, err
}
