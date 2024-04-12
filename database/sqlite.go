package database

import (
	"database/sql"
	_ "modernc.org/sqlite"
)

type sqliteDatabase struct {
	Db *sql.DB
}

func NewSqlliteDatabase(pathToDbFile string) *sqliteDatabase {
	db, err := sql.Open("sqlite", pathToDbFile)
	if err != nil {
		panic("failed to connect to database")
	}

	return &sqliteDatabase{Db: db}
}
