package main

import (
	"log"

	userhttp "github.com/Toringol/nonlinearity/app/profileService/user/delivery/http"
	"github.com/Toringol/nonlinearity/app/profileService/user/repository"
	"github.com/Toringol/nonlinearity/app/profileService/user/usecase"
	"github.com/Toringol/nonlinearity/config"
	"github.com/labstack/echo"
	"github.com/spf13/viper"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	if err := config.Init(); err != nil {
		log.Fatalf("%s", err.Error())
	}

	listenAddr := viper.GetString("listenAddr")
	databaseConfig := viper.GetString("databaseConfig")

	e := echo.New()

	userhttp.NewUserHandler(e, usecase.NewUserUsecase(repository.NewUserMemoryRepository(databaseConfig)))

	e.Logger.Fatal(e.Start(listenAddr))
}
