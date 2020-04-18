package main

import (
	userhttp "github.com/Toringol/nonlinearity/app/profileService/user/delivery/http"
	"github.com/Toringol/nonlinearity/app/profileService/user/repository"
	"github.com/Toringol/nonlinearity/app/profileService/user/usecase"
	"github.com/labstack/echo"
)

const listenAddr = "127.0.0.1:8080"

func main() {
	e := echo.New()

	userhttp.NewUserHandler(e, usecase.NewUserUsecase(repository.NewUserMemoryRepository()))

	e.Logger.Warnf("start listening on %s", listenAddr)
	err := e.Start("127.0.0.1:8080")
	if err != nil {
		e.Logger.Errorf("server error: %s", err)
	}

	e.Logger.Warnf("shutdown")
}
