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

	userRepo := repo.NewUserRepo()
	userService := service.NewUserService(userRepo, db)
	userControlelr := controller.NewUserController(userService)

	app := fiber.New()
	app.Use(middleware.AuthMiddleware)

	app.Post("api/users", userControlelr.Register)
	app.Post("api/users/login", userControlelr.Login)
	app.Get("api/users", userControlelr.FindAll)

	app.Listen("localhost:2020")

}
