package repos_mysql_need

import (
	"errors"

	"github.com/IlliaSh1/backend/internal/models"
	"gorm.io/gorm"
)

var (
	ErrNeedNotFound = errors.New("need not found")
)

func (repo *NeedRepo) GetByID(id int) (*models.Need, error) {
	need := &models.Need{}
	err := repo.db.First(&need, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrNeedNotFound
		}
		return nil, err
	}

	return need, nil
}
