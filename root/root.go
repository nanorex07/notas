package root

import (
	"github.com/charmbracelet/bubbles/help"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/nanorex07/notas/dashboard"
	edit "github.com/nanorex07/notas/edit"
	settings "github.com/nanorex07/notas/settings"
	"github.com/nanorex07/notas/types"
)

type RootModel struct {
	appSettings   *settings.AppSettings
	dashboardView dashboard.DashboardModel
	editView      edit.EditModel
}

func NewRootModel() RootModel {

	helpModel := help.New()

	appSettings := &settings.AppSettings{
		HelpModel:    &helpModel,
		EditViewMode: types.CreateMode,
	}

	return RootModel{
		appSettings:   appSettings,
		dashboardView: dashboard.NewDashboardModel(appSettings),
		editView:      edit.NewEditModel(appSettings),
	}
}

func (m RootModel) Init() tea.Cmd {
	return nil
}

func (m RootModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	cmds := make([]tea.Cmd, 0)
	switch m.appSettings.ActiveView {

	case types.DashboardView:
		updatedModel, cmd := m.dashboardView.Update(msg)
		m.dashboardView = updatedModel.(dashboard.DashboardModel)
		cmds = append(cmds, cmd)

	case types.EditView:
		updatedModel, cmd := m.editView.Update(msg)
		m.editView = updatedModel.(edit.EditModel)
		cmds = append(cmds, cmd)
	}

	return m, tea.Batch(cmds...)
}
