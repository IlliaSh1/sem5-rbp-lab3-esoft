package repos_mysql_real_estate_object

import "github.com/IlliaSh1/backend/internal/models"

func (repo *RealEstateObjectRepo) Delete(id int) (err error) {
	err = repo.db.Delete(&models.Apartment{}, id).Error
	if err != nil {
		return err
	}
	err = repo.db.Delete(&models.House{}, id).Error
	if err != nil {
		return err
	}
	err = repo.db.Delete(&models.Land{}, id).Error
	if err != nil {
		return err
	}

	err = repo.db.Delete(&models.RealEstateObject{}, id).Error
	if err != nil {
		return err
	}

	return nil
}
