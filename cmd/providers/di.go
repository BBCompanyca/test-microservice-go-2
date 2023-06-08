package providers

import (
	"test-hexagonal-go/internal/app"
	"test-hexagonal-go/internal/infra/api/handler"
	"test-hexagonal-go/internal/infra/api/router"
	"test-hexagonal-go/internal/infra/api/router/groups"
	"test-hexagonal-go/internal/infra/repo"
	"test-hexagonal-go/internal/infra/resource"

	"github.com/labstack/echo/v4"
	"go.uber.org/dig"
)

var Container *dig.Container

func BuildContainer() *dig.Container {
	Container = dig.New()

	_ = Container.Provide(func() *echo.Echo {
		return echo.New()
	})

	_ = Container.Provide(resource.New)

	_ = Container.Provide(router.New)

	_ = Container.Provide(groups.NewUserGroup)

	_ = Container.Provide(handler.NewUserHandler)

	_ = Container.Provide(app.NewUserApp)

	_ = Container.Provide(repo.NewRepository)

	return Container
}
