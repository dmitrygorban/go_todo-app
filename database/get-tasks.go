package database

import (
	"database/sql"
	"time"

	"github.com/dmitrygorban/go_todo-app/models"
	"github.com/dmitrygorban/go_todo-app/scheduler"
)

func (s *TaskStore) Get(searchQueryParam string) ([]models.Task, error) {
	var tasks []models.Task
	var rows *sql.Rows
	var err error

	if searchQueryParam != "" {
		date, err := time.Parse("02.01.2006", searchQueryParam)
		if err == nil {
			rows, err = s.Db.Query("SELECT id, date, title, comment, repeat FROM scheduler WHERE date LIKE :search ORDER BY DATE;",
				sql.Named("search", date.Format(scheduler.DATE_FORMAT)))
		} else {
			rows, err = s.Db.Query("SELECT id, date, title, comment, repeat FROM scheduler WHERE title LIKE :search OR comment LIKE :search ORDER BY DATE;",
				sql.Named("search", "%"+searchQueryParam+"%"),
			)
		}
	} else {
		query := "SELECT id, date, title, comment, repeat FROM scheduler ORDER BY DATE;"
		rows, err = s.Db.Query(query)
	}

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

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return tasks, err
}
