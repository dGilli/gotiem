package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
    openDB("example.db")

    db, err := openDB("example.db")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Cannot open database: %s\n", err)
		return
	}
	defer db.Close()

	fmt.Println("Enter a message (type 'exit' to quit):")

	// Create a new scanner reading from standard input.
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()
		if strings.ToLower(text) == "exit" {
			fmt.Println("Exiting program.")
			break
		}

        // Insert the input text into the database
		_, err := db.Exec("INSERT INTO messages (content) VALUES (?)", text)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to insert into table: %s\n", err)
			continue
		}

		fmt.Printf("Saved to database: %s\n", text)

    }

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading from input: %s\n", err)
	}
}

func openDB(dbFile string) (*sql.DB, error) {
    db, err := sql.Open("sqlite3", dbFile)
    if err != nil {
        return nil, err
    }

    // Create a table if it does not already exist
	createTableSQL := `CREATE TABLE IF NOT EXISTS messages (
		"id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,   
		"content" TEXT
	);`
	_, err = db.Exec(createTableSQL)
	if err != nil {
		return nil, err
	}

    return db, nil
}

