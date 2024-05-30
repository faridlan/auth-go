package domain

type Whitelist struct {
	ID    string `gorm:"primarykey;column:id;<-:create"`
	Token string `gorm:"column:token"`
}

func (u *Whitelist) TableName() string {
	return "whitelist"
}
