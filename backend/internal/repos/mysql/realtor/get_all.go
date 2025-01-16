package repos_mysql_realtor

import "github.com/IlliaSh1/backend/internal/models"

func (repo *RealtorRepo) GetAll() ([]*models.Realtor, error) {
	var realtors []*models.Realtor
	err := repo.db.Find(&realtors).Error
	if err != nil {
		return nil, err
	}

	return realtors, nil
}
