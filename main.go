package main

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/nanorex07/notas/models"
)

func main() {

	rootModel := models.NewRootModel()
	p := tea.NewProgram(rootModel, tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
