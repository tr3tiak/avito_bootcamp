package usecases

import (
	"avito_bootcamp/internal/entity"

	"github.com/sirupsen/logrus"
)

type UsecaseHouse struct {
	repo entity.HouseRepoInterface
}

func InitUseCaseHouse(hr entity.HouseRepoInterface) *UsecaseHouse {
	return &UsecaseHouse{
		repo: hr,
	}
}

func (uc UsecaseHouse) CreateHouse(house *entity.House) (*entity.House, error) {
	logrus.Info("usecase house create started")
	responseHouse, err := uc.repo.CreateHouse(house)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	logrus.Info("usecase house create complete")
	return responseHouse, nil
}
