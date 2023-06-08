package router

import (
	"test-hexagonal-go/internal/infra/api/handler"
	"test-hexagonal-go/internal/infra/api/router/groups"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Router struct {
	server    *echo.Echo
	userGroup groups.User
}

func New(server *echo.Echo, userGroup groups.User) *Router {
	return &Router{
		server,
		userGroup,
	}
}

func (r *Router) Init() {

	r.server.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status} latency=${latency_human}\n",
	}))
	r.server.Use(middleware.Recover())

	basePath := r.server.Group("/api")
	basePath.GET("/health", handler.HealthCheck)

	r.userGroup.Resource(basePath)

}
