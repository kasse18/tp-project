package repository

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
	"tp-project/internal/models/user"
	"tp-project/pkg/logger"
)

const (
	queryUsers = "SELECT id, login, password, email FROM users"
	queryUser  = "SELECT id, login, password, email FROM users WHERE id = $1"
	queryLogin = "SELECT id, login, password, email FROM users WHERE login = $1"

	queryCreateUser = "INSERT INTO users (login, password, email) VALUES ($1, $2, $3) RETURNING id"
	queryUpdateUser = "UPDATE users SET login = $1, password = $2, email = $3 WHERE id = $4"
	queryDeleteUser = "DELETE FROM users WHERE id = $1"

	queryGetUserByEmail = "SELECT id, login, password, email FROM users WHERE email = $1"
)

type User struct {
	db     *sqlx.DB
	logger logger.Logger
}

func (u *User) GetAllUsers(ctx context.Context) ([]models.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u *User) CreateUser(ctx context.Context, user models.User) (int, error) {
	//TODO implement me
	panic("implement me")
}

func (u *User) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u *User) UpdateUser(ctx context.Context, user models.User) error {
	//TODO implement me
	panic("implement me")
}

func (u *User) DeleteUser(ctx context.Context, id int) error {
	//TODO implement me
	panic("implement me")
}

func InitUserRepo(db *sqlx.DB, logger logger.Logger) *User {
	return &User{
		db:     db,
		logger: logger,
	}
}

func (u *User) GetUserByID(ctx context.Context, id int) (*models.User, error) {
	u.logger.Info(ctx, "Starting GetAll operation")
	out := models.User{}

	if err := u.db.PingContext(ctx); err != nil {
		u.logger.Error(ctx, "Failed to connect to database", zap.Error(err))
		return nil, fmt.Errorf("database connection failed: %w", err)
	}

	row, err := u.db.QueryContext(ctx, queryUser)
	if err != nil {
		u.logger.Error(ctx, "Failed to execute query", zap.Error(err))
		return nil, err
	}

	if err := row.Scan(&out.ID, &out.Login, &out.Password, &out.Email); err != nil {
		u.logger.Error(ctx, "failed to scan row", zap.Error(err))
	}
	return &out, nil
}
