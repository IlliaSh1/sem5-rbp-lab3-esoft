package models

import "time"

type Deal struct {
	ID        int        `json:"id" gorm:"column:id;primary_key"`
	OfferID   int        `json:"offer_id" gorm:"column:offer_id"`
	Offer     *Offer     `json:"offer" gorm:"foreignKey:OfferID"`
	NeedID    int        `json:"need_id" gorm:"column:need_id"`
	Need      *Need      `json:"need" gorm:"foreignKey:NeedID"`
	CreatedAt *time.Time `json:"created_at" gorm:"column:created_at"`
}

func (Deal) TableName() string {
	return "deals"
}
