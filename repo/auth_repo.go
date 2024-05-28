package repo

import (
	"context"
	"errors"

	"github.com/faridlan/auth-go/model/domain"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UserRepo interface {
	CreateUser(ctx context.Context, db *gorm.DB, user *domain.User) (*domain.User, error)
	FindUser(ctx context.Context, db *gorm.DB, username string) (*domain.User, error)
	FindAll(ctx context.Context, db *gorm.DB) ([]*domain.User, error)
}

type UserRepoImpl struct {
}

func NewUserRepo() UserRepo {
	return &UserRepoImpl{}
}

func (repo *UserRepoImpl) CreateUser(ctx context.Context, db *gorm.DB, user *domain.User) (*domain.User, error) {

	err := db.Omit("ID").Clauses(clause.Returning{}).Select("username", "hashed_password").Create(&user).Error
	if err != nil {
		return nil, err
	}

	return user, nil

}

func (repo *UserRepoImpl) FindUser(ctx context.Context, db *gorm.DB, username string) (*domain.User, error) {

	user := &domain.User{}
	err := db.Where("username = ?", username).Find(&user).Error
	if err != nil {
		return nil, errors.New("user not found")
	}

	return user, nil

}

func (repo *UserRepoImpl) FindAll(ctx context.Context, db *gorm.DB) ([]*domain.User, error) {

	users := []*domain.User{}
	err := db.Find(&users).Error
	if err != nil {
		return nil, err
	}

	return users, nil

}
