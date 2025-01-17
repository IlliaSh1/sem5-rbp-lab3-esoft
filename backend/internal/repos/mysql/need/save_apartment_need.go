package repos_mysql_need

import "github.com/IlliaSh1/backend/internal/models"

func (repo *NeedRepo) SaveApartmentNeed(apartment_need *models.ApartmentNeed) error {
	repo.m.Lock()
	defer repo.m.Unlock()

	err := repo.db.Save(&apartment_need).Error
	if err != nil {
		return err
	}

	return nil
}
