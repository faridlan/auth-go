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

	app.Post("/api/users/", userController.Register)

	body := strings.NewReader(
		`{
			"username" : "username_controller_009",
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

	app.Post("/api/users/", userController.Register)

	body := strings.NewReader(
		`{
			"username" : "username_controller_009"
		}`,
	)

	request := httptest.NewRequest("POST", "/api/users/", body)
	request.Header.Set("content-type", "application/json")
	response, err := app.Test(request)
	assert.Nil(t, err)

	assert.Equal(t, 500, response.StatusCode)

}

func TestLoginControllerSuccess(t *testing.T) {

	app.Post("/api/users/login", userController.Login)

	body := strings.NewReader(
		`{
			"username" : "username_controller_009",
			"password" : "secret01020304"
		}`,
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

	app.Post("/api/users/login", userController.Login)

	body := strings.NewReader(
		`{
			"username" : "username_controller_009"
		}`,
	)

	request := httptest.NewRequest("POST", "/api/users/login", body)
	request.Header.Set("content-type", "application/json")
	response, err := app.Test(request)
	assert.Nil(t, err)

	assert.Equal(t, 500, response.StatusCode)

}

func TestFindAllControllerSuccess(t *testing.T) {

	app.Use(middleware.AuthMiddleware)

	//Login First
	app.Post("/api/users/login", userController.Login)

	body := strings.NewReader(
		`{
				"username" : "username_controller_009",
				"password" : "secret01020304"
		}
		`,
	)

	request := httptest.NewRequest("POST", "/api/users/login", body)
	request.Header.Set("content-type", "application/json")
	response, err := app.Test(request)
	assert.Nil(t, err)

	assert.Equal(t, 200, response.StatusCode)

	userResponse, err := GetToken(response.Body)
	assert.Nil(t, err)

	//Get All User
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
