package repos_mysql_need

import "github.com/IlliaSh1/backend/internal/models"

func (s *NeedRepo) Save(need *models.Need) error {
	s.m.Lock()
	defer s.m.Unlock()

	err := s.db.Save(&need).Error
	if err != nil {
		return err
	}

	return nil
}
