package repos_mysql_user

import (
	"errors"

	"github.com/IlliaSh1/backend/internal/models"
	"gorm.io/gorm"
)

var ErrUserNotFound = errors.New("user not found")

func (s *UserRepo) GetById(id int) (*models.User, error) {
	user := &models.User{}
	err := s.db.First(&user, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}

	return user, nil
}
