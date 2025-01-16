package models

type RealEstateObjectType struct {
	ID   int    `json:"id" gorm:"column:id;primary_key"`
	Name string `json:"name" gorm:"column:name"`
}

func (*RealEstateObjectType) TableName() string {
	return "real_estate_object_types"
}

type ERealEstateObjectType int

const (
	KRealEstateObjectTypeUnknownID ERealEstateObjectType = iota
	KRealEstateObjectTypeApartmentID
	KRealEstateObjectTypeHouseID
	KRealEstateObjectTypeLandID
)
