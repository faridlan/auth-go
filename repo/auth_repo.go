package repo

import (
	"context"
	"errors"

	"github.com/faridlan/auth-go/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UserRepo interface {
	CreateUserHash(ctx context.Context, user *model.UserHash, db *gorm.DB) (*model.UserHash, error)
	FindUserHash(ctx context.Context, username string, db *gorm.DB) (*model.UserHash, error)
	FindAll(ctx context.Context, db *gorm.DB) ([]*model.UserHash, error)
}

type UserRepoImpl struct {
}

func NewUserRepo() UserRepo {
	return &UserRepoImpl{}
}

func (repo *UserRepoImpl) CreateUserHash(ctx context.Context, user *model.UserHash, db *gorm.DB) (*model.UserHash, error) {

	err := db.Omit("ID").Clauses(clause.Returning{}).Select("username", "hashed_password").Create(&user).Error
	if err != nil {
		return nil, err
	}

	return user, nil

}

func (repo *UserRepoImpl) FindUserHash(ctx context.Context, username string, db *gorm.DB) (*model.UserHash, error) {

	user := &model.UserHash{}
	err := db.Where("username = ?", username).Find(&user).Error
	if err != nil {
		return nil, errors.New("user not found")
	}

	return user, nil

}

func (repo *UserRepoImpl) FindAll(ctx context.Context, db *gorm.DB) ([]*model.UserHash, error) {

	users := []*model.UserHash{}
	err := db.Find(&users).Error
	if err != nil {
		return nil, err
	}

	return users, nil

}

// func (repo *UserRepoImpl) CreateUser(user *model.User, db *gorm.DB) (*model.User, error) {

// 	err := db.Omit("ID").Clauses(clause.Returning{}).Select("username", "password").Create(&user).Error
// 	if err != nil {
// 		return nil, err
// 	}

// 	return user, nil

// }

// func (repo *UserRepoImpl) FindUser(userLogin *model.UserLogin, db *gorm.DB) (*model.UserHash, error) {

// 	user := &model.UserHash{}
// 	err := db.Where("username = ?", userLogin.Username).Where("password = ?", userLogin.Password).Find(&user).Error
// 	if err != nil {
// 		return nil, errors.New("user not found")
// 	}

// 	return user, nil

// }
