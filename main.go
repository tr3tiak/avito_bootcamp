package main

import (
	"avito_bootcamp/internal/controller"
	"avito_bootcamp/internal/repo"
	"avito_bootcamp/internal/usecases"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	repo, err := repo.InitRepo()
	if err != nil {
		return
	}
	usecase := usecases.InitUseCaseUser(repo)
	controller := controller.InitController(usecase)
	router := mux.NewRouter()
	router.HandleFunc("/register", controller.HandlerRegister)
	router.HandleFunc("/login", controller.HandlerLogin)
	http.ListenAndServe("localhost:8080", router)

}
