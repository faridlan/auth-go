package repo

import (
	"context"
	"errors"

	"github.com/faridlan/auth-go/model/domain"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type WhitelistRepo interface {
	Save(ctx context.Context, db *gorm.DB, whitelist *domain.Whitelist) (*domain.Whitelist, error)
	FindById(ctx context.Context, db *gorm.DB, token string) (*domain.Whitelist, error)
	Delete(ctx context.Context, db *gorm.DB, whitelist *domain.Whitelist) error
	Truncate(ctx context.Context, db *gorm.DB) error
}

type WhitelistRepoImpl struct {
}

func NewWhitelistRepo() WhitelistRepo {
	return &WhitelistRepoImpl{}
}

func (repo *WhitelistRepoImpl) Save(ctx context.Context, db *gorm.DB, whitelist *domain.Whitelist) (*domain.Whitelist, error) {

	err := db.Omit("ID").Clauses(clause.Returning{}).Select("token").Create(&whitelist).Error
	if err != nil {
		return nil, err
	}

	return whitelist, nil

}

func (repo *WhitelistRepoImpl) FindById(ctx context.Context, db *gorm.DB, token string) (*domain.Whitelist, error) {

	whitelist := &domain.Whitelist{}
	err := db.Where("token = ?", token).Take(whitelist).Error
	if err != nil {
		return nil, errors.New("whitelist not found")
	}

	return whitelist, nil

}

func (repo *WhitelistRepoImpl) Delete(ctx context.Context, db *gorm.DB, whitelist *domain.Whitelist) error {

	err := db.Delete(&whitelist).Error
	if err != nil {
		return errors.New("delete canceled")
	}

	return nil

}

func (repo *WhitelistRepoImpl) Truncate(ctx context.Context, db *gorm.DB) error {

	err := db.Exec("TRUNCATE whitelist CASCADE").Error
	if err != nil {
		return err
	}

	return nil

}
