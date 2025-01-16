package models

type Realtor struct {
	ID                     int    `json:"id" gorm:"column:id;primary_key"`
	Name                   string `json:"name" gorm:"column:name"`
	Surname                string `json:"surname" gorm:"column:surname"`
	Patronymic             string `json:"patronymic" gorm:"column:patronymic"`
	PercentageOfCommission *uint  `json:"percentage_of_commission" gorm:"column:percentage_of_commission"`
}

func (Realtor) TableName() string {
	return "realtors"
}
