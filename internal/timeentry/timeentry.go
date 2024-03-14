package timeentry

import (
	"database/sql"
	"fmt"
)

type TimeEntry struct {
	ID   int64
	Time string
}

func WriteTimeEntry(db *sql.DB, entry TimeEntry) (int64, error) {
	result, err := db.Exec("insert into time_entries (time) values (?)", entry.Time)
	if err != nil {
		return 0, fmt.Errorf("could not insert time entry: %w", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("could not retrieve last insert ID: %w", err)
	}
	return id, nil
}
