package service

import (
	"context"

	"github.com/faridlan/auth-go/exception"
	"github.com/faridlan/auth-go/helper"
	"github.com/faridlan/auth-go/model/domain"
	"github.com/faridlan/auth-go/model/web"
	"github.com/faridlan/auth-go/repo"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type RoleService interface {
	Create(ctx context.Context, request *web.RoleCreate) (*web.RoleResponse, error)
	FindById(ctx context.Context, roleId string) (*web.RoleResponse, error)
	FindAll(ctx context.Context) ([]*web.RoleResponse, error)
}

type RoleServiceImpl struct {
	RoleRepo repo.RoleRepo
	DB       *gorm.DB
	Validate *validator.Validate
}

func NewRoleService(roleRepo repo.RoleRepo, db *gorm.DB, validate *validator.Validate) RoleService {
	return &RoleServiceImpl{
		RoleRepo: roleRepo,
		DB:       db,
		Validate: validate,
	}
}

func (service *RoleServiceImpl) Create(ctx context.Context, request *web.RoleCreate) (*web.RoleResponse, error) {

	err := service.Validate.Struct(request)
	errString := helper.TranslateError(err, service.Validate)
	if err != nil {
		return nil, exception.NewBadRequestError(errString)
	}

	role := domain.Role{
		Name: request.Name,
	}

	roleReponse, err := service.RoleRepo.Save(ctx, service.DB, &role)
	if err != nil {
		return nil, err
	}

	return helper.ToRoleResponse(roleReponse), nil

}

func (service *RoleServiceImpl) FindById(ctx context.Context, roleId string) (*web.RoleResponse, error) {

	roleResponse, err := service.RoleRepo.FindById(ctx, service.DB, roleId)
	if err != nil {
		return nil, &exception.NotFoundError{
			Message: err.Error(),
		}
	}

	return helper.ToRoleResponse(roleResponse), nil

}

func (service *RoleServiceImpl) FindAll(ctx context.Context) ([]*web.RoleResponse, error) {

	roleResponse, err := service.RoleRepo.FindAll(ctx, service.DB)
	if err != nil {
		return nil, err
	}

	return helper.ToRoleResponses(roleResponse), nil

}
