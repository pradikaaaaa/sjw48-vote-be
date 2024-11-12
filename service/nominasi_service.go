package service

import (
	"be-awards/model"
	"be-awards/repository"
)

func GetAllNominasi() (response []model.Nominasi, err error) {
	return repository.GetAllNominasi()
}
