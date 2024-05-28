package test

import (
	"context"
	"testing"

	"github.com/faridlan/auth-go/config"
	"github.com/faridlan/auth-go/controller"
	"github.com/faridlan/auth-go/helper"
	"github.com/faridlan/auth-go/model/domain"
	"github.com/faridlan/auth-go/repo"
	"github.com/faridlan/auth-go/service"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

var (
	userRepo       = repo.NewUserRepo()
	db             = config.NewDatabase()
	userService    = service.NewUserService(userRepo, db)
	userController = controller.NewUserController(userService)
	app            = fiber.New()
)

// func setupTestFlags(config string) {

// 	flag.CommandLine = flag.NewFlagSet("test", flag.ContinueOnError)
// 	helper.SetupFlags()
// 	flag.Set("config", config)
// 	flag.Parse()

// }

func TestRegisterRepoSuccess(t *testing.T) {

	// setupTestFlags("../.env")

	hashesPassword, err := helper.HashPassword("secret010203")
	assert.Nil(t, err)

	user := &domain.User{
		Username: "user_repo_7655",
		Password: hashesPassword,
	}

	userResponse, err := userRepo.CreateUser(context.Background(), db, user)
	assert.Nil(t, err)

	assert.Equal(t, "user_repo_7655", userResponse.Username)

}

func TestRegisterRepoFailed(t *testing.T) {

	user := &domain.User{
		Username: "user_repo_009",
	}

	_, err := userRepo.CreateUser(context.Background(), db, user)
	assert.NotNil(t, err)

}

func TestFindUserRepoSuccess(t *testing.T) {

	user, err := userRepo.FindUser(context.Background(), db, "user_repo_009")
	assert.Nil(t, err)

	assert.Equal(t, "user_repo_009", user.Username)

}

func TestFindUserRepoFailed(t *testing.T) {

	user, err := userRepo.FindUser(context.Background(), db, "user_repo_009")
	assert.Nil(t, err)

	assert.Equal(t, "user_repo_009", user.Username)

}

func TestFindAllRepoSuccess(t *testing.T) {

	user, err := userRepo.FindAll(context.Background(), db)
	assert.Nil(t, err)

	assert.Equal(t, 13, len(user))

}
