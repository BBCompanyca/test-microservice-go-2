package handler

import (
	"errors"
	"net/http"
	"test-hexagonal-go/internal/app"

	"github.com/labstack/echo/v4"
)

type User interface {
	GetAllUser(ctx echo.Context) error
}

type user struct {
	app app.User
}

func NewUserHandler(app app.User) User {
	return &user{app}
}

func (handler *user) GetAllUser(c echo.Context) error {

	ctx := c.Request().Context()

	u, err := handler.app.GetAllUser(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errors.New("unexpected error"))
	}

	return c.JSON(http.StatusOK, u)

}
