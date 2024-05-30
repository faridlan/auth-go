package test

import (
	"testing"

	"github.com/faridlan/auth-go/helper"
	"github.com/faridlan/auth-go/model/domain"
	"github.com/stretchr/testify/assert"
)

func CreateWhitelist() (*domain.Whitelist, error) {

	randomString := helper.RandomString(16)
	whitelist := &domain.Whitelist{
		Token: randomString,
	}
	reponse, err := whitelistRepo.Save(ctx, db, whitelist)
	if err != nil {
		return nil, err
	}

	return reponse, nil

}

func TestCreateWhitelistRepoSuccess(t *testing.T) {

	err := whitelistRepo.Truncate(ctx, db)
	assert.Nil(t, err)

	randomString := helper.RandomString(16)
	whitelist := &domain.Whitelist{
		Token: randomString,
	}

	response, err := whitelistRepo.Save(ctx, db, whitelist)
	assert.Nil(t, err)

	assert.Equal(t, randomString, response.Token)

}

func TestFindWhitelistRepoSuccess(t *testing.T) {

	err := whitelistRepo.Truncate(ctx, db)
	assert.Nil(t, err)

	whitelist, err := CreateWhitelist()
	assert.Nil(t, err)

	response, err := whitelistRepo.FindById(ctx, db, whitelist.Token)
	assert.Nil(t, err)

	assert.Equal(t, whitelist.Token, response.Token)

}

func TestFindWhitelistRepoFailed(t *testing.T) {

	err := whitelistRepo.Truncate(ctx, db)
	assert.Nil(t, err)

	_, err = CreateWhitelist()
	assert.Nil(t, err)

	_, err = whitelistRepo.FindById(ctx, db, "salah")
	assert.NotNil(t, err)

}

func TestDeleteWhitelistRepoSuccess(t *testing.T) {

	err := whitelistRepo.Truncate(ctx, db)
	assert.Nil(t, err)

	whitelist, err := CreateWhitelist()
	assert.Nil(t, err)

	err = whitelistRepo.Delete(ctx, db, whitelist)
	assert.Nil(t, err)

}

func TestDeleteWhitelistRepoFailed(t *testing.T) {

	err := whitelistRepo.Truncate(ctx, db)
	assert.Nil(t, err)

	whitelist := &domain.Whitelist{
		Token: "salah",
	}

	err = whitelistRepo.Delete(ctx, db, whitelist)
	assert.NotNil(t, err)

}
