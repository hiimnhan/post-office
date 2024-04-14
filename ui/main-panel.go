package ui

import "github.com/charmbracelet/lipgloss"

func (m model) NewMainPanelModel() string {
	var tabs []string
	for _, tab := range m.mainPanelModel.requestTabs {
		tabs = append(tabs, m.NewRequestTabModel(tab, tab.sidebarItem.uuid))

	}
	t := lipgloss.JoinHorizontal(lipgloss.Top, tabs...)
	wrapper := m.MainPanelStyle().Render(t)
	return wrapper
}
