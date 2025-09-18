package models

import "github.com/charmbracelet/bubbles/key"

var (
	helpKey = key.NewBinding(key.WithKeys("h"))
	quitKey = key.NewBinding(key.WithKeys("q", "ctrl+c"))
)
