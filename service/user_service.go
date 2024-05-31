package service

import (
	"context"
	"errors"

	"github.com/faridlan/auth-go/exception"
	"github.com/faridlan/auth-go/helper"
	jwtconfig "github.com/faridlan/auth-go/helper/jwt_config"
	"github.com/faridlan/auth-go/model/domain"
	"github.com/faridlan/auth-go/model/web"
	"github.com/faridlan/auth-go/repo"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type UserService interface {
	Register(ctx context.Context, request *web.UserCreate) (*web.UserResponse, error)
	Login(ctx context.Context, request *web.UserCreate) (*web.UserResponseLogin, error)
	Logout(ctx context.Context, whitelistID string) error
	FindAll(ctx context.Context) ([]*web.UserResponse, error)
}

type UserServiceImpl struct {
	UserRepo  repo.UserRepo
	Whitelist repo.WhitelistRepo
	DB        *gorm.DB
	Validate  *validator.Validate
}

func NewUserService(userRepo repo.UserRepo, whitelist repo.WhitelistRepo, db *gorm.DB, validate *validator.Validate) UserService {
	return &UserServiceImpl{
		UserRepo:  userRepo,
		Whitelist: whitelist,
		DB:        db,
		Validate:  validate,
	}
}

func (service *UserServiceImpl) Register(ctx context.Context, request *web.UserCreate) (*web.UserResponse, error) {

	err := service.Validate.Struct(request)
	errString := helper.TranslateError(err, service.Validate)
	if err != nil {
		return nil, exception.NewBadRequestError(errString)
	}

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

	err := service.Validate.Struct(request)
	errString := helper.TranslateError(err, service.Validate)
	if err != nil {
		return nil, exception.NewBadRequestError(errString)
	}

	tx := service.DB.Begin()
	defer tx.Rollback()

	user, err := service.UserRepo.FindUser(ctx, tx, request.Username)
	if err != nil {
		return nil, err
	}

	if !helper.ComparePasswords(user.Password, request.Password) {
		return nil, errors.New("username or password incorrect")
	}

	randomString := helper.RandomString(16)
	tokenWhitelist := &domain.Whitelist{
		Token: randomString,
	}

	token, err := service.Whitelist.Save(ctx, tx, tokenWhitelist)
	if err == nil {
		tx.Commit()
	} else {
		return nil, err
	}

	claims := &web.Claim{
		User: web.UserResponse{
			ID:        user.ID,
			Username:  user.Username,
			Whitelist: token.Token,
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

func (service *UserServiceImpl) Logout(ctx context.Context, whitelistID string) error {

	tx := service.DB.Begin()
	defer tx.Rollback()

	user, err := service.Whitelist.FindById(ctx, tx, whitelistID)
	if err != nil {
		return err
	}

	err = service.Whitelist.Delete(ctx, tx, user)
	if err == nil {
		tx.Commit()
	} else {
		return err
	}

	return nil

}
