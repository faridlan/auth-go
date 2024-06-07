package controller

import (
	"github.com/faridlan/auth-go/model/web"
	"github.com/faridlan/auth-go/service"
	"github.com/gofiber/fiber/v2"
)

type RoleController interface {
	Create(ctx *fiber.Ctx) error
	FindById(ctx *fiber.Ctx) error
	FindAll(ctx *fiber.Ctx) error
}

type RoleControllerImpl struct {
	RoleService service.RoleService
}

func NewRoleController(roleService service.RoleService) RoleController {
	return &RoleControllerImpl{
		RoleService: roleService,
	}
}

func (controller *RoleControllerImpl) Create(ctx *fiber.Ctx) error {

	roleCreate := new(web.RoleCreate)
	err := ctx.BodyParser(roleCreate)
	if err != nil {
		return err
	}

	roleResponse, err := controller.RoleService.Create(ctx.Context(), roleCreate)
	if err != nil {
		return err
	}

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   roleResponse,
	}

	return ctx.JSON(webResponse)

}

func (controller *RoleControllerImpl) FindById(ctx *fiber.Ctx) error {

	id := ctx.Params("id")

	roleResponse, err := controller.RoleService.FindById(ctx.Context(), id)
	if err != nil {
		return err
	}

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   roleResponse,
	}

	return ctx.JSON(webResponse)
}

func (controller *RoleControllerImpl) FindAll(ctx *fiber.Ctx) error {

	roleResponses, err := controller.RoleService.FindAll(ctx.Context())
	if err != nil {
		return err
	}

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   roleResponses,
	}

	return ctx.JSON(webResponse)
}
