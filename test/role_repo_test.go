package test

import (
	"fmt"
	"testing"

	"github.com/faridlan/auth-go/model/domain"
	"github.com/stretchr/testify/assert"
)

func CreateRole(roleName string) (*domain.Role, error) {

	role := &domain.Role{
		Name: roleName,
	}

	roleRespone, err := roleRepo.Save(ctx, db, role)
	if err != nil {
		return nil, err
	}

	return roleRespone, nil

}

func TestCreateRoleRepoSuccess(t *testing.T) {

	err := userRepo.Truncate(ctx, db)
	assert.Nil(t, err)

	err = roleRepo.Truncate(ctx, db)
	assert.Nil(t, err)

	role := &domain.Role{
		Name: "role_test",
	}

	roleResponse, err := roleRepo.Save(ctx, db, role)
	assert.Nil(t, err)

	assert.Equal(t, "role_test", roleResponse.Name)

}

func TestCreateRoleRepoFailed(t *testing.T) {

	err := userRepo.Truncate(ctx, db)
	assert.Nil(t, err)

	err = roleRepo.Truncate(ctx, db)
	assert.Nil(t, err)

	_, err = CreateRole("role_test")
	assert.Nil(t, err)

	role := &domain.Role{
		Name: "role_test",
	}

	_, err = roleRepo.Save(ctx, db, role)
	assert.NotNil(t, err)

}

func TestFindByIdRoleRepoSuccess(t *testing.T) {

	err := userRepo.Truncate(ctx, db)
	assert.Nil(t, err)

	err = roleRepo.Truncate(ctx, db)
	assert.Nil(t, err)

	role, err := CreateRole("role_test")
	assert.Nil(t, err)

	roleResponse, err := roleRepo.FindById(ctx, db, role.ID)
	assert.Nil(t, err)

	assert.Equal(t, "role_test", roleResponse.Name)

}

func TestFindByIdRoleRepoFailed(t *testing.T) {

	err := userRepo.Truncate(ctx, db)
	assert.Nil(t, err)

	err = roleRepo.Truncate(ctx, db)
	assert.Nil(t, err)

	_, err = CreateRole("role_test")
	assert.Nil(t, err)

	_, err = roleRepo.FindById(ctx, db, "salah")
	assert.NotNil(t, err)

	assert.Equal(t, "role not found", err.Error())

}

func TestFindAllRoleRepoSuccess(t *testing.T) {

	err := userRepo.Truncate(ctx, db)
	assert.Nil(t, err)

	err = roleRepo.Truncate(ctx, db)
	assert.Nil(t, err)

	for i := 1; i <= 3; i++ {
		_, err := CreateRole(fmt.Sprintf("role_test_%d", i))
		assert.Nil(t, err)
	}

	role, err := roleRepo.FindAll(ctx, db)
	assert.Nil(t, err)

	assert.Equal(t, 4, len(role))

}
