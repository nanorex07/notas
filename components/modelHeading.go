package components

import "github.com/charmbracelet/lipgloss"

var (
	modelHeading = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("80"))
)

type ModelHeading struct {
	text string
}

func NewModelHeading(text string) *ModelHeading {
	return &ModelHeading{text: text}
}

func (m *ModelHeading) SetText(text string) {
	m.text = text
}

func (m *ModelHeading) Render() string {
	return modelHeading.Render(m.text)
}
