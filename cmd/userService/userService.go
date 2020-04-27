package main

import (
	"log"

	userhttp "github.com/Toringol/nonlinearity/app/user/delivery/http"
	"github.com/Toringol/nonlinearity/app/user/repository"
	"github.com/Toringol/nonlinearity/app/user/usecase"
	"github.com/Toringol/nonlinearity/config"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/spf13/viper"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	if err := config.Init(); err != nil {
		log.Fatalf("%s", err.Error())
	}

	listenAddr := viper.GetString("listenAddr")

	e := echo.New()

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${time_rfc3339} [${method}] ${remote_ip}, ${uri} ${status} 'error':'${error}'\n",
	}))

	userhttp.NewUserHandler(e, usecase.NewUserUsecase(repository.NewUserMemoryRepository()))

	e.Logger.Fatal(e.Start(listenAddr))
}
