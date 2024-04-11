package ui

import ()

func (m model) NewSidebarModel() string {
	s := "\n"
	s = blankStyle.Render(" ") + sideBarTitleStyle.Render("\ueb01 Endpoints")
	s += "\n\n"

	for i, item := range m.sidebarModel.items {
		cursor := " "
		if m.sidebarModel.selected == i {
			cursor = ""
			s += cursorStyle.Render(cursor) + blankStyle.Render(" ") + m.EndpointsStyle(item.request.Request.Method, item.request.Request.URL.String(), true)
		} else {
			s += blankStyle.Render(cursor) + blankStyle.Render(" ") + m.EndpointsStyle(item.request.Request.Method, item.request.Request.URL.String(), false)

		}
		s += "\n"
	}

	s = m.SidebarStyle().Render(s)

	return s
}
