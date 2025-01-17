package repos_mysql_need

import (
	"sync"

	"gorm.io/gorm"
)

type NeedRepo struct {
	db *gorm.DB
	m  *sync.RWMutex
}

func NewNeedRepo(db *gorm.DB) *NeedRepo {
	return &NeedRepo{
		db: db,
		m:  &sync.RWMutex{},
	}
}
