package database

import (
	"github.com/dmitrygorban/go_todo-app/models"
)

func (s *TaskStore) Get() ([]models.Task, error) {
	var tasks []models.Task

	rows, err := s.Db.Query("SELECT id, date, title, comment, repeat FROM scheduler ORDER BY DATE;")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		task := models.Task{}

		err := rows.Scan(&task.Id, &task.Date, &task.Title, &task.Comment, &task.Repeat)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	return tasks, err

}
