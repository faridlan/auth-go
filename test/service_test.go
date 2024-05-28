package test

import (
	"context"
	"testing"

	"github.com/faridlan/auth-go/model/web"
	"github.com/stretchr/testify/assert"
)

func TestRegisterServiceSucces(t *testing.T) {

	user := &web.UserCreate{
		Username: "user_service_909",
		Password: "secret010203",
	}

	userResponse, err := userService.Register(context.Background(), user)
	assert.Nil(t, err)

	assert.Equal(t, "user_service_909", userResponse.Username)

}

func TestRegisterServiceFailed(t *testing.T) {

	user := &web.UserCreate{
		Username: "user_service_009",
	}

	_, err := userService.Register(context.Background(), user)
	assert.NotNil(t, err)

}

func TestLoginServiceSuccess(t *testing.T) {
	user := &web.UserCreate{
		Username: "user_service_009",
		Password: "secret010203",
	}

	userResponse, err := userService.Login(context.Background(), user)
	assert.Nil(t, err)

	assert.Equal(t, "user_service_009", userResponse.User.Username)
	assert.NotNil(t, userResponse.Token)
}

func TestLoginServiceFailed(t *testing.T) {
	user := &web.UserCreate{
		Username: "user_service_009",
	}

	_, err := userService.Login(context.Background(), user)
	assert.NotNil(t, err)
}

func TestFindAllServiceSuccess(t *testing.T) {

	userResponse, err := userService.FindAll(context.Background())
	assert.Nil(t, err)

	assert.Equal(t, 13, len(userResponse))

}
