package app

import (
	"context"
	"test-hexagonal-go/internal/domain/ports"
	"test-hexagonal-go/internal/infra/adapters/models"
)

type User interface {
	GetAllUser(ctx context.Context) ([]models.User, error)
}

type user struct {
	userRepo ports.Repository
}

func NewUserApp(userRepo ports.Repository) User {
	return &user{userRepo}
}

func (app *user) GetAllUser(ctx context.Context) ([]models.User, error) {

	u, err := app.userRepo.GetAllUser(ctx)
	if err != nil {
		return nil, err
	}

	user := []models.User{}

	for _, u := range u {
		user = append(user, models.User{
			UserID:        u.UserID,
			Name:          u.Name,
			Username:      u.Username,
			Permission:    u.Permission,
			Status:        u.Status,
			Date_Register: u.Date_Register,
			Date_Update:   u.Date_Update,
			Registered_By: u.Registered_By,
		})
	}

	return user, nil

}
