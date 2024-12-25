package repository

import (
	"context"
	"github.com/jmoiron/sqlx"
	models "tp-project/internal/models/tattoo"
	"tp-project/pkg/logger"
)

const (
	queryTattoo        = "SELECT id, name, artist, description, price FROM tattoos"
	queryCreateTattoo  = "INSERT INTO tattoos (name, artist, description, price) VALUES ($1, $2, $3, $4) RETURNING id"
	queryUpdateTattoo  = "UPDATE tattoos SET name = $1, artist = $2, description = $3, price = $4 WHERE id = $5"
	queryDeleteTattoo  = "DELETE FROM tattoos WHERE id = $1"
	queryGetAllTattoos = "SELECT * FROM tattoos ORDER BY id"
	queryGetTattooTags = "SELECT tag_id FROM tattoo_tags WHERE tattoo_id = $1"
	queryGetTagByID    = "SELECT name FROM tags WHERE id = $1"
	queryGetAllTags    = "SELECT name FROM tags"
)

type Tattoo struct {
	db     *sqlx.DB
	logger logger.Logger
}

func InitTattooRepo(db *sqlx.DB, logger logger.Logger) *Tattoo {
	return &Tattoo{
		db:     db,
		logger: logger,
	}
}

func (t Tattoo) GetTattooByID(ctx context.Context, id int) (*models.Tattoo, error) {
	var tattoo models.Tattoo
	err := t.db.GetContext(ctx, &tattoo, queryTattoo, id)
	if err != nil {
		return nil, err
	}
	return &tattoo, nil
}

func (t Tattoo) GetAllTattoos(ctx context.Context) ([]models.Tattoo, error) {
	var tattoos []models.Tattoo
	err := t.db.SelectContext(ctx, &tattoos, queryGetAllTattoos)
	if err != nil {
		return nil, err
	}
	return tattoos, nil
}

func (t Tattoo) CreateTattoo(ctx context.Context, tattoo models.Tattoo) (int, error) {
	var newID int
	err := t.db.GetContext(ctx, &newID, queryCreateTattoo, tattoo.Name, tattoo.Artist, tattoo.Description, tattoo.Price)
	if err != nil {
		return 0, err
	}
	return newID, nil
}

func (t Tattoo) UpdateTattoo(ctx context.Context, tattooId int, tattoo models.Tattoo) error {
	_, err := t.db.ExecContext(ctx, queryUpdateTattoo, tattoo.Name, tattoo.Artist, tattoo.Description, tattoo.Price, tattooId)
	return err
}

func (t Tattoo) DeleteTattoo(ctx context.Context, id int) error {
	_, err := t.db.ExecContext(ctx, queryDeleteTattoo, id)
	return err
}

func (t Tattoo) GetTattooTags(ctx context.Context, tattooId int) ([]int, error) {
	var tagIDs []int
	err := t.db.SelectContext(ctx, &tagIDs, queryGetTattooTags, tattooId)
	if err != nil {
		return nil, err
	}
	return tagIDs, nil
}

func (t Tattoo) GetTagByName(ctx context.Context, tagName string) (*string, error) {
	tn := &tagName
	err := t.db.GetContext(ctx, tn, queryGetTagByID, tagName)
	if err != nil {
		return nil, err
	}
	return tn, nil
}

func (t Tattoo) GetAllTags(ctx context.Context) ([]string, error) {
	var tags []string
	err := t.db.SelectContext(ctx, &tags, queryGetAllTags)
	if err != nil {
		return nil, err
	}
	return tags, nil
}

func (t Tattoo) GetTattoosByTag(ctx context.Context, tagID int) ([]models.Tattoo, error) {
	var tattoos []models.Tattoo
	err := t.db.SelectContext(ctx, &tattoos, queryTattoo, tagID)
	if err != nil {
		return nil, err
	}
	return tattoos, nil
}
