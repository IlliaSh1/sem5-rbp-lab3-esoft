package models

type HouseNeed struct {
	NeedID int   `json:"need_id" gorm:"column:need_id;primary_key"`
	Need   *Need `json:"need" gorm:"foreignKey:NeedID"`

	MinFloorsCount *uint `json:"min_floors_count" gorm:"column:min_floors_count"`
	MaxFloorsCount *uint `json:"max_floors_count" gorm:"column:max_floors_count"`
	MinRoomsCount  *uint `json:"min_rooms_count" gorm:"column:min_rooms_count"`
	MaxRoomsCount  *uint `json:"max_rooms_count" gorm:"column:max_rooms_count"`
	MinArea        *uint `json:"min_area" gorm:"column:min_area"`
	MaxArea        *uint `json:"max_area" gorm:"column:max_area"`
}

func (HouseNeed) TableName() string {
	return "house_needs"
}
