package edit

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/nanorex07/notas/components"
	settings "github.com/nanorex07/notas/settings"
)

type EditModel struct {
	keys        keyBindings
	appSettings *settings.AppSettings
	heading     *components.ModelHeading
}

func NewEditModel(appSettings *settings.AppSettings) EditModel {
	return EditModel{
		keys:        editKeys,
		appSettings: appSettings,
		heading:     components.NewModelHeading(""),
	}
}

func (m EditModel) Init() tea.Cmd {
	return nil
}

func (m EditModel) KeyBindings() help.KeyMap {
	return m.keys
}

func (m EditModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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
		case key.Matches(msg, editKeys.Help):
			m.appSettings.HelpModel.ShowAll = !m.appSettings.HelpModel.ShowAll
		case key.Matches(msg, editKeys.Quit):
			return m, tea.Quit
		}
	}
	return m, tea.Batch(cmds...)
}
