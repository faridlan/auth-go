package test

import (
	"context"

	"github.com/faridlan/auth-go/config"
	"github.com/faridlan/auth-go/controller"
	"github.com/faridlan/auth-go/exception"

	"github.com/faridlan/auth-go/repo"
	"github.com/faridlan/auth-go/service"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var (
	userRepo       = repo.NewUserRepo()
	whitelistRepo  = repo.NewWhitelistRepo()
	db             = config.NewDatabase()
	validate       = validator.New()
	userService    = service.NewUserService(userRepo, whitelistRepo, db, validate)
	userController = controller.NewUserController(userService)
	app            = fiber.New(
		fiber.Config{
			ErrorHandler: exception.ExceptionError,
		},
	)
	ctx            = context.Background()
	roleRepo       = repo.NewRoleRepo()
	roleService    = service.NewRoleService(roleRepo, db, validate)
	roleController = controller.NewRoleController(roleService)
)
