package controller

import (
	"avito_bootcamp/internal/entity"
	"encoding/json"
	"net/http"

	"github.com/sirupsen/logrus"
)

type ControllerHouse struct {
	usecases entity.HouseUseCase
}

func InitControllerHouse(uc entity.HouseUseCase) *ControllerHouse {
	return &ControllerHouse{
		usecases: uc,
	}
}

func (c *ControllerHouse) HandlerCreateHouse(w http.ResponseWriter, r *http.Request) {
	logrus.Info("controller house started")
	var house entity.House
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&house)
	responseHouse, err := c.usecases.CreateHouse(&house)
	if err != nil {
		return
	}
	w.Header().Set("Content-type", "application/json")
	encoder := json.NewEncoder(w)
	encoder.Encode(responseHouse)
	logrus.Info("controller house complete")
}
