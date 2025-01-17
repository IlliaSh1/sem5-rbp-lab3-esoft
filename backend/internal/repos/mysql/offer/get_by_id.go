package repos_mysql_offer

import (
	"errors"

	"github.com/IlliaSh1/backend/internal/models"
)

var ErrOfferNotFound = errors.New("offer not found")

func (repo *OfferRepo) GetByID(id int) (*models.Offer, error) {
	offer := &models.Offer{}
	err := repo.db.First(&offer, id).Error
	if err != nil {
		return nil, err
	}

	if offer == nil {
		return nil, ErrOfferNotFound
	}

	return offer, nil
}
