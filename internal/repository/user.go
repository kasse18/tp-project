package repository

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	"tp-project/internal/models/user"
	"tp-project/pkg/logger"
)

const (
	queryGetUsers          = "SELECT id, login, password, email FROM users"
	queryGetUser           = "SELECT id, login, password, email FROM users WHERE id = $1"
	queryGetByLogin        = "SELECT id, login, password, email FROM users WHERE login = $1"
	queryCreateUser        = "INSERT INTO users (username, email, password, role) VALUES ($1, $2, $3, $4) RETURNING id"
	queryGetUserByID       = "SELECT id, username, email, role, created_at FROM users WHERE id = $1"
	queryUpdateUser        = "UPDATE users SET username = $1, email = $2, password = $3, role = $4 WHERE id = $5"
	queryDeleteUser        = "DELETE FROM users WHERE id = $1"
	queryGetUserByUsername = "SELECT id, username, email, role, created_at FROM users WHERE username = $1"
	queryGetUserByEmail    = "SELECT id, username, email, role, created_at FROM users WHERE email = $2"
)

type User struct {
	db     *sqlx.DB
	logger logger.Logger
}

func InitUserRepo(db *sqlx.DB, logger logger.Logger) *User {
	return &User{
		db:     db,
		logger: logger,
	}
}

func (u User) GetAllUsers(ctx context.Context) ([]models.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u User) CreateUser(ctx context.Context, user models.User) (int, error) {
	var newUserID int
	err := u.db.GetContext(ctx, &newUserID, queryCreateUser, user.Username, user.Email, user.Password, user.Role)
	if err != nil {
		return 0, fmt.Errorf("failed to create user: %w", err)
	}
	return newUserID, nil
}

func (u User) GetUserByID(ctx context.Context, userID int) (*models.User, error) {
	var user models.User
	err := u.db.GetContext(ctx, &user, queryGetUserByID, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get user by ID: %w", err)
	}
	return &user, nil
}

func (u User) UpdateUser(ctx context.Context, user models.User) error {
	result, err := u.db.ExecContext(ctx, queryUpdateUser, user.Username, user.Email, user.Password, user.Role, user.ID)
	if err != nil {
		return fmt.Errorf("failed to update user: %w", err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected during user deletion: %w", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("no rows affected during user update")
	}
	return nil
}

func (u User) DeleteUser(ctx context.Context, userID int) error {
	result, err := u.db.ExecContext(ctx, queryDeleteUser, userID)
	if err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected during user deletion: %w", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("no rows affected during user deletion")
	}
	return nil
}

func (u User) GetUserByUsername(ctx context.Context, username string) (*models.User, error) {
	var user models.User
	err := u.db.GetContext(ctx, &user, queryGetUserByUsername, username)
	if err != nil {
		return nil, fmt.Errorf("failed to get user by username: %w", err)
	}
	return &user, nil
}

func (u User) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User
	err := u.db.GetContext(ctx, &user, queryGetUserByEmail, email)
	if err != nil {
		return nil, fmt.Errorf("failed to get user by email: %w", err)
	}
	return &user, nil
}

func (u User) GetUser(ctx context.Context, userID int) (*models.User, error) {
	var user models.User
	err := u.db.GetContext(ctx, &user, queryGetUser, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get user by ID: %w", err)
	}
	return &user, nil
}

func (u User) Login(ctx context.Context, login string) (*models.User, error) {
	var user models.User
	err := u.db.GetContext(ctx, &user, queryGetByLogin, login)
	if err != nil {
		return nil, fmt.Errorf("failed to get user by login: %w", err)
	}
	return &user, nil
}

func (u User) GetUsers(ctx context.Context) ([]models.User, error) {
	var users []models.User
	err := u.db.SelectContext(ctx, &users, queryGetUsers)
	if err != nil {
		return nil, fmt.Errorf("failed to get users: %w", err)
	}
	return users, nil
}
