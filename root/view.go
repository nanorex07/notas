package root

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/nanorex07/notas/dashboard"
	"github.com/nanorex07/notas/types"
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

func (m RootModel) ActiveViewInstance() types.ViewInstance {
	switch m.ActiveView {
	case DashboardView:
		return m.dashboardView
	}
	return nil
}

func (m RootModel) UpdateActionViewInstance(model tea.Model) {
	switch m.ActiveView {
	case DashboardView:
		m.dashboardView = model.(dashboard.DashboardModel)
	}
}

func (m RootModel) View() string {

	viewBuilder := strings.Builder{}
	activeViewInstance := m.ActiveViewInstance()

	title := "♡ ♥ Notas App ♥ ♡"
	viewBuilder.WriteString(titleStyle.Render(title))
	viewBuilder.WriteString("\n\n")

	viewBuilder.WriteString(activeViewInstance.View())
	viewBuilder.WriteString("\n\n")

	helpSection := m.appSettings.HelpModel.View(activeViewInstance.KeyBindings())
	viewBuilder.WriteString(helpSection)

	return globalStyle.Render(viewBuilder.String())

}
