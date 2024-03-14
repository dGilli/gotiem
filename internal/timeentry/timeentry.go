package timeentry

import (
	"database/sql"
	"fmt"
	"time"
)

type TimeEntry struct {
	ID        int64
	StartTime time.Time
    StopTime  time.Time
    Project   string
    Tags      []string
}

func NewTimeEntry(timeStr string) (TimeEntry, error) {
    time, err := time.Parse(time.DateTime, timeStr)
    if err != nil {
        return TimeEntry{}, fmt.Errorf("could not parse time: %w", err)
    }
    
    entry := TimeEntry{
        StartTime: time,
    }

    return entry, nil
}

func WriteTimeEntry(db *sql.DB, entry TimeEntry) (int64, error) {
	result, err := db.Exec("insert into time_entries (time) values (?)", entry.StartTime.Format(time.DateTime))
	if err != nil {
		return 0, fmt.Errorf("could not insert time entry: %w", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("could not retrieve last insert ID: %w", err)
	}
	return id, nil
}

