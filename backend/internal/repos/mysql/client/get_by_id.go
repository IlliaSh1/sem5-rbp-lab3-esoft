package repos_mysql_client

import (
	"errors"

	"github.com/IlliaSh1/backend/internal/models"
)

var (
	ErrClientNotFound = errors.New("client not found")
)

func (repo *ClientRepo) GetById(id int) (*models.Client, error) {
	var client *models.Client
	err := repo.db.First(&client, id).Error
	if err != nil {
		return nil, err
	}
	if client == nil {
		return nil, ErrClientNotFound
	}

	return client, nil
}
