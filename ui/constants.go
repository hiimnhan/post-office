package ui

type focusedPanel int

// panel focus
const (
	nonePanelFocused focusedPanel = iota
	sidebarPanelFocused
	mainPanelFocused
)

// dimensions
const (
	sidebarWidth  = 40
	minimumWidth  = 96
	minimumHeight = 35
)
