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
		reqVote := model.KodeVote{
			Kode:   utils.GenerateRandomCode(16),
			Jumlah: request.JumlahVote,
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

func GetJumlahVote(request model.RequestKodeVote) (response []model.KodeVote, err error) {

	if !utils.CheckDuplicates(request.KodeVote) {
		for _, vote := range request.KodeVote {
			data, err := repository.GetJumlahVote(vote)

			if err != nil {
				return response, err
			}

			if data.Jumlah == 0 {
				return response, errors.New("kode sudah terpakai")
			}

			response = append(response, data)
		}
	} else {
		return response, errors.New("duplikasi kode vote")
	}

	return response, err
}
