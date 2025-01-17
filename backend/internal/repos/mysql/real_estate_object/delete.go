package repos_mysql_real_estate_object

import "github.com/IlliaSh1/backend/internal/models"

func (repo *RealEstateObjectRepo) Delete(id int) error {
	err := repo.db.Delete(&models.RealEstateObject{}, id).Error
	if err != nil {
		return err
	}

	return nil
}
