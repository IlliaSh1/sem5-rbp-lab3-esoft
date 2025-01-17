package repos_mysql_need

import "github.com/IlliaSh1/backend/internal/models"

func (repo *NeedRepo) SaveLandNeed(land_need *models.LandNeed) error {
	repo.m.Lock()
	defer repo.m.Unlock()

	err := repo.db.Save(&land_need).Error
	if err != nil {
		return err
	}

	return nil
}
