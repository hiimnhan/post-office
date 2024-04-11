package ui

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/hiimnhan/post-office/theme"
	"github.com/hiimnhan/post-office/util"
)

var (
	borderStyle lipgloss.Style
	blankStyle  lipgloss.Style
	cursorStyle lipgloss.Style
)

var (
	sideBarTitleStyle    lipgloss.Style
	sideBarItemStyle     lipgloss.Style
	sideBarSelectedStyle lipgloss.Style
)

func (m model) InitStyles() {
	// SIDEBAR
	sideBarTitleStyle = lipgloss.NewStyle().Foreground(lipgloss.Color(m.theme.SideBarTitle))
	sideBarItemStyle = lipgloss.NewStyle().Background(lipgloss.Color(m.theme.BackgroundWindow))
	blankStyle = lipgloss.NewStyle().Background(lipgloss.Color(m.theme.BackgroundWindow))
	sideBarSelectedStyle = lipgloss.NewStyle()
	borderStyle = lipgloss.NewStyle().BorderStyle(lipgloss.ThickBorder()).BorderBackground(lipgloss.Color(m.theme.BackgroundWindow))
	cursorStyle = blankStyle.Foreground(lipgloss.Color(m.theme.Cursor))
}

func MethodLabel(method string, theme theme.Theme) string {
	var methodColorMap = map[string]string{
		"GET":     theme.MethodGet,
		"POST":    theme.MethodPost,
		"PUT":     theme.MethodPut,
		"PATCH":   theme.MethodPatch,
		"DELETE":  theme.MethodDelete,
		"OPTIONS": theme.MethodOptions,
	}

	color := methodColorMap[method]
	method = lipgloss.NewStyle().Foreground(lipgloss.Color(color)).Render(method)

	return method

}

// SIDEBAR

func (m model) EndpointsStyle(method string, url string, selected bool) string {
	method = MethodLabel(method, *m.theme)
	if selected {
		return method +
			sideBarSelectedStyle.Render(" "+util.TruncateText(url, sidebarWidth-4)) + "\n"
	}
	return method +
		sideBarItemStyle.Render(" "+util.TruncateText(url, sidebarWidth-4)) + "\n"

}

func (m model) SidebarStyle() lipgloss.Style {
	if m.panelFocus == sidebarPanelFocused {
		return borderStyle.
			BorderForeground(lipgloss.Color(m.theme.SidebarFocus)).
			Width(sidebarWidth).
			Height(m.mainPanelHeight).Bold(true).Background(lipgloss.Color(m.theme.BackgroundWindow))
	}
	return borderStyle.
		BorderForeground(lipgloss.Color(m.theme.SidebarFocus)).
		Width(sidebarWidth).
		Height(m.mainPanelHeight).Background(lipgloss.Color(m.theme.BackgroundWindow))
}
