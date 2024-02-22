package keys

import (
	tea "github.com/charmbracelet/bubbletea"
	//	"github.com/charmbracelet/lipgloss"
	"github.com/jedib0t/go-pretty/v6/table"
)

//var (
//	titleStyle = lipgloss.NewStyle().
//		Foreground(lipgloss.Color("#FFFDF5")).
//		Background(lipgloss.Color("#25A065")).
//		Padding(0, 1)
//
//	statusMessageStyle = lipgloss.NewStyle().
//		Foreground(lipgloss.AdaptiveColor{Light: "#04B575", Dark: "#04B575"}).
//		Render
//)

type History []string

func NewHistory() History {
	return []string{}
}

func (history History) View() string {
	// TODO(ramses): pre-set table height so we avoid all this movement or use the list util idk 
	status := "Listening for keys..."
	if len(history) > 1 {
		tbl := table.NewWriter()
		tbl.AppendHeader(table.Row{"#", "String Repr."})

		for i, entry := range history {
			tbl.AppendRow(table.Row{len(history) - i, entry})
		}

		status = tbl.Render()
	}

	return status
}

func (history History) Add(msg tea.Msg) History {
	return append([]string{msg.(tea.KeyMsg).String()}, history...)
}
