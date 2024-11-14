package service

import (
	"be-awards/config"
	"be-awards/model"
	"be-awards/repository"
	"errors"

	"be-awards/utils"
)

func AddKoteVote(request model.RequestVote) (response model.KodeVote, err error) {
	tx := config.DB.Begin()

	for i := 0; i < request.Kuota; i++ {
		var reqVote model.KodeVote
		for {
			kode := utils.GenerateRandomCode(16)

			exists, err := repository.CheckKodeExist(tx, kode)
			if err != nil {
				tx.Rollback()
				return response, err
			}

			if !exists {
				reqVote = model.KodeVote{
					Kode:   kode,
					Jumlah: request.JumlahVote,
				}
				break
			}
		}

		_, err = repository.CreateVote(tx, reqVote)
		if err != nil {
			tx.Rollback()
			return response, err
		}
	}

	tx.Commit()
	return response, err
}

func GetJumlahVote(request model.RequestKodeVote) (response []model.KodeVote, status bool, err error) {

	if !utils.CheckDuplicates(request.KodeVote) {
		for _, vote := range request.KodeVote {
			data, err := repository.GetJumlahVote(vote)

			if err != nil {
				if err.Error() == "record not found" {
					return response, true, errors.New("kode serial voting tidak sah")
				}
				return response, false, err
			}

			if data.Jumlah == 0 {
				return response, true, errors.New("kode serial voting telah digunakan")
			}

			response = append(response, data)
		}
	} else {
		return response, true, errors.New("duplikasi kode vote")
	}

	return response, true, err
}
