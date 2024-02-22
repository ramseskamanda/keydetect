package main

import (
	"fmt"
	"github.com/ramseskamanda/keydetect/keys"
	"os"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	appStyle = lipgloss.NewStyle().Padding(1, 2)
)

type model struct {
	history  keys.History
	bindings keys.KeyMap
	help     help.KeyMap
	quitting bool
}

func newModel() model {
	h := help.New()
	h.ShowAll = true

	m := model{
		bindings: keys.Bindings,
		history:  keys.NewHistory(),
	}

	return m
}

func (m model) Init() tea.Cmd {
	return tea.SetWindowTitle("Key detection")
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		// If we set a width on the help menu it can gracefully truncate
		// its view as needed.
		m.bindings.Width = msg.Width

	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.bindings.Help):
			m.bindings.ShowAll = !m.bindings.ShowAll
		case key.Matches(msg, m.bindings.Quit):
			m.quitting = true
			return m, tea.Quit
		default:
			m.history = m.history.Add(msg)
		}
	}

	return m, nil
}

func (m model) View() string {
	if m.quitting {
		return fmt.Sprintf("Pressed %d keys. \n\nBye!\n", len(m.history))
	}

	return appStyle.Render(
		lipgloss.JoinHorizontal(
			lipgloss.Top,
			m.history.View(),
			m.bindings.View(),
		),
	)
}

func main() {
	if _, err := tea.NewProgram(newModel(), tea.WithAltScreen()).Run(); err != nil {
		fmt.Printf("Could not start program :(\n%v\n", err)
		os.Exit(1)
	}
}
