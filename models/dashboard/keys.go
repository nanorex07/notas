package dashboard

import "github.com/charmbracelet/bubbles/key"

type keyBindings struct {
	New  key.Binding
	Open key.Binding
	Up   key.Binding
	Down key.Binding
	Quit key.Binding
	Help key.Binding
}

func (k keyBindings) ShortHelp() []key.Binding {
	return []key.Binding{k.Quit, k.Help}
}

func (k keyBindings) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.New, k.Open},
		{k.Up, k.Down},
		{k.Quit, k.Help},
	}
}

var dashboardKeys = keyBindings{
	New:  key.NewBinding(key.WithKeys("n"), key.WithHelp("n", "new")),
	Open: key.NewBinding(key.WithKeys("enter"), key.WithHelp("enter", "open")),
	Up:   key.NewBinding(key.WithKeys("up"), key.WithHelp("up", "↑")),
	Down: key.NewBinding(key.WithKeys("down"), key.WithHelp("down", "↓")),
	Quit: key.NewBinding(key.WithKeys("q"), key.WithHelp("q", "quit")),
	Help: key.NewBinding(key.WithKeys("h"), key.WithHelp("h", "help")),
}
