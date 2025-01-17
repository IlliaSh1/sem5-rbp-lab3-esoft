package repos_mysql_client

import (
	"errors"

	"github.com/IlliaSh1/backend/internal/models"
	"gorm.io/gorm"
)

var (
	ErrClientNotFound = errors.New("client not found")
)

func (repo *ClientRepo) GetById(id int) (*models.Client, error) {
	var client *models.Client
	err := repo.db.First(&client, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrClientNotFound
		}
		return nil, err
	}

	return client, nil
}
