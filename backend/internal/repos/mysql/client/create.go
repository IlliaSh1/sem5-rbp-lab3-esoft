package repos_mysql_client

import "github.com/IlliaSh1/backend/internal/models"

func (repo *ClientRepo) Create(client *models.Client) (*models.Client, error) {
	err := repo.db.Create(&client).Error
	if err != nil {
		return nil, err
	}

	return client, nil
}
