package service

import (
	"context"
	models "tp-project/internal/models/user"
)

type UserRepo interface {
	GetUserByID(ctx context.Context, id int) (*models.User, error)
	CreateUser(ctx context.Context, user models.User) (int, error)
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
	UpdateUser(ctx context.Context, user models.User) error
	DeleteUser(ctx context.Context, id int) error
	GetAllUsers(ctx context.Context) ([]models.User, error)
}

type UserService struct {
	UserRepo UserRepo
}

func NewUserService(userRepo UserRepo) UserService {
	return UserService{
		UserRepo: userRepo,
	}
}

func (s UserService) GetAllUsers(ctx context.Context) ([]models.User, error) {
	return s.UserRepo.GetAllUsers(ctx)
}

func (s UserService) CreateUser(ctx context.Context, user models.User) (int, error) {
	return s.UserRepo.CreateUser(ctx, user)
}

func (s UserService) GetUserByID(ctx context.Context, id int) (*models.User, error) {
	return s.UserRepo.GetUserByID(ctx, id)
}

func (s UserService) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	return s.UserRepo.GetUserByEmail(ctx, email)
}

func (s UserService) UpdateUser(ctx context.Context, user models.User) error {
	return s.UserRepo.UpdateUser(ctx, user)
}

func (s UserService) DeleteUser(ctx context.Context, id int) error {
	return s.UserRepo.DeleteUser(ctx, id)
}
