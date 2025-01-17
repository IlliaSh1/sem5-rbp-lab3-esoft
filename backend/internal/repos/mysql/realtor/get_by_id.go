package repos_mysql_realtor

import (
	"errors"

	"github.com/IlliaSh1/backend/internal/models"
	"gorm.io/gorm"
)

var (
	ErrRealtorNotFound = errors.New("realtor not found")
)

func (repo *RealtorRepo) GetByID(id int) (*models.Realtor, error) {
	realtor := &models.Realtor{}
	err := repo.db.First(&realtor, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrRealtorNotFound
		}
		return nil, err
	}

	return realtor, nil
}
