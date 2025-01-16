package repos_mysql_realtor

import (
	"errors"

	"github.com/IlliaSh1/backend/internal/models"
)

var (
	ErrRealtorNotFound = errors.New("realtor not found")
)

func (repo *RealtorRepo) GetByID(id int) (*models.Realtor, error) {
	realtor := &models.Realtor{}
	err := repo.db.First(&realtor, id).Error
	if err != nil {
		return nil, err
	}

	if realtor == nil {
		return nil, ErrRealtorNotFound
	}

	return realtor, nil
}
