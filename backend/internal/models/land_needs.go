package models

type LandNeed struct {
	NeedID int   `json:"need_id" gorm:"column:need_id;primary_key"`
	Need   *Need `json:"need" gorm:"foreignKey:NeedID"`

	MinArea *uint `json:"min_area" gorm:"column:min_area"`
	MaxArea *uint `json:"max_area" gorm:"column:max_area"`
}

func (*LandNeed) TableName() string {
	return "land_needs"
}
