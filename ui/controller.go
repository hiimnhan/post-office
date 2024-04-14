package ui

// sidebar
func ControllerSidebarUp(m model) model {
	if m.sidebarModel.cursor > 0 {
		m.sidebarModel.cursor--
	}
	return m
}

func ControllerSidebarDown(m model) model {
	if m.sidebarModel.cursor < len(m.sidebarModel.items)-1 {
		m.sidebarModel.cursor++
	}
	return m
}

func SwitchPanelFocus(m model) model {
	m.panelFocus += 1
	if m.panelFocus > mainPanelFocused {
		m.panelFocus = sidebarPanelFocused
	}
	return m
}

func OpenNewRequesTab(m model) model {
	m.sidebarModel.selected = m.sidebarModel.cursor
	m.sidebarModel.selectedItem = &m.sidebarModel.items[m.sidebarModel.selected]
	if m.mainPanelModel.openTabUuids[m.sidebarModel.selectedItem.uuid] {
		m.mainPanelModel.selectedTab = m.sidebarModel.selectedItem.uuid
	} else {
		m.mainPanelModel.requestTabs = append(m.mainPanelModel.requestTabs, requestTabModel{
			sidebarItem: m.sidebarModel.selectedItem,
		})
		m.mainPanelModel.selectedTab = m.sidebarModel.selectedItem.uuid
		m.mainPanelModel.openTabUuids[m.sidebarModel.selectedItem.uuid] = true
	}

	return m

}
