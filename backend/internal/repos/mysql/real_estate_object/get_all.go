package repos_mysql_real_estate_object

import "github.com/IlliaSh1/backend/internal/models"

func (repo *RealEstateObjectRepo) GetAll() ([]*models.RealEstateObject, error) {
	var realEstateObjects []*models.RealEstateObject
	err := repo.db.Find(&realEstateObjects).Error
	if err != nil {
		return nil, err
	}

	return realEstateObjects, nil
}
