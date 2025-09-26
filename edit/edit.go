package edit

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textarea"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/nanorex07/notas/components"
	settings "github.com/nanorex07/notas/settings"
)

type EditModel struct {
	appSettings *settings.AppSettings
	heading     *components.ModelHeading

	focusedMode  bool
	focusedInput int
	inputs       []textarea.Model
}

func NewEditModel(appSettings *settings.AppSettings) EditModel {
	var inputs []textarea.Model = make([]textarea.Model, 3)

	// title input
	inputs[0] = textarea.New()
	inputs[0].Placeholder = "Title ..."
	inputs[0].SetHeight(1)
	inputs[0].SetWidth(50)
	inputs[0].ShowLineNumbers = false
	inputs[0].CharLimit = 40
	inputs[0].Focus()

	// body input
	inputs[1] = textarea.New()
	inputs[1].Placeholder = "Enter your note content ..."
	inputs[1].SetHeight(15)
	inputs[1].SetWidth(50)
	inputs[1].ShowLineNumbers = false
	inputs[1].CharLimit = 300
	inputs[1].Focus()

	// tags input
	inputs[2] = textarea.New()
	inputs[2].Placeholder = "',' separated tags ..."
	inputs[2].SetHeight(1)
	inputs[2].SetWidth(50)
	inputs[2].ShowLineNumbers = false
	inputs[2].CharLimit = 300

	return EditModel{
		appSettings:  appSettings,
		heading:      components.NewModelHeading(""),
		inputs:       inputs,
		focusedMode:  true,
		focusedInput: 0,
	}
}

func (m EditModel) Init() tea.Cmd {
	return nil
}

func (m EditModel) KeyBindings() help.KeyMap {
	if m.focusedMode {
		return focusedModeKeys
	}
	return unfocusedModeKeys
}

func (m EditModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	cmds := make([]tea.Cmd, 0)
	var cmd tea.Cmd

	if m.focusedMode {
		switch msg := msg.(type) {
		case tea.WindowSizeMsg:
			m.appSettings.HelpModel.Width = msg.Width
		case tea.KeyMsg:
			switch {
			case key.Matches(msg, focusedModeKeys.Esc):
				m.focusedMode = false
				m.inputs[m.focusedInput].Blur()
			case key.Matches(msg, key.NewBinding(key.WithKeys("shift+tab"))):
				m.focusedInput = (m.focusedInput + 1) % len(m.inputs)
			}
			for i := range m.inputs {
				m.inputs[i].Blur()
			}
			if m.focusedMode {
				m.inputs[m.focusedInput].Focus()
			}
		}
	} else {
		switch msg := msg.(type) {
		case tea.WindowSizeMsg:
			m.appSettings.HelpModel.Width = msg.Width
		case tea.KeyMsg:
			switch {
			case key.Matches(msg, unfocusedModeKeys.Edit):
				m.focusedMode = true
				m.focusedInput = 0
				m.inputs[m.focusedInput].Focus()
			case key.Matches(msg, unfocusedModeKeys.Quit):
				return m, tea.Quit
			case key.Matches(msg, unfocusedModeKeys.Help):
				m.appSettings.HelpModel.ShowAll = !m.appSettings.HelpModel.ShowAll
			}
		}
	}

	model, cmd := m.appSettings.HelpModel.Update(msg)
	m.appSettings.HelpModel = &model
	cmds = append(cmds, cmd)

	for i := range m.inputs {
		m.inputs[i], cmd = m.inputs[i].Update(msg)
		cmds = append(cmds, cmd)
	}

	return m, tea.Batch(cmds...)
}
