package repos_mysql_real_estate_object

import (
	"errors"

	"github.com/IlliaSh1/backend/internal/models"
)

var ErrRealEstateObjectNotFound = errors.New("real estate object not found")

func (repo *RealEstateObjectRepo) GetById(id int) (*models.RealEstateObject, error) {
	realEstateObject := &models.RealEstateObject{}
	err := repo.db.First(&realEstateObject, id).Error
	if err != nil {
		return nil, err
	}

	if realEstateObject == nil {
		return nil, ErrRealEstateObjectNotFound
	}

	return realEstateObject, nil
}
