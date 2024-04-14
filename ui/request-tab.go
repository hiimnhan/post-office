package ui

import "github.com/google/uuid"

func (m model) NewRequestTabModel(rt requestTabModel, uuid uuid.UUID) string {
	title := m.EndpointsStyle(
		rt.sidebarItem.request.Request.Method,
		rt.sidebarItem.request.Request.URL.String(),
		false, 14)

	selected := m.mainPanelModel.selectedTab == uuid

	return m.RequestTabStyle(selected).Render(title)

}
