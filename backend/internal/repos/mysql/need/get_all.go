package repos_mysql_need

import "github.com/IlliaSh1/backend/internal/models"

func (repo *NeedRepo) GetAll() ([]*models.Need, error) {
	var needs []*models.Need
	err := repo.db.Find(&needs).Error
	if err != nil {
		return nil, err
	}

	return needs, nil
}
