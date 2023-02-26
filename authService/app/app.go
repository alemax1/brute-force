package app

import (
	"authService/config"
	"authService/db"
	"authService/internal/delivery/route"
	"authService/internal/repository"
	"authService/internal/usecase"
	"fmt"

	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

func Run() {
	if err := config.InitCFG(); err != nil {
		log.Fatal().Err(err).Msg("error trying init config")
	}

	DB, err := db.CreateConnection()
	if err != nil {
		log.Fatal().Err(err).Msg("error trying create database connection")
	}

	userRepo := repository.New(DB)

	userUC := usecase.New(userRepo)

	e := route.CreateRoutes(userUC)

	e.Logger.Fatal(e.Start(fmt.Sprintf("%v:%d", viper.GetString("authService.host"), viper.GetInt("authService.port"))))
}
