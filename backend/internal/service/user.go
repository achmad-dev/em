package service

/*
--- MIT License (c) 2024 achmad
--- See LICENSE for more details
*/

import (
	"context"
	"errors"

	"github.com/achmad/em/backend/api/dto"
	"github.com/achmad/em/backend/internal/domain"
	"github.com/achmad/em/backend/internal/repository"
	"github.com/achmad/em/backend/pkg/utils"
	"github.com/sirupsen/logrus"
)

// service interface for user
type UserService interface {
	SignIn(ctx context.Context, username, password string) (*dto.AuthResponseDto, error)
	GetUserByID(ctx context.Context, id string) (*domain.User, error)
	GetUsersCompany(ctx context.Context, userType string) ([]*domain.UserCompanyResponse, error)
}

// user service implementation
type userServiceImpl struct {
	userRepo   repository.UserRepo
	bcryptUtil utils.BcryptUtil
	secret     string
	log        *logrus.Logger
}

// GetUserByID implements UserService.
func (u *userServiceImpl) GetUserByID(ctx context.Context, id string) (*domain.User, error) {
	user, err := u.userRepo.GetUserByID(ctx, id)
	if err != nil {
		u.log.Error(err)
		return nil, errors.New("user not found")
	}
	return user, nil
}

// GetUsersCompany implements UserService.
func (u *userServiceImpl) GetUsersCompany(ctx context.Context, userType string) ([]*domain.UserCompanyResponse, error) {
	UserCompanyResponse, err := u.userRepo.GetUsersByType(ctx, userType)
	if err != nil {
		u.log.Error(err)
		return nil, errors.New("something went wrong")
	}

	return UserCompanyResponse, nil
}

// SignIn implements UserService.
func (u *userServiceImpl) SignIn(ctx context.Context, username string, password string) (*dto.AuthResponseDto, error) {
	user, err := u.userRepo.GetUserByName(ctx, username)
	if err != nil {
		u.log.Error(err)
		return nil, errors.New("user not found")
	}
	if !u.bcryptUtil.CheckPasswordHash(password, user.Password) {
		u.log.Error("invalid password")
		return nil, errors.New("invalid password")
	}
	token, err := utils.GenerateToken(user.ID, user.Role, u.secret)
	if err != nil {
		u.log.Error(err)
		return nil, errors.New("something went wrong")
	}

	response := &dto.AuthResponseDto{
		Token: token,
		Role:  user.Role,
	}

	return response, nil
}

// NewUserService creates a new user service
func NewUserService(userRepo repository.UserRepo, bcryptUtil utils.BcryptUtil, secret string, log *logrus.Logger) UserService {
	return &userServiceImpl{userRepo: userRepo, bcryptUtil: bcryptUtil, secret: secret, log: log}
}
