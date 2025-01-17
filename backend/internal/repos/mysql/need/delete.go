package repos_mysql_need

import "github.com/IlliaSh1/backend/internal/models"

func (repo *NeedRepo) Delete(id int) (err error) {
	err = repo.db.Delete(&models.HouseNeed{}, id).Error
	if err != nil {
		return err
	}

	err = repo.db.Delete(&models.LandNeed{}, id).Error
	if err != nil {
		return err
	}

	err = repo.db.Delete(&models.ApartmentNeed{}, id).Error
	if err != nil {
		return err
	}

	err = repo.db.Delete(&models.Need{}, id).Error
	if err != nil {
		return err
	}

	return nil
}
