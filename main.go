package main

import (
	"github.com/faridlan/auth-go/config"
	"github.com/faridlan/auth-go/controller"
	"github.com/faridlan/auth-go/exception"
	"github.com/faridlan/auth-go/middleware"
	"github.com/faridlan/auth-go/repo"
	"github.com/faridlan/auth-go/service"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func main() {

	db := config.NewDatabase()
	validator := validator.New()

	whitelistRepo := repo.NewWhitelistRepo()

	userRepo := repo.NewUserRepo()
	userService := service.NewUserService(userRepo, whitelistRepo, db, validator)
	userController := controller.NewUserController(userService)

	roleRepo := repo.NewRoleRepo()
	roleService := service.NewRoleService(roleRepo, db, validator)
	roleController := controller.NewRoleController(roleService)

	app := fiber.New(
		fiber.Config{
			ErrorHandler: exception.ExceptionError,
		},
	)
	app.Use(middleware.AuthMiddleware)

	app.Post("api/users", userController.Register)
	app.Post("api/users/login", userController.Login)
	app.Post("api/users/logout", userController.Logout)
	app.Get("api/users", userController.FindAll)

	app.Post("api/roles", roleController.Create)
	app.Get("api/roles/:id", roleController.FindById)
	app.Get("api/roles", roleController.FindAll)

	app.Listen("localhost:2020")

}
