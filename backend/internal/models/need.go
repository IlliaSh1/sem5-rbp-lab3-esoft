package models

import "time"

type Need struct {
	ID int `json:"id" gorm:"column:id;primary_key"`

	RealEstateObjectTypeID int `json:"real_estate_object_type_id" gorm:"column:real_estate_object_type_id"`
	// RealEstateObject       *RealEstateObjectType `json:"real_estate_object" gorm:"foreignKey:RealEstateObjectID"`

	ClientID  int      `json:"client_id" gorm:"column:client_id"`
	Client    *Client  `json:"client" gorm:"foreignKey:ClientID"`
	RealtorID int      `json:"realtor_id" gorm:"column:realtor_id"`
	Realtor   *Realtor `json:"realtor" gorm:"foreignKey:RealtorID"`

	MinPrice  *uint      `json:"min_price" gorm:"column:min_price"`
	MaxPrice  *uint      `json:"max_price" gorm:"column:max_price"`
	CreatedAt *time.Time `json:"created_at" gorm:"column:created_at"`
}

func (*Need) TableName() string {
	return "needs"
}
