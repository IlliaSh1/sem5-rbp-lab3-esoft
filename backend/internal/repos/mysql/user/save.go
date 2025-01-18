package repos_mysql_user

import "github.com/IlliaSh1/backend/internal/models"

func (s *UserRepo) Save(user *models.User) error {
	s.m.Lock()
	defer s.m.Unlock()
	return s.db.Save(user).Error
}
