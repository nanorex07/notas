package dashboard

import (
	"github.com/charmbracelet/bubbles/help"
	tea "github.com/charmbracelet/bubbletea"
)

type DashboardModel struct {
	keys keyBindings
}

func NewDashboardModel() DashboardModel {
	return DashboardModel{
		keys: dashboardKeys,
	}
}

func (m DashboardModel) Init() tea.Cmd {
	return nil
}

func (m DashboardModel) KeyBindings() help.KeyMap {
	return m.keys
}

func (m DashboardModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}
