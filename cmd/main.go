package main

import (
	"log"

	todo "github.com/Sancene/go-rest-learn"
	"github.com/Sancene/go-rest-learn/pkg/handler"
	"github.com/Sancene/go-rest-learn/pkg/repository"
	"github.com/Sancene/go-rest-learn/pkg/service"
)

func main() {

	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error has occured while running http server: %s", err.Error())
	}
}
