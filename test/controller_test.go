package test

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/faridlan/auth-go/middleware"
	"github.com/faridlan/auth-go/model/web"
	"github.com/stretchr/testify/assert"
)

func GetToken(r io.Reader) (*web.UserResponseLogin, error) {

	byte, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}

	var webResponse web.WebResponse
	err = json.Unmarshal(byte, &webResponse)
	if err != nil {
		return nil, err
	}

	dataBytes, err := json.Marshal(webResponse.Data)
	if err != nil {
		return nil, err
	}

	var userResponse web.UserResponseLogin
	err = json.Unmarshal(dataBytes, &userResponse)
	if err != nil {
		return nil, err
	}

	return &userResponse, nil

}

func TestRegisterControllerSuccess(t *testing.T) {

	err := userRepo.Truncate(ctx, db)
	assert.Nil(t, err)

	app.Post("/api/users/", userController.Register)

	body := strings.NewReader(
		`{
			"username" : "username_controller_test_001",
			"password" : "secret01020304"
		}`,
	)

	request := httptest.NewRequest("POST", "/api/users/", body)
	request.Header.Set("content-type", "application/json")
	response, err := app.Test(request)
	assert.Nil(t, err)

	assert.Equal(t, 200, response.StatusCode)
	byte, err := io.ReadAll(response.Body)
	assert.Nil(t, err)

	fmt.Println(string(byte))

}

func TestRegisterControllerFailed(t *testing.T) {

	err := userRepo.Truncate(ctx, db)
	assert.Nil(t, err)

	user, err := CreateUser("username_controller_test_001")
	assert.Nil(t, err)

	app.Post("/api/users/", userController.Register)

	body := strings.NewReader(
		fmt.Sprintf(
			`{
			"username" : "%s",
			"password" : "secret010203"
		}`, user.Username),
	)

	request := httptest.NewRequest("POST", "/api/users/", body)
	request.Header.Set("content-type", "application/json")
	response, err := app.Test(request)
	assert.Nil(t, err)

	assert.Equal(t, 500, response.StatusCode)

}

func TestLoginControllerSuccess(t *testing.T) {

	err := userRepo.Truncate(ctx, db)
	assert.Nil(t, err)

	err = whitelistRepo.Truncate(ctx, db)
	assert.Nil(t, err)

	user, err := CreateUser("username_controller_test_001")
	assert.Nil(t, err)

	app.Post("/api/users/login", userController.Login)

	body := strings.NewReader(
		fmt.Sprintf(
			`{
			"username" : "%s",
			"password" : "secret010203"
		}`, user.Username),
	)
	request := httptest.NewRequest("POST", "/api/users/login", body)
	request.Header.Set("content-type", "application/json")
	response, err := app.Test(request)
	assert.Nil(t, err)

	assert.Equal(t, 200, response.StatusCode)
	byte, err := io.ReadAll(response.Body)
	assert.Nil(t, err)

	fmt.Println(string(byte))

}

func TestLoginControllerFailed(t *testing.T) {

	err := userRepo.Truncate(ctx, db)
	assert.Nil(t, err)

	err = whitelistRepo.Truncate(ctx, db)
	assert.Nil(t, err)

	_, err = CreateUser("username_controller_test_001")
	assert.Nil(t, err)

	app.Post("/api/users/login", userController.Login)

	body := strings.NewReader(
		`{
			"username" : "username_controller_009",
			"password" : "secret010203"
		}`,
	)

	request := httptest.NewRequest("POST", "/api/users/login", body)
	request.Header.Set("content-type", "application/json")
	response, err := app.Test(request)
	assert.Nil(t, err)

	assert.Equal(t, 500, response.StatusCode)

}

func TestFindAllControllerSuccess(t *testing.T) {

	err := userRepo.Truncate(ctx, db)
	assert.Nil(t, err)

	err = whitelistRepo.Truncate(ctx, db)
	assert.Nil(t, err)

	user, err := CreateUser("username_controller_test_001")
	assert.Nil(t, err)

	app.Use(middleware.AuthMiddleware)

	//Login First
	app.Post("/api/users/login", userController.Login)

	body := strings.NewReader(
		fmt.Sprintf(
			`{
			"username" : "%s",
			"password" : "secret010203"
		}`, user.Username),
	)

	request := httptest.NewRequest("POST", "/api/users/login", body)
	request.Header.Set("content-type", "application/json")
	response, err := app.Test(request)
	assert.Nil(t, err)

	assert.Equal(t, 200, response.StatusCode)

	userResponse, err := GetToken(response.Body)
	assert.Nil(t, err)

	//Get All User

	for i := 1; i <= 3; i++ {
		_, err := CreateUser(fmt.Sprintf("user_controller_test_%d", i))
		assert.Nil(t, err)
	}

	app.Get("/api/users/", userController.FindAll)
	request1 := httptest.NewRequest("GET", "/api/users/", nil)
	request1.Header.Set("content-type", "application/json")
	request1.Header.Set("Authorization", "Bearer "+userResponse.Token)
	response1, err := app.Test(request1)
	assert.Nil(t, err)

	assert.Equal(t, 200, response1.StatusCode)
	byte, err := io.ReadAll(response1.Body)
	assert.Nil(t, err)

	fmt.Println(string(byte))

}

func TestFindAllControllerFailed(t *testing.T) {

	app.Use(middleware.AuthMiddleware)
	app.Get("/api/users/", userController.FindAll)

	request := httptest.NewRequest("GET", "/api/users/", nil)
	request.Header.Set("content-type", "application/json")
	response, err := app.Test(request)
	assert.Nil(t, err)

	assert.Equal(t, 401, response.StatusCode)
	byte, err := io.ReadAll(response.Body)
	assert.Nil(t, err)

	fmt.Println(string(byte))

}

func TestLogoutControllerSuccess(t *testing.T) {

	err := userRepo.Truncate(ctx, db)
	assert.Nil(t, err)

	err = whitelistRepo.Truncate(ctx, db)
	assert.Nil(t, err)

	user, err := CreateUser("username_controller_test_001")
	assert.Nil(t, err)

	app.Use(middleware.AuthMiddleware)

	//Login First
	app.Post("/api/users/login", userController.Login)

	body := strings.NewReader(
		fmt.Sprintf(
			`{
			"username" : "%s",
			"password" : "secret010203"
		}`, user.Username),
	)

	request := httptest.NewRequest("POST", "/api/users/login", body)
	request.Header.Set("content-type", "application/json")
	response, err := app.Test(request)
	assert.Nil(t, err)

	assert.Equal(t, 200, response.StatusCode)

	userResponse, err := GetToken(response.Body)
	assert.Nil(t, err)

	//Logout User

	app.Post("/api/users/logout", userController.Logout)
	request1 := httptest.NewRequest("POST", "/api/users/logout", nil)
	request1.Header.Set("content-type", "application/json")
	request1.Header.Set("Authorization", "Bearer "+userResponse.Token)
	response1, err := app.Test(request1)
	assert.Nil(t, err)

	assert.Equal(t, 200, response1.StatusCode)

}
