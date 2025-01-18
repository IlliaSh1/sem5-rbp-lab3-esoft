package repos_mysql_deal

import "github.com/IlliaSh1/backend/internal/models"

func (dealRepo *DealRepo) GetAll() ([]*models.Deal, error) {
	var deals []*models.Deal
	err := dealRepo.db.Find(&deals).Error
	if err != nil {
		return nil, err
	}

	return deals, nil
}
