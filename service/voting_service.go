package service

import (
	"be-awards/config"
	"be-awards/model"
	"be-awards/repository"
)

func VotingProcess(request model.VoteProcess) (status bool, err error) {
	tx := config.DB.Begin()

	for _, vote := range request.IDKode {
		_, err = repository.UpdateKodeVote(tx, vote)
		if err != nil {
			tx.Rollback()
			return false, err
		}
	}

	for _, nominasi := range request.Vote {
		_, err = repository.UpdateJumlahVote(tx, nominasi)
		if err != nil {
			tx.Rollback()
			return false, err
		}
	}

	tx.Commit()
	return true, err
}
