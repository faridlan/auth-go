package main

import (
	"github.com/faridlan/auth-go/config"
	"github.com/faridlan/auth-go/controller"
	"github.com/faridlan/auth-go/middleware"
	"github.com/faridlan/auth-go/repo"
	"github.com/faridlan/auth-go/service"
	"github.com/gofiber/fiber/v2"
)

func main() {

	db := config.NewDatabase()

	whitelistRepo := repo.NewWhitelistRepo()

	userRepo := repo.NewUserRepo()
	userService := service.NewUserService(userRepo, whitelistRepo, db)
	userControlelr := controller.NewUserController(userService)

	app := fiber.New()
	app.Use(middleware.AuthMiddleware)

	app.Post("api/users", userControlelr.Register)
	app.Post("api/users/login", userControlelr.Login)
	app.Post("api/users/logout", userControlelr.Logout)
	app.Get("api/users", userControlelr.FindAll)

	app.Listen("localhost:2020")

}
