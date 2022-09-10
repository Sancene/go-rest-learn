package main

import (
	"log"

	todo "github.com/Sancene/go-rest-learn"
	"github.com/Sancene/go-rest-learn/pkg/handler"
	"github.com/Sancene/go-rest-learn/pkg/repository"
	"github.com/Sancene/go-rest-learn/pkg/service"
	"github.com/spf13/viper"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error has occured while running http server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("../configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
