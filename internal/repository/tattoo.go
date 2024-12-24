package repository

import (
	"context"
	"github.com/jmoiron/sqlx"
	models "tp-project/internal/models/tattoo"
	"tp-project/pkg/logger"
)

const (
	queryTattoo       = "SELECT id, name, artist, description, price FROM tattoos"
	queryCreateTattoo = "INSERT INTO tattoos (name, artist, description, price) VALUES ($1, $2, $3, $4) RETURNING id"
	queryUpdateTattoo = "UPDATE tattoos SET name = $1, artist = $2, description = $3, price = $4 WHERE id = $5"
	queryDeleteTattoo = "DELETE FROM tattoos WHERE id = $1"
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
