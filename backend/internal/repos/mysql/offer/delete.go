package repos_mysql_offer

import "github.com/IlliaSh1/backend/internal/models"

func (repo *OfferRepo) Delete(id int) error {
	err := repo.db.Delete(&models.Offer{}, id).Error
	if err != nil {
		return err
	}

	return nil
}
