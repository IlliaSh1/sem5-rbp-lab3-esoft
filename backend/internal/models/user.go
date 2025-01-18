package models

type User struct {
	ID           int    `json:"id" gorm:"column:id;primary_key"`
	Email        string `json:"email" gorm:"column:email"`
	PasswordHash string `json:"-" gorm:"column:password_hash"`
}

func (User) TableName() string {
	return "users"
}
