package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/dGilli/gotiem/internal/db"
	"github.com/dGilli/gotiem/internal/timeentry"
)

func main() {
	db, err := db.OpenDB("tmp/example.db")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Cannot open database: %s\n", err)
		return
	}
	defer db.Close()

	fmt.Println("Enter a time entry (type 'exit' to quit):")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()
		if strings.ToLower(text) == "exit" {
			break
		}

		fmt.Printf("You typed: %s\n", text)

        entry, err := timeentry.NewTimeEntry("2024-11-01 10:00:00")
        if err != nil {
            fmt.Fprintf(os.Stderr, "Error: %s\n", err)
            continue
        }

		if _, err := timeentry.WriteTimeEntry(db, entry); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %s\n", err)
			continue
		}

		fmt.Println("Time entry saved.")
        break
	}
}

