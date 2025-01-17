package repos_mysql_offer

import (
	"sync"

	"gorm.io/gorm"
)

type OfferRepo struct {
	db *gorm.DB
	m  *sync.RWMutex
}

func NewOfferRepo(db *gorm.DB) *OfferRepo {
	return &OfferRepo{
		db: db,
		m:  &sync.RWMutex{},
	}
}
