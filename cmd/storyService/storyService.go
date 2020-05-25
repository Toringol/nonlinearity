package storyService

import (
	"log"

	storyhttp "github.com/Toringol/nonlinearity/app/story/delivery/http"

	"github.com/Toringol/nonlinearity/app/story/repository"
	"github.com/Toringol/nonlinearity/app/story/usecase"
	"github.com/Toringol/nonlinearity/config"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/spf13/viper"
)

func main() {
	if err := config.Init(); err != nil {
		log.Fatalf("%s", err.Error())
	}

	listenAddr := viper.GetString("listenAddrStoryService")

	e := echo.New()

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${time_rfc3339} [${method}] ${remote_ip}, ${uri} ${status} 'error':'${error}'\n",
	}))

	storyhttp.NewStoryHandler(e, usecase.NewStoryUsecase(repository.NewStoryMemoryRepository()))

	e.Logger.Fatal(e.Start(listenAddr))
}
