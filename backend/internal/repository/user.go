package repository

/*
--- MIT License (c) 2024 achmad
--- See LICENSE for more details
*/

import (
	"context"

	"github.com/achmad/em/backend/internal/domain"
	"github.com/jmoiron/sqlx"
)

// base interface for user repository
type UserRepo interface {
	GetUserByID(ctx context.Context, id string) (*domain.User, error)
	GetUserByName(ctx context.Context, username string) (*domain.User, error)
	GetUsersByType(ctx context.Context, userType string) ([]*domain.UserCompanyResponse, error)
}

// user repository implementation
type userRepositoryImpl struct {
	sqlx *sqlx.DB
}

// GetUserByName implements UserRepo.
func (u *userRepositoryImpl) GetUserByName(ctx context.Context, username string) (*domain.User, error) {
	user := &domain.User{}
	query := `SELECT id, username, password, role, company_name, created_at, updated_at FROM users WHERE username = $1`
	err := u.sqlx.GetContext(ctx, user, query, username)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// GetUserByID implements UserRepo.
func (u *userRepositoryImpl) GetUserByID(ctx context.Context, id string) (*domain.User, error) {
	user := &domain.User{}
	query := `SELECT id, username, password, role, company_name, created_at, updated_at FROM users WHERE id = $1`
	err := u.sqlx.GetContext(ctx, user, query, id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// GetUsersByType implements User.
func (u *userRepositoryImpl) GetUsersByType(ctx context.Context, userType string) ([]*domain.UserCompanyResponse, error) {
	users := []*domain.UserCompanyResponse{}
	query := `SELECT company_name FROM users WHERE role = $1`
	err := u.sqlx.SelectContext(ctx, &users, query, userType)
	if err != nil {
		return nil, err
	}
	return users, nil
}

// NewUserRepository creates a new user repository
func NewUserRepository(sqlx *sqlx.DB) UserRepo {
	return &userRepositoryImpl{sqlx: sqlx}
}
