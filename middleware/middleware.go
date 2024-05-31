package middleware

import (
	"context"
	"fmt"
	"time"

	"github.com/faridlan/auth-go/config"
	"github.com/faridlan/auth-go/exception"
	"github.com/faridlan/auth-go/helper"
	jwtconfig "github.com/faridlan/auth-go/helper/jwt_config"
	"github.com/faridlan/auth-go/repo"
	"github.com/gofiber/fiber/v2"
)

func AuthMiddleware(ctx *fiber.Ctx) error {

	authorization := ctx.Get("Authorization")

	if ctx.Path() == "/api/users/login" || (ctx.Path() == "/api/users" && ctx.Method() == "POST") {
		return ctx.Next()
	}

	if len(authorization) < 8 || authorization[:7] != "Bearer " {
		return exception.NewUnauthorizedError("Missing or invalid token format")
	}

	tokenBearer := authorization[7:]

	//
	db := config.NewDatabase()
	authToken, _ := jwtconfig.ParseJwtAuth(ctx)
	whitelistRepo := repo.NewWhitelistRepo()
	_, err := whitelistRepo.FindById(context.Background(), db, authToken)
	if err != nil {
		return exception.NewUnauthorizedError(err.Error())
	}
	//

	config, err := helper.GetEnv()
	if err != nil {
		return err
	}

	path := config.GetString("PRIVATE_KEY")

	privateKey, err := jwtconfig.LoadPrivateKey(path)
	if err != nil {
		return exception.NewUnauthorizedError(err.Error())
	}

	claims, _, err := jwtconfig.VerifyToken(tokenBearer, &privateKey.PublicKey)
	if err != nil {
		return exception.NewUnauthorizedError(err.Error())
	}

	// Check token expiration time
	exp := claims.ExpiresAt
	fmt.Println(exp.Time)

	timeUntilExpiry := time.Until(exp.Time)

	if timeUntilExpiry < 5*time.Minute {
		newToken, err := jwtconfig.GenerateJWT(claims)
		// newToken, err := GenerateToken(claims.User)
		if err != nil {
			return err
			// ctx.Status(fiber.StatusInternalServerError)
			// return ctx.SendString("Failed to generate new token: " + err.Error())
		}

		fmt.Println(newToken)
		// Set the new token in the response header
		ctx.Set("Authorization", "Bearer "+newToken)
	}

	fmt.Println(exp.Time)
	return ctx.Next()

}
