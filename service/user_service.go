package service

import (
	"context"
	"errors"

	"github.com/faridlan/auth-go/helper"
	jwtconfig "github.com/faridlan/auth-go/helper/jwt_config"
	"github.com/faridlan/auth-go/model"
	"github.com/faridlan/auth-go/repo"
	"gorm.io/gorm"
)

type UserService interface {
	Register(ctx context.Context, request *model.UserCreate) (*model.UserResponse, error)
	Login(ctx context.Context, request *model.UserLogin) (*model.UserResponseLogin, error)
	FindAll(ctx context.Context) ([]*model.UserResponse, error)
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

func (service *UserServiceImpl) Register(ctx context.Context, request *model.UserCreate) (*model.UserResponse, error) {

	passwordHash, err := helper.HashPassword(request.Password)
	if err != nil {
		return nil, err
	}

	user := model.UserHash{
		Username: request.Username,
		Password: passwordHash,
	}

	userResponse, err := service.UserRepo.CreateUserHash(ctx, &user, service.DB)
	if err != nil {
		return nil, err
	}

	return helper.ToUserResponse(userResponse), err

}

func (service *UserServiceImpl) Login(ctx context.Context, request *model.UserLogin) (*model.UserResponseLogin, error) {

	user, err := service.UserRepo.FindUserHash(ctx, request.Username, service.DB)
	if err != nil {
		return nil, err
	}

	if !helper.ComparePasswords(user.Password, request.Password) {
		return nil, errors.New("username or password incorrect")
	}

	claims := &model.Claim{
		User: model.UserResponse{
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

func (service *UserServiceImpl) FindAll(ctx context.Context) ([]*model.UserResponse, error) {

	users, err := service.UserRepo.FindAll(ctx, service.DB)
	if err != nil {
		return nil, err
	}

	return helper.ToUserResponses(users), nil

}
