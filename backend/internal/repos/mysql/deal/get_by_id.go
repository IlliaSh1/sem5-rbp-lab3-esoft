package repos_mysql_deal

import (
	"errors"

	"github.com/IlliaSh1/backend/internal/models"
	"gorm.io/gorm"
)

var ErrDealNotFound = errors.New("deal not found")

func (dealRepo *DealRepo) GetByID(id int) (*models.Deal, error) {
	var deal models.Deal
	err := dealRepo.db.First(&deal, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrDealNotFound
		}
		return nil, err
	}

	return &deal, nil
}
