package edit

import "github.com/nanorex07/notas/types"

func (m EditModel) View() string {

	switch m.appSettings.EditViewMode {
	case types.CreateMode:
		m.heading.SetText("Create new note")
	case types.EditMode:
		m.heading.SetText("Edit note")
	}

	return m.heading.Render()
}
