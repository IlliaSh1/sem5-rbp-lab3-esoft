package repos_mysql_offer

import "github.com/IlliaSh1/backend/internal/models"

func (repo *OfferRepo) GetAll() ([]*models.Offer, error) {
	var offers []*models.Offer
	err := repo.db.Find(&offers).Error
	if err != nil {
		return nil, err
	}

	return offers, nil
}
