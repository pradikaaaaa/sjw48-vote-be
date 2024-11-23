package service

import (
	"be-awards/config"
	"be-awards/model"
	"be-awards/repository"
	"errors"
)

func VotingProcess(request model.VoteProcess) (status bool, err error) {
	tx := config.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	isFirstLoopSuccessful := true

	for _, vote := range request.IDKode {
		check, err := repository.CheckKodeCount(vote)
		if err != nil {
			tx.Rollback()
			return false, err
		}

		if check.Jumlah > 0 {
			_, err = repository.UpdateKodeVote(tx, vote)
			if err != nil {
				tx.Rollback()
				return false, err
			}
		} else {
			isFirstLoopSuccessful = false
		}

	}

	if !isFirstLoopSuccessful {
		return true, errors.New("kode serial voting telah digunakan")
	}

	for _, nominasi := range request.Vote {
		_, err = repository.UpdateJumlahVote(tx, nominasi)
		if err != nil {
			tx.Rollback()
			return false, err
		}
	}

	tx.Commit()
	if err != nil {
		return false, err
	}
	return true, err
}
