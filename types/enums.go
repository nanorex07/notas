package types

type ActiveView int

const (
	DashboardView ActiveView = iota
	EditView
	DisplayView
)

type EditViewMode int

const (
	CreateMode EditViewMode = iota
	EditMode
)
