package repos_mysql_realtor

import "github.com/IlliaSh1/backend/internal/models"

func (repo *RealtorRepo) Save(realtor *models.Realtor) (*models.Realtor, error) {
	err := repo.db.Save(&realtor).Error
	if err != nil {
		return nil, err
	}

	return realtor, nil
}
