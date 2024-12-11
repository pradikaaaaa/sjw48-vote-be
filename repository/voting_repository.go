package repository

import (
	"be-awards/lib"
	"be-awards/model"

	"gorm.io/gorm"
)

func UpdateKodeVote(tx *gorm.DB, id int) (status bool, err error) {
	timeNow := lib.GetTimeNow("timestime")
	query := "UPDATE kode_vote SET jumlah = 0, updated_at = ? WHERE id = ?"

	err = tx.Exec(query, timeNow, id).Error
	if err != nil {
		return false, err
	}

	return true, err
}

func UpdateJumlahVote(tx *gorm.DB, request model.NominasiVote) (status bool, err error) {
	timeNow := lib.GetTimeNow("timestime")
	query := "UPDATE nominasi SET total_vote = total_vote + ?, updated_at = ? WHERE id = ?"

	err = tx.Exec(query, request.JumlahVote, timeNow, request.ID).Error
	if err != nil {
		return false, err
	}

	return true, err
}
