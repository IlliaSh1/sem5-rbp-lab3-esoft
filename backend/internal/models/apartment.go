package models

type Apartment struct {
	// ID int `json:"id" gorm:"column:id;primary_key"`

	RealEstateObjectID int               `json:"real_estate_object_id" gorm:"column:real_estate_object_id;primary_key"`
	RealEstateObject   *RealEstateObject `json:"real_estate_object" gorm:"foreignKey:RealEstateObjectID"`

	FloorNumber *int  `json:"floor_number" gorm:"column:floor_number"`
	RoomsCount  *uint `json:"rooms_count" gorm:"column:rooms_count"`
	Area        *uint `json:"area" gorm:"column:area"`
}

func (Apartment) TableName() string {
	return "apartments"
}
