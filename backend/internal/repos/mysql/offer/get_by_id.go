package repos_mysql_offer

import (
	"errors"

	"github.com/IlliaSh1/backend/internal/models"
	"gorm.io/gorm"
)

var ErrOfferNotFound = errors.New("offer not found")

func (repo *OfferRepo) GetByID(id int) (*models.Offer, error) {
	offer := &models.Offer{}
	err := repo.db.First(&offer, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrOfferNotFound
		}
		return nil, err
	}

	return offer, nil
}
