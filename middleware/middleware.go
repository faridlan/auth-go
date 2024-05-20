package middleware

import (
	"fmt"
	"os"
	"time"

	jwtconfig "github.com/faridlan/auth-go/helper/jwt_config"
	"github.com/gofiber/fiber/v2"
)

func AuthMiddleware(ctx *fiber.Ctx) error {

	authorization := ctx.Get("Authorization")

	if ctx.Path() == "/api/users/login" || (ctx.Path() == "/api/users" && ctx.Method() == "POST") {
		return ctx.Next()
	}

	if len(authorization) < 8 || authorization[:7] != "Bearer " {
		ctx.Status(fiber.StatusUnauthorized)
		return ctx.SendString("UNAUTHORIZED")
	}

	tokenBearer := authorization[7:]
	path := os.Getenv("PRIVATE_KEY")

	privateKey, err := jwtconfig.LoadPrivateKey(path)
	if err != nil {
		ctx.Status(fiber.StatusUnauthorized)
		return ctx.SendString("Failed to load private key: " + err.Error())
	}

	claims, _, err := jwtconfig.VerifyToken(tokenBearer, &privateKey.PublicKey)
	if err != nil {
		ctx.Status(fiber.StatusUnauthorized)
		return ctx.SendString("invalid token")
	}

	// Check token expiration time
	exp := claims.ExpiresAt
	fmt.Println(exp.Time)

	timeUntilExpiry := time.Until(exp.Time)

	if timeUntilExpiry < 5*time.Minute {
		newToken, err := jwtconfig.GenerateJWT(claims)
		// newToken, err := GenerateToken(claims.User)
		if err != nil {
			ctx.Status(fiber.StatusInternalServerError)
			return ctx.SendString("Failed to generate new token: " + err.Error())
		}

		fmt.Println(newToken)
		// Set the new token in the response header
		ctx.Set("Authorization", "Bearer "+newToken)
	}

	fmt.Println(exp.Time)
	return ctx.Next()

}
