package controller

import (
	"avito_bootcamp/internal/entity"
	"encoding/json"
	"net/http"
)

type ControllerUser struct {
	usecases entity.UserUsecase
}

func InitController(uc entity.UserUsecase) ControllerUser {
	return ControllerUser{
		usecases: uc,
	}
}

func (c *ControllerUser) HandlerRegister(w http.ResponseWriter, r *http.Request) {
	user := entity.User{}
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&user)
	err := c.usecases.Register(&user)
	if err != nil {
		return
	}

	w.WriteHeader(http.StatusOK)

}

func (c *ControllerUser) HandlerLogin(w http.ResponseWriter, r *http.Request) {
	user := entity.User{}
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&user)

	token, err := c.usecases.Login(&user)
	if err != nil {
		return
	}
	w.Header().Set("Authorization", "Bearer "+token)
	w.WriteHeader(http.StatusOK)
}
