package edit

import "github.com/charmbracelet/bubbles/key"

type keyBindings struct {
	Save key.Binding
	View key.Binding
	Quit key.Binding
	Help key.Binding
}

func (k keyBindings) ShortHelp() []key.Binding {
	return []key.Binding{k.Quit, k.Help}
}

func (k keyBindings) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Save, k.View},
		{k.Quit, k.Help},
	}
}

var editKeys = keyBindings{
	Save: key.NewBinding(key.WithKeys("s"), key.WithHelp("s", "save")),
	View: key.NewBinding(key.WithKeys("v"), key.WithHelp("v", "view")),
	Quit: key.NewBinding(key.WithKeys("q"), key.WithHelp("q", "quit")),
	Help: key.NewBinding(key.WithKeys("h"), key.WithHelp("h", "help")),
}
