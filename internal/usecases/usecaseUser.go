package usecases

import (
	"avito_bootcamp/internal/entity"
	"avito_bootcamp/pkg"

	"github.com/sirupsen/logrus"
)

type UsecaseUser struct {
	repo entity.UserRepo
}

func InitUseCaseUser(ur entity.UserRepo) UsecaseUser {
	return UsecaseUser{
		repo: ur,
	}
}

func (uc UsecaseUser) Register(user *entity.User) error {
	logrus.Info("usecase register started")
	encPass, err := pkg.EncryptedPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = encPass
	err = uc.repo.Create(user)
	if err != nil {
		return err
	}
	logrus.Info("usecase register complete")
	return nil
}

func (uc UsecaseUser) Login(user *entity.User) (string, error) {
	logrus.Info("usecase login started")
	expectedUser, err := uc.repo.Get(user.Name)
	if err != nil {
		logrus.Error(err)
		return "", err
	}
	err = pkg.ComparePassword(expectedUser.Password, user.Password)
	if err != nil {
		return "", err
	}
	tokenString, err := pkg.GenerateJWTToken(expectedUser)
	if err != nil {
		logrus.Error("generate token error", err)
	}
	logrus.Info("usecase login complete")
	return tokenString, nil

}
