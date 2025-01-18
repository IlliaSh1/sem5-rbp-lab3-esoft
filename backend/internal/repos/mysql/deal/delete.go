package repos_mysql_deal

import "github.com/IlliaSh1/backend/internal/models"

func (dealRepo *DealRepo) Delete(id int) error {
	return dealRepo.db.Delete(&models.Deal{}, id).Error
}
