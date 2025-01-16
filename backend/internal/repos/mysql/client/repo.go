package repos_mysql_client

import (
	"sync"

	"gorm.io/gorm"
)

type ClientRepo struct {
	db *gorm.DB
	m  *sync.RWMutex
}

func NewClientRepo(db *gorm.DB) *ClientRepo {
	return &ClientRepo{
		db: db,
		m:  &sync.RWMutex{},
	}
}
