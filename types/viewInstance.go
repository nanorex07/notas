package types

import (
	"github.com/charmbracelet/bubbles/help"
	tea "github.com/charmbracelet/bubbletea"
)

type ViewInstance interface {
	tea.Model
	KeyBindings() help.KeyMap
}
