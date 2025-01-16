package repos_mysql_client

import "github.com/IlliaSh1/backend/internal/models"

func (repo *ClientRepo) GetAll() ([]*models.Client, error) {
	var clients []*models.Client
	err := repo.db.Find(&clients).Error
	if err != nil {
		return nil, err
	}

	return clients, nil
}
