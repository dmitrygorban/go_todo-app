package database

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type sqliteDatabase struct {
	Db *sql.DB
}

func NewSqlliteDatabase(pathToDbFile string) *sqliteDatabase {
	db, err := sql.Open("sqlite3", pathToDbFile)
	if err != nil {
		panic("failed to connect to database")
	}

	return &sqliteDatabase{Db: db}
}
