package models

type Client struct {
	ID int `json:"id" gorm:"column:id;primary_key"`

	Name       *string `json:"name" gorm:"column:name"`
	Surname    *string `json:"surname" gorm:"column:surname"`
	Patronymic *string `json:"patronymic" gorm:"column:patronymic"`

	Phone *string `json:"phone" gorm:"column:phone"`
	Email *string `json:"email" gorm:"column:email"`
}

func (Client) TableName() string {
	return "clients"
}
