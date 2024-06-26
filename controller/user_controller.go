package controller

import (
	jwtconfig "github.com/faridlan/auth-go/helper/jwt_config"
	"github.com/faridlan/auth-go/model/web"
	"github.com/faridlan/auth-go/service"
	"github.com/gofiber/fiber/v2"
)

type UserController interface {
	Register(ctx *fiber.Ctx) error
	Login(ctx *fiber.Ctx) error
	FindAll(ctx *fiber.Ctx) error
	Logout(ctx *fiber.Ctx) error
}

type UserControllerImpl struct {
	UserService service.UserService
}

func NewUserController(userService service.UserService) UserController {
	return &UserControllerImpl{
		UserService: userService,
	}
}

func (controller *UserControllerImpl) Register(ctx *fiber.Ctx) error {

	user := new(web.UserCreate)
	err := ctx.BodyParser(user)
	if err != nil {
		return err
	}

	userResponse, err := controller.UserService.Register(ctx.Context(), user)
	if err != nil {
		return err
	}

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   userResponse,
	}

	return ctx.JSON(webResponse)

}

func (controller *UserControllerImpl) Login(ctx *fiber.Ctx) error {

	user := new(web.UserLogin)
	err := ctx.BodyParser(user)
	if err != nil {
		return err
	}

	userResponse, err := controller.UserService.Login(ctx.Context(), user)
	if err != nil {
		return err
	}

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   userResponse,
	}

	return ctx.JSON(webResponse)

}
func (controller *UserControllerImpl) FindAll(ctx *fiber.Ctx) error {

	userResponses, err := controller.UserService.FindAll(ctx.Context())
	if err != nil {
		return err
	}

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   userResponses,
	}

	return ctx.JSON(webResponse)

}

func (controller *UserControllerImpl) Logout(ctx *fiber.Ctx) error {

	whitelist, _ := jwtconfig.ParseJwtAuth(ctx)
	err := controller.UserService.Logout(ctx.Context(), whitelist)
	if err != nil {
		return err
	}

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}

	return ctx.JSON(webResponse)

}
