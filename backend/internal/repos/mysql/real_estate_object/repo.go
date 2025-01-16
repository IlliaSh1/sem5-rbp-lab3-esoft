package repos_mysql_real_estate_object

import (
	"sync"

	"gorm.io/gorm"
)

type RealEstateObjectRepo struct {
	db *gorm.DB
	m  *sync.RWMutex
}

func NewRealEstateObjectRepo(db *gorm.DB) *RealEstateObjectRepo {
	return &RealEstateObjectRepo{
		db: db,
		m:  &sync.RWMutex{},
	}
}
