package repos_mysql_realtor

import (
	"sync"

	"gorm.io/gorm"
)

type RealtorRepo struct {
	db *gorm.DB
	m  *sync.RWMutex
}

func NewRealtorRepo(db *gorm.DB) *RealtorRepo {
	return &RealtorRepo{
		db: db,
		m:  &sync.RWMutex{},
	}
}
