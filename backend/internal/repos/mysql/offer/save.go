package repos_mysql_offer

import "github.com/IlliaSh1/backend/internal/models"

func (repo *OfferRepo) Save(offer *models.Offer) (*models.Offer, error) {
	err := repo.db.Save(&offer).Error
	if err != nil {
		return nil, err
	}

	return offer, nil
}
