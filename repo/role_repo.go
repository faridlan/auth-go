package repo

import (
	"context"
	"errors"

	"github.com/faridlan/auth-go/model/domain"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type RoleRepo interface {
	Save(ctx context.Context, db *gorm.DB, role *domain.Role) (*domain.Role, error)
	FindById(ctx context.Context, db *gorm.DB, roleID string) (*domain.Role, error)
	FindAll(ctx context.Context, db *gorm.DB) ([]*domain.Role, error)
	Truncate(ctx context.Context, db *gorm.DB) error
}

type RoleRepoImpl struct {
}

func NewRoleRepo() RoleRepo {
	return &RoleRepoImpl{}
}

func (repo *RoleRepoImpl) Save(ctx context.Context, db *gorm.DB, role *domain.Role) (*domain.Role, error) {

	err := db.Omit("ID").Clauses(clause.Returning{}).Select("name").Create(&role).Error
	if err != nil {
		return nil, err
	}

	return role, nil

}

func (repo *RoleRepoImpl) FindById(ctx context.Context, db *gorm.DB, roleID string) (*domain.Role, error) {

	role := &domain.Role{}
	err := db.First(&role, "ID = ?", roleID).Error
	if err != nil {
		return nil, errors.New("role not found")
	}

	return role, nil

}

func (repo *RoleRepoImpl) FindAll(ctx context.Context, db *gorm.DB) ([]*domain.Role, error) {

	roles := []*domain.Role{}
	err := db.Find(&roles).Error
	if err != nil {
		return nil, err
	}

	return roles, nil

}

func (repo *RoleRepoImpl) Truncate(ctx context.Context, db *gorm.DB) error {

	err := db.Exec("DELETE FROM roles WHERE name <> 'role_root'").Error
	if err != nil {
		return err
	}

	return nil

}
