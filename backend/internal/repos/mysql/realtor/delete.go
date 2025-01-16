package repos_mysql_realtor

import "github.com/IlliaSh1/backend/internal/models"

func (repo *RealtorRepo) Delete(id int) error {
	return repo.db.Delete(&models.Realtor{}, id).Error
}
