package repo

import (
	"context"
	"test-hexagonal-go/internal/domain/entity"
	"test-hexagonal-go/internal/domain/ports"

	"github.com/jmoiron/sqlx"
)

var (
	qryGetAllUser = `select userID, name, username, permission_level, status, date_register, date_update, registered_by from user`
)

type repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) ports.Repository {
	return repository{db}
}

func (repo repository) GetAllUser(ctx context.Context) ([]entity.User, error) {

	uu := []entity.User{}

	err := repo.db.SelectContext(ctx, &uu, qryGetAllUser)
	if err != nil {
		return nil, err
	}

	return uu, nil

}
