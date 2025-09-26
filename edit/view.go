package edit

import (
	"strings"

	"github.com/nanorex07/notas/types"
)

func (m EditModel) View() string {

	switch m.appSettings.EditViewMode {
	case types.CreateMode:
		m.heading.SetText("Create new note")
	case types.EditMode:
		m.heading.SetText("Edit note")
	}

	viewBuilder := strings.Builder{}
	for i := range m.inputs {
		viewBuilder.WriteString(m.inputs[i].View())
		viewBuilder.WriteString("\n\n")
	}

	return viewBuilder.String()
}
