package service

import (
	"context"
	models "tp-project/internal/models/tattoo"
)

type TattooRepo interface {
	GetTattooByID(ctx context.Context, id int) (*models.Tattoo, error)
	GetAllTattoos(ctx context.Context) ([]models.Tattoo, error)
	CreateTattoo(ctx context.Context, tattoo models.Tattoo) (int, error)
	UpdateTattoo(ctx context.Context, tattooId int, tattoo models.Tattoo) error
	DeleteTattoo(ctx context.Context, id int) error
}

type TattooService struct {
	TattooRepo TattooRepo
}

func NewTattooService(tattooRepo TattooRepo) TattooService {
	return TattooService{
		TattooRepo: tattooRepo,
	}
}

func (s TattooService) GetAllTattoos(ctx context.Context) ([]models.Tattoo, error) {
	return s.TattooRepo.GetAllTattoos(ctx)
}

func (s TattooService) GetTattooByID(ctx context.Context, id int) (*models.Tattoo, error) {
	return s.TattooRepo.GetTattooByID(ctx, id)
}

func (s TattooService) CreateTattoo(ctx context.Context, tattoo models.Tattoo) (int, error) {
	return s.TattooRepo.CreateTattoo(ctx, tattoo)
}

func (s TattooService) UpdateTattoo(ctx context.Context, tattooId int, tattoo models.Tattoo) error {
	return s.TattooRepo.UpdateTattoo(ctx, tattooId, tattoo)
}

func (s TattooService) DeleteTattoo(ctx context.Context, id int) error {
	return s.TattooRepo.DeleteTattoo(ctx, id)
}
