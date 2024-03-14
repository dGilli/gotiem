package timeentry

import (
	"testing"

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
    testEntry := TimeEntry{Time: "foobar"}
    if _, err := WriteTimeEntry(testDB, testEntry); err != nil {
        t.Fatalf("Failed to write time entry: %s", err)
    }

	// Verify the entry was written correctly
	var retrievedTime string
	err = testDB.QueryRow("select time from time_entries order by id desc limit 1").Scan(&retrievedTime)
	if err != nil {
		t.Fatalf("Failed to retrieve time entry: %s", err)
	}

	if retrievedTime != testEntry.Time {
		t.Errorf("Expected to retrieve entry '%s', got '%s'", testEntry.Time, retrievedTime)
	}
}

