package main

import (
	"log"

	todo "github.com/Sancene/go-rest-learn"
	"github.com/Sancene/go-rest-learn/pkg/handler"
)

func main() {
	srv := new(todo.Server)
	handlers := new(handler.Handler)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error has occured while running http server: %s", err.Error())
	}
}
