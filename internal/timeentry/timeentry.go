package timeentry

import (
	"database/sql"
	"fmt"
	"time"
)

type TimeEntry struct {
	ID   int64
	Time time.Time
}

func NewTimeEntry(timeStr string) (TimeEntry, error) {
    time, err := time.Parse(time.DateTime, timeStr)
    if err != nil {
        return TimeEntry{}, fmt.Errorf("could not parse time: %w", err)
    }
    
    entry := TimeEntry{
        Time: time,
    }

    return entry, nil
}

func WriteTimeEntry(db *sql.DB, entry TimeEntry) (int64, error) {
	result, err := db.Exec("insert into time_entries (time) values (?)", entry.Time.Format(time.DateTime))
	if err != nil {
		return 0, fmt.Errorf("could not insert time entry: %w", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("could not retrieve last insert ID: %w", err)
	}
	return id, nil
}

