package test

import (
	"fmt"
	"testing"

	"github.com/faridlan/auth-go/model/web"
	"github.com/stretchr/testify/assert"
)

func TestCreateRoleServiceSuccess(t *testing.T) {

	err := userRepo.Truncate(ctx, db)
	assert.Nil(t, err)

	err = roleRepo.Truncate(ctx, db)
	assert.Nil(t, err)

	role := &web.RoleCreate{
		Name: "role_test_service",
	}

	roleResponse, err := roleService.Create(ctx, role)
	assert.Nil(t, err)

	assert.Equal(t, "role_test_service", roleResponse.Name)

}

func TestCreateRoleServiceFailed(t *testing.T) {

	err := userRepo.Truncate(ctx, db)
	assert.Nil(t, err)

	err = roleRepo.Truncate(ctx, db)
	assert.Nil(t, err)

	role := &web.RoleCreate{
		Name: "",
	}

	_, err = roleService.Create(ctx, role)
	assert.NotNil(t, err)
	assert.Equal(t, "Name is a required field", err.Error())

}

func TestFindByIdRoleServicecSuccess(t *testing.T) {

	err := userRepo.Truncate(ctx, db)
	assert.Nil(t, err)

	err = roleRepo.Truncate(ctx, db)
	assert.Nil(t, err)

	role, err := CreateRole("role_test_service")
	assert.Nil(t, err)

	roleResponse, err := roleService.FindById(ctx, role.ID)

	assert.Nil(t, err)

	assert.Equal(t, "role_test_service", roleResponse.Name)

}

func TestFindByIdRoleServicecFailed(t *testing.T) {

	err := userRepo.Truncate(ctx, db)
	assert.Nil(t, err)

	err = roleRepo.Truncate(ctx, db)
	assert.Nil(t, err)

	_, err = CreateRole("role_test_service")
	assert.Nil(t, err)

	_, err = roleService.FindById(ctx, "salah")
	assert.NotNil(t, err)
	assert.Equal(t, "role not found", err.Error())

	// SELECT * FROM "roles" WHERE ID = 'salah' AND "roles"."deleted_at" IS NULL ORDER BY "roles"."id" LIMIT 1
}

func TestFindAllRoleServiceSuccess(t *testing.T) {

	err := userRepo.Truncate(ctx, db)
	assert.Nil(t, err)

	err = roleRepo.Truncate(ctx, db)
	assert.Nil(t, err)

	for i := 1; i <= 3; i++ {
		_, err := CreateRole(fmt.Sprintf("role_test_service_%d", i))
		assert.Nil(t, err)
	}

	roleResponse, err := roleService.FindAll(ctx)
	assert.Nil(t, err)

	assert.Equal(t, 4, len(roleResponse))

}
