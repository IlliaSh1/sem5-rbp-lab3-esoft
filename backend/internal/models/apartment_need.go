package models

type ApartmentNeed struct {
	NeedID int   `json:"need_id" gorm:"column:need_id;primary_key"`
	Need   *Need `json:"need" gorm:"foreignKey:NeedID"`

	MinFloorNumber *int `json:"min_floor_number" gorm:"column:min_floor_number"`
	MaxFloorNumber *int `json:"max_floor_number" gorm:"column:max_floor_number"`

	MinRoomsCount *uint `json:"min_rooms_count" gorm:"column:min_rooms_count"`
	MaxRoomsCount *uint `json:"max_rooms_count" gorm:"column:max_rooms_count"`

	MinArea *uint `json:"min_area" gorm:"column:min_area"`
	MaxArea *uint `json:"max_area" gorm:"column:max_area"`
}

func (ApartmentNeed) TableName() string {
	return "apartment_needs"
}
