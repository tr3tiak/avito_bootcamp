package main

import (
	"avito_bootcamp/internal/controller"
	"avito_bootcamp/internal/controller/middleware"
	"avito_bootcamp/internal/repo"
	"avito_bootcamp/internal/usecases"
	"net/http"

	"github.com/sirupsen/logrus"

	"github.com/gorilla/mux"
)

func main() {
	repoUser, err := repo.InitRepo()
	if err != nil {
		return
	}
	repoHouse, err := repo.InitHouseRepo()
	if err != nil {
		return
	}
	repoFlat, err := repo.InitFlatRepo()
	if err != nil {
		return
	}
	usecaseUser := usecases.InitUseCaseUser(repoUser)
	usecaseHouse := usecases.InitUseCaseHouse(repoHouse)
	usecaseFlat := usecases.InitUseCaseFlat(repoFlat)
	ControllerUser := controller.InitController(usecaseUser)
	ControllerHouse := controller.InitControllerHouse(usecaseHouse)
	ControllerFlat := controller.InitControllerFlat(usecaseFlat)
	router := mux.NewRouter()
	router.HandleFunc("/register", ControllerUser.HandlerRegister)
	router.HandleFunc("/login", ControllerUser.HandlerLogin)
	router.HandleFunc("/house/create", middleware.AuthMiddleware(middleware.AccessMiddleware(ControllerHouse.HandlerCreateHouse)))
	router.HandleFunc("/flat/create", middleware.AuthMiddleware(ControllerFlat.HandlerCreateFlat))
	router.HandleFunc("/flat/update", middleware.AuthMiddleware(middleware.AccessMiddleware(ControllerFlat.HandlerUpdateStatus)))
	logrus.Info("starting server")
	http.ListenAndServe("localhost:8080", router)

}
