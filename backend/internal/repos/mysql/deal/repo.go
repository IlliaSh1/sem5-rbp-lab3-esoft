package repos_mysql_deal

import (
	"sync"

	"gorm.io/gorm"
)

type DealRepo struct {
	db *gorm.DB
	m  *sync.RWMutex
}

func NewDealRepo(db *gorm.DB) *DealRepo {
	return &DealRepo{
		db: db,
		m:  &sync.RWMutex{},
	}
}
