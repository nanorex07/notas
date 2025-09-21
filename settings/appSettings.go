package models

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/nanorex07/notas/types"
)

type AppSettings struct {
	ActiveView   types.ActiveView
	EditViewMode types.EditViewMode
	HelpModel    *help.Model
}
