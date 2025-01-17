package repos_mysql_need

import "github.com/IlliaSh1/backend/internal/models"

func (repo *NeedRepo) SaveHouseNeed(house_need *models.HouseNeed) error {
	repo.m.Lock()
	defer repo.m.Unlock()

	err := repo.db.Save(&house_need).Error
	if err != nil {
		return err
	}

	return nil
}
