package exception

import (
	"github.com/faridlan/auth-go/model/web"
	"github.com/gofiber/fiber/v2"
)

func ExceptionError(ctx *fiber.Ctx, err error) error {

	switch e := err.(type) {
	case *NotFoundError:
		return notFoundError(ctx, e.Error())
	case *BadRequestError:
		return badRequestError(ctx, e.Error())
	case *UnauthorizedError:
		return unauthorizedError(ctx, e.Error())
	default:
		return internalServerError(ctx, err)
	}

}

func internalServerError(ctx *fiber.Ctx, err error) error {

	ctx.Request().Header.Add("content-type", "application/json")
	ctx.Status(fiber.StatusBadRequest)

	webRespone := web.WebResponse{
		Code:   fiber.StatusBadRequest,
		Status: "INTERNAL SERVER ERROR",
		Data:   err.Error(),
	}

	return ctx.JSON(webRespone)

}

func badRequestError(ctx *fiber.Ctx, err string) error {

	ctx.Request().Header.Add("content-type", "application/json")
	ctx.Status(fiber.StatusInternalServerError)

	webRespone := web.WebResponse{
		Code:   fiber.StatusInternalServerError,
		Status: "BAD REQUEST",
		Data:   err,
	}

	return ctx.JSON(webRespone)

}

func unauthorizedError(ctx *fiber.Ctx, err string) error {

	ctx.Request().Header.Add("content-type", "application/json")
	ctx.Status(fiber.StatusUnauthorized)

	webRespone := web.WebResponse{
		Code:   fiber.StatusUnauthorized,
		Status: "UNAUTHORIZED",
		Data:   err,
	}

	return ctx.JSON(webRespone)

}

func notFoundError(ctx *fiber.Ctx, err string) error {

	ctx.Request().Header.Add("content-type", "application/json")
	ctx.Status(fiber.StatusNotFound)

	webRespone := web.WebResponse{
		Code:   fiber.StatusNotFound,
		Status: "NOT FOUND",
		Data:   err,
	}

	return ctx.JSON(webRespone)

}
