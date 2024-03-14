package main

import (
    "fmt"
    "os"

    tea "github.com/charmbracelet/bubbletea"
)

type model struct {
    message string
}

func (m model) Init() tea.Cmd {
    // Return nil as we do not need to initialize anything further
    return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    // Handle application quit when 'q' is pressed
    switch msg := msg.(type) {
    case tea.KeyMsg:
        switch msg.String() {
            case "q":
                return m, tea.Quit
        }
    }
    return m, nil // Return the model as-is if no updates
}

func (m model) View() string {
    // This function describes how to render the model to the screen.
    return fmt.Sprintf("Hello, Bubble Tea! Press 'q' to quit.\n")
}

func main() {
    initialModel := model{message: "Hello, Bubble Tea!"}

    p := tea.NewProgram(initialModel)
    if _, err := p.Run(); err != nil {
        fmt.Fprintf(os.Stderr, "Could not start the Bubble Tea program: %v", err)
        os.Exit(1)
    }
}

