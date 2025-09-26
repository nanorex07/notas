package edit

import (
	"github.com/charmbracelet/bubbles/key"
)

type focusedKeyBindings struct {
	Esc      key.Binding
	ShiftTab key.Binding
}

func (k focusedKeyBindings) ShortHelp() []key.Binding {
	return []key.Binding{k.Esc, k.ShiftTab}
}

func (k focusedKeyBindings) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Esc, k.ShiftTab},
	}
}

var focusedModeKeys = focusedKeyBindings{
	Esc:      key.NewBinding(key.WithKeys("esc"), key.WithHelp("esc", "<- view")),
	ShiftTab: key.NewBinding(key.WithKeys("shift+tab"), key.WithHelp("shift+tab", "shift+tab")),
}

type unfocusedKeyBindings struct {
	Edit key.Binding
	Save key.Binding
	Quit key.Binding
	Help key.Binding
}

func (k unfocusedKeyBindings) ShortHelp() []key.Binding {
	return []key.Binding{k.Edit, k.Save, k.Quit, k.Help}
}

func (k unfocusedKeyBindings) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Edit, k.Save},
		{k.Quit, k.Help},
	}
}

var unfocusedModeKeys = unfocusedKeyBindings{
	Edit: key.NewBinding(key.WithKeys("e"), key.WithHelp("e", "edit ->")),
	Save: key.NewBinding(key.WithKeys("s"), key.WithHelp("s", "save")),
	Quit: key.NewBinding(key.WithKeys("q"), key.WithHelp("q", "quit")),
	Help: key.NewBinding(key.WithKeys("h"), key.WithHelp("h", "help")),
}
