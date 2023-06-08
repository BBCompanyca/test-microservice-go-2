package ports

import (
	"context"
	"test-hexagonal-go/internal/domain/entity"
)

type Repository interface {
	GetAllUser(ctx context.Context) ([]entity.User, error)
}
