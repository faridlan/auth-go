package service

import (
	"context"
	"errors"

	"github.com/faridlan/auth-go/helper"
	jwtconfig "github.com/faridlan/auth-go/helper/jwt_config"
	"github.com/faridlan/auth-go/model/domain"
	"github.com/faridlan/auth-go/model/web"
	"github.com/faridlan/auth-go/repo"
	"gorm.io/gorm"
)

type UserService interface {
	Register(ctx context.Context, request *web.UserCreate) (*web.UserResponse, error)
	Login(ctx context.Context, request *web.UserCreate) (*web.UserResponseLogin, error)
	FindAll(ctx context.Context) ([]*web.UserResponse, error)
}

type UserServiceImpl struct {
	UserRepo repo.UserRepo
	DB       *gorm.DB
}

func NewUserService(userRepo repo.UserRepo, db *gorm.DB) UserService {
	return &UserServiceImpl{
		UserRepo: userRepo,
		DB:       db,
	}
}

func (service *UserServiceImpl) Register(ctx context.Context, request *web.UserCreate) (*web.UserResponse, error) {

	passwordHash, err := helper.HashPassword(request.Password)
	if err != nil {
		return nil, err
	}

	user := domain.User{
		Username: request.Username,
		Password: passwordHash,
	}

	userResponse, err := service.UserRepo.CreateUser(ctx, service.DB, &user)
	if err != nil {
		return nil, err
	}

	return helper.ToUserResponse(userResponse), err

}

func (service *UserServiceImpl) Login(ctx context.Context, request *web.UserCreate) (*web.UserResponseLogin, error) {

	user, err := service.UserRepo.FindUser(ctx, service.DB, request.Username)
	if err != nil {
		return nil, err
	}

	if !helper.ComparePasswords(user.Password, request.Password) {
		return nil, errors.New("username or password incorrect")
	}

	claims := &web.Claim{
		User: web.UserResponse{
			ID:       user.ID,
			Username: user.Username,
		},
	}

	tokenString, err := jwtconfig.GenerateJWT(claims)
	if err != nil {
		return nil, err
	}

	userLogin := helper.ToUserResponseLogin(user)
	userLogin.Token = tokenString

	return userLogin, nil

}

func (service *UserServiceImpl) FindAll(ctx context.Context) ([]*web.UserResponse, error) {

	users, err := service.UserRepo.FindAll(ctx, service.DB)
	if err != nil {
		return nil, err
	}

	return helper.ToUserResponses(users), nil

}
