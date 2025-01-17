package models

import "time"

type Offer struct {
	ID int `json:"id" gorm:"column:id;primary_key"`

	RealEstateObjectID int               `json:"real_estate_object_id" gorm:"column:real_estate_object_id"`
	RealEstateObject   *RealEstateObject `json:"real_estate_object" gorm:"foreignKey:RealEstateObjectID"`

	ClientID  int      `json:"client_id" gorm:"column:client_id"`
	Client    *Client  `json:"client" gorm:"foreignKey:ClientID"`
	RealtorID int      `json:"realtor_id" gorm:"column:realtor_id"`
	Realtor   *Realtor `json:"realtor" gorm:"foreignKey:RealtorID"`

	Price     int        `json:"price" gorm:"column:price"`
	CreatedAt *time.Time `json:"created_at" gorm:"column:created_at"`
}

func (Offer) TableName() string {
	return "offers"
}
