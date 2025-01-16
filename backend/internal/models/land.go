package models

type Land struct {
	// ID int `json:"id" gorm:"column:id;primary_key"`

	RealEstateObjectID int               `json:"real_estate_object_id" gorm:"column:real_estate_object_id;primary_key"`
	RealEstateObject   *RealEstateObject `json:"real_estate_object" gorm:"foreignKey:RealEstateObjectID"`

	Area int `json:"area" gorm:"column:area"`
}

func (Land) TableName() string {
	return "lands"
}
