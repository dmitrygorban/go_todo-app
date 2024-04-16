package storage

import (
	"database/sql"
	"os"
	"path/filepath"
	"time"

	"github.com/dmitrygorban/go_todo-app/models"
	"github.com/dmitrygorban/go_todo-app/scheduler"
	_ "modernc.org/sqlite"
)

type TaskStore struct {
	Db *sql.DB
}

func NewTaskStore(pathToDatabaseFile string) TaskStore {
	db, err := sql.Open("sqlite", pathToDatabaseFile)
	if err != nil {
		panic("failed to connect to database")
	}

	return TaskStore{Db: db}
}

func (s *TaskStore) Get(filter string) ([]models.Task, error) {
	var tasks []models.Task
	var rows *sql.Rows
	var err error

	if filter == "" {
		query := "SELECT id, date, title, comment, repeat FROM scheduler ORDER BY DATE;"
		rows, err = s.Db.Query(query)
	}

	if filter != "" {
		var query string
		var searchParam interface{}

		date, err := time.Parse("02.01.2006", filter)
		if err == nil {
			query = "SELECT id, date, title, comment, repeat FROM scheduler WHERE date LIKE :search ORDER BY DATE;"
			searchParam = date.Format(scheduler.DATE_FORMAT)
		} else {
			query = "SELECT id, date, title, comment, repeat FROM scheduler WHERE title LIKE :search OR comment LIKE :search ORDER BY DATE;"
			searchParam = "%" + filter + "%"
		}

		rows, err = s.Db.Query(query, sql.Named("search", searchParam))
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

func (s *TaskStore) GetTaskById(id int) (models.Task, error) {
	task := models.Task{}

	row := s.Db.QueryRow("SELECT id, date, title, comment, repeat FROM scheduler WHERE id = :id;",
		sql.Named("id", id))

	err := row.Scan(&task.Id, &task.Date, &task.Title, &task.Comment, &task.Repeat)
	if err != nil {
		return task, err
	}

	return task, nil
}

func (s *TaskStore) Add(t models.Task) (int, error) {
	res, err := s.Db.Exec("INSERT INTO scheduler (date, title, comment, repeat) VALUES (:date, :title, :comment, :repeat)",
		sql.Named("date", t.Date),
		sql.Named("title", t.Title),
		sql.Named("comment", t.Comment),
		sql.Named("repeat", t.Repeat),
	)

	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (s *TaskStore) Delete(id int) error {
	_, err := s.Db.Exec("DELETE FROM scheduler WHERE id = :id;", sql.Named("id", id))
	return err
}

func (s *TaskStore) Update(task models.Task) error {
	_, err := s.Db.Exec("UPDATE scheduler SET date = :date, title = :title, comment = :comment, repeat = :repeat WHERE id = :id",
		sql.Named("date", task.Date),
		sql.Named("title", task.Title),
		sql.Named("comment", task.Comment),
		sql.Named("repeat", task.Repeat),
		sql.Named("id", task.Id),
	)
	if err != nil {
		return err
	}

	return nil
}

func (s *TaskStore) NeedMigration(pathToFile string) bool {
	appPath, err := os.Getwd()

	if err != nil {
		return false
	}

	dbFile := filepath.Join(appPath, pathToFile)
	_, err = os.Stat(dbFile)

	return os.IsNotExist(err)
}
