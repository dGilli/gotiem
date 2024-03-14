package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

func OpenDB(dbFile string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", dbFile)
	if err != nil {
		return nil, err
	}

	createTableSQL := `create table if not exists time_entries (
		"id" integer not null primary key autoincrement,   
		"time" text not null
	);`
	if _, err = db.Exec(createTableSQL); err != nil {
		return nil, err
	}

	return db, nil
}

