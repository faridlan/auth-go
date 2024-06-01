package domain

import "gorm.io/gorm"

type Role struct {
	ID        string         `gorm:"primarykey;column:id;<-:create"`
	Name      string         `gorm:"column:name"`
	CreatedAt int64          `gorm:"column:created_at;autoCreateTime:milli"`
	UpdatedAt int64          `gorm:"column:updated_at;autoCreateTime:milli;autoUpdateTime:milli"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at"`
}

func (u *Role) TableName() string {
	return "roles"
}
