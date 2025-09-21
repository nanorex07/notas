package root

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
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

func (m RootModel) ActiveViewInstance() types.ViewInstance {
	switch m.appSettings.ActiveView {
	case types.DashboardView:
		return m.dashboardView
	case types.EditView:
		return m.editView
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

	helpSection := m.appSettings.HelpModel.View(activeViewInstance.KeyBindings())
	viewBuilder.WriteString(helpSection)

	return globalStyle.Render(viewBuilder.String())

}
