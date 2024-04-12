package database

import "database/sql"

func (s *TaskStore) Delete(id int) error {
	_, err := s.Db.Exec("DELETE FROM scheduler WHERE id = :id;", sql.Named("id", id))
	return err
}
