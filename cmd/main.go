package main

import (
	"fmt"
	"log"

	"test-hexagonal-go/cmd/providers"
	"test-hexagonal-go/config"
	"test-hexagonal-go/internal/infra/api/router"

	"github.com/labstack/echo/v4"
)

var (
	serverHost = config.Environments().ServerHost
	serverPort = config.Environments().ServerPort
)

func main() {
	container := providers.BuildContainer()

	err := container.Invoke(func(router *router.Router, server *echo.Echo) {
		router.Init()
		server.Logger.Fatal(server.Start(fmt.Sprintf("%s:%d", "localhost", 3000)))
	})

	if err != nil {
		log.Panic(err)
	}
}
