package repos_mysql_client

import "github.com/IlliaSh1/backend/internal/models"

func (repo *ClientRepo) Delete(id int) error {
	return repo.db.Delete(&models.Client{}, id).Error
}
