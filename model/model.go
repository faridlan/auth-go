package model

import "gorm.io/gorm"

type User struct {
	ID        string         `gorm:"primarykey;column:id;<-:create"`
	Username  string         `gorm:"column:username"`
	Password  string         `gorm:"column:password"`
	CreatedAt int64          `gorm:"column:created_at;autoCreateTime:milli"`
	UpdatedAt int64          `gorm:"column:updated_at;autoCreateTime:milli;autoUpdateTime:milli"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at"`
}

func (u *User) TableName() string {
	return "users"
}

type UserHash struct {
	ID        string         `gorm:"primarykey;column:id;<-:create"`
	Username  string         `gorm:"column:username"`
	Password  []byte         `gorm:"column:hashed_password"`
	CreatedAt int64          `gorm:"column:created_at;autoCreateTime:milli"`
	UpdatedAt int64          `gorm:"column:updated_at;autoCreateTime:milli;autoUpdateTime:milli"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at"`
}

func (u *UserHash) TableName() string {
	return "users_hash"
}

type UserLogin struct {
	Username string
	Password string
}

type UserLoginResponse struct {
	User  *UserResponse
	Token string
}
