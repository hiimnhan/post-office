package ui

// sidebar
func ControllerSidebarUp(m model) model {
	if m.sidebarModel.selected > 0 {
		m.sidebarModel.selected--
	}
	return m
}

func ControllerSidebarDown(m model) model {
	if m.sidebarModel.selected < len(m.sidebarModel.items)-1 {
		m.sidebarModel.selected++
	}
	return m
}
