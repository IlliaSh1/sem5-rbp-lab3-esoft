package repos_mysql_user

import (
	"sync"

	"gorm.io/gorm"
)

type UserRepo struct {
	db *gorm.DB
	m  *sync.RWMutex
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{
		db: db,
		m:  &sync.RWMutex{},
	}
}
