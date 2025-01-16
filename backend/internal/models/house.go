package models

type House struct {
	// ID int `json:"id" gorm:"column:id;primary_key"`

	RealEstateObjectID int               `json:"real_estate_object_id" gorm:"column:real_estate_object_id;primary_key"`
	RealEstateObject   *RealEstateObject `json:"real_estate_object" gorm:"foreignKey:RealEstateObjectID"`

	FloorsCount *uint `json:"floors_count" gorm:"column:floors_count"`
	RoomsCount  *uint `json:"rooms_count" gorm:"column:rooms_count"`
	Area        *uint `json:"area" gorm:"column:area"`
}

func (House) TableName() string {
	return "houses"
}
