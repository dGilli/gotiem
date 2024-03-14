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
	db, err := db.OpenDB("example.db")
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

        entry := timeentry.TimeEntry{Time: "foobar"}
		if _, err := timeentry.WriteTimeEntry(db, entry); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %s\n", err)
			continue
		}

		fmt.Println("Time entry saved.")
        break
	}
}

