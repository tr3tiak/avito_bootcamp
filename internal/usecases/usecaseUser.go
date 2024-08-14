package usecases

import (
	"avito_bootcamp/internal/entity"
	"avito_bootcamp/pkg"
)

type UsecaseUser struct {
	repo entity.UserRepo
}

func InitUseCaseUser(ur entity.UserRepo) *UsecaseUser {
	return &UsecaseUser{
		repo: ur,
	}
}

func (uc UsecaseUser) Register(user *entity.User) error {
	encPass, err := pkg.EncryptedPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = encPass
	err = uc.repo.Create(user)
	if err != nil {
		return err
	}
	return nil
}

func (uc UsecaseUser) Login(user *entity.User) (string, error) {

	expectedUser, err := uc.repo.Get(user.Name)
	if err != nil {
		return "", err
	}
	err = pkg.ComparePassword(expectedUser.Password, user.Password)
	if err != nil {
		return "", err
	}
	tokenString, err := pkg.GenerateJWTToken(expectedUser)
	return tokenString, nil

}
