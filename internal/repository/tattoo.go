package repository

import (
	"context"
	"github.com/jmoiron/sqlx"
	models "tp-project/internal/models/tattoo"
	"tp-project/pkg/logger"
)

type Tattoo struct {
	db     *sqlx.DB
	logger logger.Logger
}

func (t Tattoo) GetTattooByID(ctx context.Context, id int) (*models.Tattoo, error) {
	//TODO implement me
	panic("implement me")
}

func (t Tattoo) GetAllTattoos(ctx context.Context) ([]models.Tattoo, error) {
	//TODO implement me
	panic("implement me")
}

func (t Tattoo) CreateTattoo(ctx context.Context, tattoo models.Tattoo) (int, error) {
	//TODO implement me
	panic("implement me")
}

func (t Tattoo) UpdateTattoo(ctx context.Context, tattooId int, tattoo models.Tattoo) error {
	//TODO implement me
	panic("implement me")
}

func (t Tattoo) DeleteTattoo(ctx context.Context, id int) error {
	//TODO implement me
	panic("implement me")
}

func InitTattooRepo(db *sqlx.DB, logger logger.Logger) *Tattoo {
	return &Tattoo{
		db:     db,
		logger: logger,
	}
}
