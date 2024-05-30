package test

import (
	"context"
	"fmt"
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
	whitelistRepo  = repo.NewWhitelistRepo()
	db             = config.NewDatabase()
	userService    = service.NewUserService(userRepo, whitelistRepo, db)
	userController = controller.NewUserController(userService)
	app            = fiber.New()
	ctx            = context.Background()
)

func CreateUser(username string) (*domain.User, error) {

	hashesPassword, err := helper.HashPassword("secret010203")
	if err != nil {
		return nil, err
	}

	user := &domain.User{
		Username: username,
		Password: hashesPassword,
	}

	userResponse, err := userRepo.CreateUser(ctx, db, user)
	if err != nil {
		return nil, err
	}

	return userResponse, nil

}

func TestRegisterRepoSuccess(t *testing.T) {

	err := userRepo.Truncate(ctx, db)
	assert.Nil(t, err)

	hashesPassword, err := helper.HashPassword("secret010203")
	assert.Nil(t, err)

	user := &domain.User{
		Username: "user_repo_7655",
		Password: hashesPassword,
	}

	userResponse, err := userRepo.CreateUser(ctx, db, user)
	assert.Nil(t, err)

	assert.Equal(t, "user_repo_7655", userResponse.Username)

}

func TestRegisterRepoFailed(t *testing.T) {

	err := userRepo.Truncate(ctx, db)
	assert.Nil(t, err)

	user, err := CreateUser("user_test_001")
	assert.Nil(t, err)

	userResponse := &domain.User{
		Username: user.Username,
	}

	_, err = userRepo.CreateUser(ctx, db, userResponse)
	assert.NotNil(t, err)

}

func TestFindUserRepoSuccess(t *testing.T) {

	err := userRepo.Truncate(ctx, db)
	assert.Nil(t, err)

	user, err := CreateUser("user_test_001")
	assert.Nil(t, err)

	userResponse, err := userRepo.FindUser(ctx, db, user.Username)
	assert.Nil(t, err)

	assert.Equal(t, "user_test_001", userResponse.Username)

}

func TestFindUserRepoFailed(t *testing.T) {

	_, err := userRepo.FindUser(ctx, db, "not_found_user")
	assert.NotNil(t, err)

	assert.Equal(t, "user not found", err.Error())

}

func TestFindAllRepoSuccess(t *testing.T) {

	err := userRepo.Truncate(ctx, db)
	assert.Nil(t, err)

	for i := 1; i <= 2; i++ {
		_, err := CreateUser(fmt.Sprintf("user_test_%d", i))
		assert.Nil(t, err)
	}

	user, err := userRepo.FindAll(ctx, db)
	assert.Nil(t, err)

	assert.Equal(t, 2, len(user))

}
