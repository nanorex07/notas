package root

import (
	"github.com/charmbracelet/bubbles/help"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/nanorex07/notas/dashboard"
	settings "github.com/nanorex07/notas/settings"
)

type RootModel struct {
	ActiveView
	appSettings   *settings.AppSettings
	dashboardView dashboard.DashboardModel
}

func NewRootModel() RootModel {

	helpModel := help.New()

	appSettings := &settings.AppSettings{
		HelpModel: &helpModel,
	}

	return RootModel{
		ActiveView:    DashboardView,
		appSettings:   appSettings,
		dashboardView: dashboard.NewDashboardModel(appSettings),
	}
}

func (m RootModel) Init() tea.Cmd {
	return nil
}

func (m RootModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	updatedModel, cmd := m.ActiveViewInstance().Update(msg)
	m.UpdateActionViewInstance(updatedModel)
	return m, cmd
}
