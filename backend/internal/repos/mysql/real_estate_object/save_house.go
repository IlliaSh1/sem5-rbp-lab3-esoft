package repos_mysql_real_estate_object

import (
	"context"

	"github.com/IlliaSh1/backend/internal/models"
	"github.com/IlliaSh1/backend/internal/storage/transactor"
	"gorm.io/gorm"
)

func (repo *RealEstateObjectRepo) SaveHouse(ctx context.Context, house *models.House) (*models.House, error) {
	var db *gorm.DB = transactor.GetTransaction(ctx)
	if db == nil {
		db = repo.db
	}

	err := db.Save(&house).Error
	if err != nil {
		return nil, err
	}

	return house, nil
}
