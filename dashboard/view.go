package dashboard

func (m DashboardModel) View() string {
	return m.list.Render()
}
