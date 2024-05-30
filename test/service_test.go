package test

import (
	"fmt"
	"testing"

	"github.com/faridlan/auth-go/model/web"
	"github.com/stretchr/testify/assert"
)

func TestRegisterServiceSucces(t *testing.T) {

	err := userRepo.Truncate(ctx, db)
	assert.Nil(t, err)

	user := &web.UserCreate{
		Username: "user_service_test",
		Password: "secret010203",
	}

	userResponse, err := userService.Register(ctx, user)
	assert.Nil(t, err)

	assert.Equal(t, "user_service_test", userResponse.Username)

}

func TestRegisterServiceFailed(t *testing.T) {

	err := userRepo.Truncate(ctx, db)
	assert.Nil(t, err)

	user, err := CreateUser("user_service_test")
	assert.Nil(t, err)

	userResponse := &web.UserCreate{
		Username: user.Username,
	}

	_, err = userService.Register(ctx, userResponse)
	assert.NotNil(t, err)

}

func TestLoginServiceSuccess(t *testing.T) {

	err := userRepo.Truncate(ctx, db)
	assert.Nil(t, err)

	err = whitelistRepo.Truncate(ctx, db)
	assert.Nil(t, err)

	user, err := CreateUser("user_service_test")
	assert.Nil(t, err)

	userResponse := &web.UserCreate{
		Username: user.Username,
		Password: "secret010203",
	}

	result, err := userService.Login(ctx, userResponse)
	assert.Nil(t, err)

	assert.Equal(t, "user_service_test", result.User.Username)
	assert.NotNil(t, result.Token)

}

func TestLoginServiceFailed(t *testing.T) {

	err := userRepo.Truncate(ctx, db)
	assert.Nil(t, err)

	user := &web.UserCreate{
		Username: "user_service_009",
	}

	_, err = userService.Login(ctx, user)
	assert.NotNil(t, err)

}

func TestFindAllServiceSuccess(t *testing.T) {

	err := userRepo.Truncate(ctx, db)
	assert.Nil(t, err)

	for i := 1; i <= 2; i++ {
		CreateUser(fmt.Sprintf("user_service_test_%d", i))

	}

	userResponse, err := userService.FindAll(ctx)
	assert.Nil(t, err)

	assert.Equal(t, 2, len(userResponse))

}

func TestLogoutServiceSuccess(t *testing.T) {

	err := userRepo.Truncate(ctx, db)
	assert.Nil(t, err)

	err = whitelistRepo.Truncate(ctx, db)
	assert.Nil(t, err)

	whitelist, err := CreateWhitelist()
	assert.Nil(t, err)

	err = userService.Logout(ctx, whitelist.Token)
	assert.Nil(t, err)

}
