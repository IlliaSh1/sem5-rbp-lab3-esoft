package models

type RealEstateObject struct {
	ID int `json:"id"`

	RealEstateObjectTypeID int `json:"real_estate_object_type_id" gorm:"column:real_estate_object_type_id"`

	Coordinates *Point `json:"coordinates" gorm:"column:coordinates"`

	City            *string `json:"city" gorm:"column:city"`
	Street          *string `json:"street" gorm:"column:street"`
	HouseNumber     *string `json:"house_number" gorm:"column:house_number"`
	ApartmentNumber *string `json:"apartment_number" gorm:"column:apartment_number"`
}

func (*RealEstateObject) TableName() string {
	return "real_estate_objects"
}
