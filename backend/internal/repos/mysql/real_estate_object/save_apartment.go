package repos_mysql_real_estate_object

import (
	"context"

	"github.com/IlliaSh1/backend/internal/models"
	"github.com/IlliaSh1/backend/internal/storage/transactor"
	"gorm.io/gorm"
)

func (repo *RealEstateObjectRepo) SaveApartment(ctx context.Context, apartment *models.Apartment) (*models.Apartment, error) {
	var db *gorm.DB = transactor.GetTransaction(ctx)
	if db == nil {
		db = repo.db
	}

	err := db.Save(&apartment).Error
	if err != nil {
		return nil, err
	}

	return apartment, nil
}
