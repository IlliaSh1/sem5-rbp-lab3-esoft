package repos_mysql_deal

import "github.com/IlliaSh1/backend/internal/models"

func (dealRepo *DealRepo) Save(deal *models.Deal) error {
	return dealRepo.db.Save(&deal).Error
}
