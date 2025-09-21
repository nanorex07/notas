package dashboard

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/nanorex07/notas/components"
	settings "github.com/nanorex07/notas/settings"
	"github.com/nanorex07/notas/types"
)

type DashboardModel struct {
	keys        keyBindings
	list        *components.ListComponent
	appSettings *settings.AppSettings
}

func NewDashboardModel(appSettings *settings.AppSettings) DashboardModel {
	list := components.NewListComponent()
	list.AddItem(&components.ListItem{
		ID:    "1",
		Title: "Item 1",
		Body:  "Resolving deltas: 100% (1/1), done.",
		Tags:  []string{"tag1", "tag2"},
	})
	list.AddItem(&components.ListItem{
		ID:    "2",
		Title: "Item 2",
		Body:  "All done!",
		Tags:  []string{"tag1", "tag2"},
	})
	return DashboardModel{
		keys:        dashboardKeys,
		list:        list,
		appSettings: appSettings,
	}
}

func (m DashboardModel) Init() tea.Cmd {
	return nil
}

func (m DashboardModel) KeyBindings() help.KeyMap {
	return m.keys
}

func (m DashboardModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	cmds := make([]tea.Cmd, 0)
	var cmd tea.Cmd

	model, cmd := m.appSettings.HelpModel.Update(msg)
	m.appSettings.HelpModel = &model
	cmds = append(cmds, cmd)

	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		m.appSettings.HelpModel.Width = msg.Width

	case tea.KeyMsg:
		switch {
		case key.Matches(msg, dashboardKeys.Help):
			m.appSettings.HelpModel.ShowAll = !m.appSettings.HelpModel.ShowAll
		case key.Matches(msg, dashboardKeys.Quit):
			return m, tea.Quit

		case key.Matches(msg, dashboardKeys.Up):
			m.list.Up()
		case key.Matches(msg, dashboardKeys.Down):
			m.list.Down()
		case key.Matches(msg, dashboardKeys.New):
			m.appSettings.EditViewMode = types.CreateMode
			m.appSettings.ActiveView = types.EditView
		}
	}
	return m, tea.Batch(cmds...)
}
