package timeentry

import (
	"testing"
	"time"

	"github.com/dGilli/gotiem/internal/db"

	_ "github.com/mattn/go-sqlite3"
)

func TestWriteTimeEntry(t *testing.T) {
	// Setup: open the test database
	testDB, err := db.OpenDB(":memory:")
	if err != nil {
		t.Fatalf("Failed to open database: %s", err)
	}
	defer testDB.Close()

	// Define a test entry
    timeStr := "2024-11-01 10:00:00"
    testTime, err := time.Parse(time.DateTime, timeStr)
    if err != nil {
        t.Fatalf("Failed to parse time: %s", err)
    }
    testEntry := TimeEntry{
        StartTime: testTime,
    }

    if _, err := WriteTimeEntry(testDB, testEntry); err != nil {
        t.Fatalf("Failed to write time entry: %s", err)
    }

	// Verify the entry was written correctly
	var retrievedTime string
	err = testDB.QueryRow("select time from time_entries order by id desc limit 1").Scan(&retrievedTime)
	if err != nil {
		t.Fatalf("Failed to retrieve time entry: %s", err)
	}

	if retrievedTime != timeStr {
		t.Errorf("Expected to retrieve entry '%s', got '%s'", timeStr, retrievedTime)
	}
}

