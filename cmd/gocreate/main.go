package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	choices  []string
	cursor   int
	selected map[int]struct{}
}

func initialModel() model {
	return model{
		choices:  []string{"Multiple commands or packages", "Basic web server", "Basic package"},
		selected: make(map[int]struct{}),
	}
}

func (m model) Init() tea.Cmd {
	return nil // No command to run on startup
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// Handle key presses
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q", "esc":
			return m, tea.Quit
		case "up", "k", "w":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j", "s":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}
		case "enter", " ":
			if _, ok := m.selected[m.cursor]; ok {
				delete(m.selected, m.cursor)
			} else {
				m.selected[m.cursor] = struct{}{}
			}
		}
	}
	return m, nil
}

func (m model) View() string {
	s := "What type of project would you like to create?\n\n"

	// Is the cursor pointing at a choice that's selected?
	for i, choice := range m.choices {
		cursor := " " // no cursor
		if m.cursor == i {
			cursor = ">"
		}
		// Render a checkbox for the selected choice
		checked := " "
		if _, ok := m.selected[i]; ok {
			checked = "x"
		}
		// Render the choice
		s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice)
	}

	s += "\nPress q or esc to quit.\n"
	return s // Return the view string
}

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Could not start program :(\n%v\n", err)
		os.Exit(1)
	}
}
