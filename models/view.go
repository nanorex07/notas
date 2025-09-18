package models

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/nanorex07/notas/models/dashboard"
)

var (
	globalStyle = lipgloss.NewStyle().Padding(2, 3)

	titleStyle = lipgloss.NewStyle().
			BorderStyle(lipgloss.Border{
			Top:         "-*",
			Bottom:      "-*",
			Left:        "|",
			Right:       "|",
			TopLeft:     "*",
			TopRight:    "*",
			BottomLeft:  "*",
			BottomRight: "*",
		}).
		BorderForeground(lipgloss.Color("63")).
		Padding(0, 2).Bold(true).Align(lipgloss.Center)
)

type ActiveView int

const (
	DashboardView ActiveView = iota
	EditView
	DisplayView
)

func (v ActiveView) String() string {
	switch v {
	case DashboardView:
		return "Dashboard"
	case EditView:
		return "Edit"
	case DisplayView:
		return "View"
	default:
		return "Unknown"
	}
}

func (m RootModel) ActiveViewInstance() ViewInstance {
	switch m.ActiveView {
	case DashboardView:
		return dashboard.NewDashboardModel()
	}
	return nil
}

func (m RootModel) View() string {

	viewBuilder := strings.Builder{}
	activeViewInstance := m.ActiveViewInstance()

	title := "♡ ♥ Notas App ♥ ♡"
	viewBuilder.WriteString(titleStyle.Render(title))
	viewBuilder.WriteString("\n\n")

	viewBuilder.WriteString(activeViewInstance.View())
	viewBuilder.WriteString("\n\n")

	helpSection := m.Help.View(activeViewInstance.KeyBindings())
	viewBuilder.WriteString(helpSection)

	return globalStyle.Render(viewBuilder.String())

}
