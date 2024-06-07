package test

import (
	"fmt"
	"io"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateRoleControllerSuccess(t *testing.T) {

	err := userRepo.Truncate(ctx, db)
	assert.Nil(t, err)

	err = roleRepo.Truncate(ctx, db)
	assert.Nil(t, err)

	app.Post("/api/roles", roleController.Create)

	body := strings.NewReader(
		`{
			"name" : "role_test"
		}`,
	)

	request := httptest.NewRequest("POST", "/api/roles", body)
	request.Header.Set("content-type", "application/json")
	response, err := app.Test(request)
	assert.Nil(t, err)

	assert.Equal(t, 200, response.StatusCode)
	byte, err := io.ReadAll(response.Body)
	assert.Nil(t, err)

	fmt.Println(string(byte))

}

func TestCreateRoleControllerFailed(t *testing.T) {

	err := userRepo.Truncate(ctx, db)
	assert.Nil(t, err)

	err = roleRepo.Truncate(ctx, db)
	assert.Nil(t, err)

	app.Post("/api/roles", roleController.Create)

	body := strings.NewReader(
		`{
			"name" : ""
		}`,
	)

	request := httptest.NewRequest("POST", "/api/roles", body)
	request.Header.Set("content-type", "application/json")
	response, err := app.Test(request)
	assert.Nil(t, err)

	assert.Equal(t, 500, response.StatusCode)

}

func TestFindByIdRoleControllerSuccess(t *testing.T) {

	err := userRepo.Truncate(ctx, db)
	assert.Nil(t, err)

	err = roleRepo.Truncate(ctx, db)
	assert.Nil(t, err)

	app.Get("/api/roles/:id", roleController.FindById)

	role, err := CreateRole("role_test")
	assert.Nil(t, err)

	request := httptest.NewRequest("GET", fmt.Sprintf("/api/roles/%s", role.ID), nil)
	request.Header.Set("content-type", "application/json")
	response, err := app.Test(request)
	assert.Nil(t, err)

	assert.Equal(t, 200, response.StatusCode)
	byte, err := io.ReadAll(response.Body)
	assert.Nil(t, err)

	fmt.Println(string(byte))

}

func TestFindByIdRoleControllerFailed(t *testing.T) {

	err := userRepo.Truncate(ctx, db)
	assert.Nil(t, err)

	err = roleRepo.Truncate(ctx, db)
	assert.Nil(t, err)

	app.Get("/api/roles/:id", roleController.FindById)

	_, err = CreateRole("role_test")
	assert.Nil(t, err)

	request := httptest.NewRequest("GET", fmt.Sprintf("/api/roles/%s", "salah"), nil)
	response, err := app.Test(request)
	assert.Nil(t, err)

	assert.Equal(t, 404, response.StatusCode)

}

func TestFindAllRoleControllerSuccess(t *testing.T) {

	err := userRepo.Truncate(ctx, db)
	assert.Nil(t, err)

	err = roleRepo.Truncate(ctx, db)
	assert.Nil(t, err)

	app.Get("/api/roles", roleController.FindAll)

	for i := 1; i <= 3; i++ {
		_, err := CreateRole(fmt.Sprintf("role_test_%d", i))
		assert.Nil(t, err)
	}

	request := httptest.NewRequest("GET", "/api/roles", nil)
	request.Header.Set("content-type", "application/json")
	response, err := app.Test(request)
	assert.Nil(t, err)

	assert.Equal(t, 200, response.StatusCode)
	byte, err := io.ReadAll(response.Body)
	assert.Nil(t, err)

	fmt.Println(string(byte))

}
