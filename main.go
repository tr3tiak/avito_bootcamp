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
	usecaseUser := usecases.InitUseCaseUser(repoUser)
	usecaseHouse := usecases.InitUseCaseHouse(repoHouse)
	ControllerUser := controller.InitController(usecaseUser)
	ControllerHouse := controller.InitControllerHouse(usecaseHouse)
	router := mux.NewRouter()
	router.HandleFunc("/register", ControllerUser.HandlerRegister)
	router.HandleFunc("/login", ControllerUser.HandlerLogin)
	router.HandleFunc("/house/create", middleware.AuthMiddleware(middleware.AccessMiddleware(ControllerHouse.HandlerCreateHouse)))
	logrus.Info("starting server")
	http.ListenAndServe("localhost:8080", router)

}
