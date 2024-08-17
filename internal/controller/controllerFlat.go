package controller

import (
	"avito_bootcamp/internal/entity"
	"encoding/json"
	"net/http"

	"github.com/sirupsen/logrus"
)

type ControllerFlat struct {
	ucFlat entity.UsecaseFlatInterface
}

func InitControllerFlat(uc entity.UsecaseFlatInterface) *ControllerFlat {
	return &ControllerFlat{
		ucFlat: uc,
	}
}

func (cf ControllerFlat) HandlerCreateFlat(w http.ResponseWriter, r *http.Request) {
	logrus.Info("HandlerCreateFlat started")
	var flat entity.Flat
	err := json.NewDecoder(r.Body).Decode(&flat)
	if err != nil {
		logrus.Error(err)
		return
	}
	err = cf.ucFlat.CreateFlat(&flat)
	if err != nil {
		logrus.Error(err)
		return
	}
	err = json.NewEncoder(w).Encode(flat)
	if err != nil {
		logrus.Error(err)
		return
	}
	w.WriteHeader(http.StatusOK)
	logrus.Info("HandlerCreateFlat complete")
}

func (cf ControllerFlat) HandlerUpdateStatus(w http.ResponseWriter, r *http.Request) {
	logrus.Info("HandlerUpdateStatus started")
	var flat entity.Flat
	err := json.NewDecoder(r.Body).Decode(&flat)
	if err != nil {
		return
	}
	err = cf.ucFlat.UpdateStatusFlat(&flat)
	if err != nil {
		return
	}
	err = json.NewEncoder(w).Encode(flat)
	if err != nil {
		return
	}
	w.WriteHeader(http.StatusOK)
	logrus.Info("HandlerUpdateStatus complete")
}
