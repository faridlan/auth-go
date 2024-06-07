package test

import (
	"fmt"
	"testing"

	"github.com/faridlan/auth-go/helper"
	"github.com/faridlan/auth-go/model/domain"
	"github.com/stretchr/testify/assert"
)

func CreateUser(username string) (*domain.User, error) {

	hashesPassword, err := helper.HashPassword("secret010203")
	if err != nil {
		return nil, err
	}

	err = roleRepo.Truncate(ctx, db)
	if err != nil {
		return nil, err
	}

	role, err := CreateRole("role_test")
	if err != nil {
		return nil, err
	}

	user := &domain.User{
		Username: username,
		Password: hashesPassword,
		RoleId:   role.ID,
	}

	userResponse, err := userRepo.CreateUser(ctx, db, user)
	if err != nil {
		return nil, err
	}

	return userResponse, nil

}

func CreateUserMany(username string) (*domain.User, error) {

	hashesPassword, err := helper.HashPassword("secret010203")
	if err != nil {
		return nil, err
	}

	role, err := CreateRole("role_test_many")
	if err != nil {
		return nil, err
	}

	userResponse := &domain.User{}

	for i := 1; i <= 3; i++ {
		userResponse.Username = fmt.Sprintf(username+"_%d", i)
		userResponse.Password = hashesPassword
		userResponse.RoleId = role.ID

		userResponse, err = userRepo.CreateUser(ctx, db, userResponse)
		if err != nil {
			return nil, err
		}
	}

	return userResponse, nil

}

func TestRegisterRepoSuccess(t *testing.T) {

	err := userRepo.Truncate(ctx, db)
	assert.Nil(t, err)

	err = roleRepo.Truncate(ctx, db)
	assert.Nil(t, err)

	hashesPassword, err := helper.HashPassword("secret010203")
	assert.Nil(t, err)

	role, err := CreateRole("role_test")
	assert.Nil(t, err)

	user := &domain.User{
		Username: "user_repo_7655",
		Password: hashesPassword,
		RoleId:   role.ID,
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

	err = roleRepo.Truncate(ctx, db)
	assert.Nil(t, err)

	_, err = CreateUserMany("user_test")
	assert.Nil(t, err)

	user, err := userRepo.FindAll(ctx, db)
	assert.Nil(t, err)

	assert.Equal(t, 4, len(user))

}
