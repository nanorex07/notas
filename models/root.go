package models

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

type RootModel struct {
	ActiveView
	Help help.Model
}

func NewRootModel() RootModel {
	return RootModel{
		ActiveView: DashboardView,
		Help:       help.New(),
	}
}

func (m RootModel) Init() tea.Cmd {
	return nil
}

func (m RootModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	cmds := make([]tea.Cmd, 0)
	var cmd tea.Cmd

	m.Help, cmd = m.Help.Update(msg)
	cmds = append(cmds, cmd)

	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		m.Help.Width = msg.Width

	case tea.KeyMsg:
		switch {
		case key.Matches(msg, helpKey):
			m.Help.ShowAll = !m.Help.ShowAll
		case key.Matches(msg, quitKey):
			return RootModel{}, tea.Quit
		}
	}

	return m, tea.Batch(cmds...)
}
