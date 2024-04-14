package ui

func (m model) NewSidebarModel() string {
	s := blankStyle.Render(" ") + sideBarTitleStyle.Render("\ueb01 Endpoints")
	s += "\n\n"

	for i, item := range m.sidebarModel.items {
		cursor := " "
		if m.sidebarModel.cursor == i {
			cursor = ""
			if m.sidebarModel.selected == i {
				s += cursorStyle.Render(cursor) + cursorStyle.Render(" "+"*") + blankStyle.Render(" ") + m.EndpointsStyle(item.request.Request.Method, item.request.Request.URL.String(), true, sidebarWidth) + "\n"
			} else {
				s += cursorStyle.Render(cursor) + blankStyle.Render(" ") + m.EndpointsStyle(item.request.Request.Method, item.request.Request.URL.String(), true, sidebarWidth) + "\n"

			}
		} else {
			if m.sidebarModel.selected == i {
				s += blankStyle.Render(cursor) + cursorStyle.Render(" "+"*") + blankStyle.Render(" ") + m.EndpointsStyle(item.request.Request.Method, item.request.Request.URL.String(), false, sidebarWidth) + "\n"
			} else {
				s += blankStyle.Render(cursor) + blankStyle.Render(" ") + m.EndpointsStyle(item.request.Request.Method, item.request.Request.URL.String(), false, sidebarWidth) + "\n"
			}

		}

		s += "\n"
	}

	s = m.SidebarStyle().Render(s)

	return s
}
