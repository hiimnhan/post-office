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
var (
	tabStyle          lipgloss.Style
	tabSelectedStyle  lipgloss.Style
	tabSelectedBorder = lipgloss.Border{
		Top:         "─",
		Bottom:      " ",
		Left:        "│",
		Right:       "│",
		TopLeft:     "╭",
		TopRight:    "╮",
		BottomLeft:  "┘",
		BottomRight: "└",
	}
	tabBorder = lipgloss.Border{
		Top:         "─",
		Bottom:      "─",
		Left:        "│",
		Right:       "│",
		TopLeft:     "╭",
		TopRight:    "╮",
		BottomLeft:  "┴",
		BottomRight: "┴",
	}
)

func (m model) InitStyles() {
	// SIDEBAR
	sideBarTitleStyle = lipgloss.NewStyle().Foreground(lipgloss.Color(m.theme.SideBarTitle))
	sideBarItemStyle = lipgloss.NewStyle().Background(lipgloss.Color(m.theme.BackgroundWindow))

	blankStyle = lipgloss.NewStyle().Background(lipgloss.Color(m.theme.BackgroundWindow))

	sideBarSelectedStyle = lipgloss.NewStyle().Inherit(blankStyle)

	borderStyle = lipgloss.NewStyle().BorderStyle(lipgloss.ThickBorder()).BorderBackground(lipgloss.Color(m.theme.BackgroundWindow))

	cursorStyle = blankStyle.Foreground(lipgloss.Color(m.theme.Cursor))

	tabStyle = lipgloss.NewStyle().Border(tabBorder, true).BorderForeground(lipgloss.Color(m.theme.Border)).Width(20).BorderBackground(lipgloss.Color(m.theme.BackgroundWindow)).Inherit(blankStyle)
	tabSelectedStyle = tabStyle.Copy().Border(tabSelectedBorder, true)
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
	method = lipgloss.NewStyle().Foreground(lipgloss.Color(color)).Background(lipgloss.Color(theme.BackgroundWindow)).Render(method)

	return method

}

// SIDEBAR

func (m model) EndpointsStyle(method string, url string, selected bool, width int) string {
	method = MethodLabel(method, *m.theme)
	if selected {
		return method +
			sideBarSelectedStyle.Render(" "+util.TruncateText(url, width-4))
	}
	return method +
		sideBarItemStyle.Render(" "+util.TruncateText(url, width-4))

}

func (m model) SidebarStyle() lipgloss.Style {
	if m.panelFocus == sidebarPanelFocused {
		return borderStyle.
			BorderForeground(lipgloss.Color(m.theme.SidebarFocus)).
			Width(sidebarWidth).
			Height(m.mainPanelHeight).Bold(true).Background(lipgloss.Color(m.theme.BackgroundWindow)).
			PaddingTop(1).
			MarginTop(5)
	}
	return borderStyle.
		BorderForeground(lipgloss.Color(m.theme.Border)).
		Width(sidebarWidth).
		Height(m.mainPanelHeight).Background(lipgloss.Color(m.theme.BackgroundWindow))
}

func (m model) MainPanelStyle() lipgloss.Style {
	if m.panelFocus == mainPanelFocused {
		return borderStyle.
			BorderForeground(lipgloss.Color(m.theme.MainPanelFocus)).
			Width(m.fullWidth - sidebarWidth).
			Height(m.mainPanelHeight).Bold(true).Background(lipgloss.Color(m.theme.BackgroundWindow))
	}
	return borderStyle.
		BorderForeground(lipgloss.Color(m.theme.Border)).
		Width(m.fullWidth - sidebarWidth).
		Height(m.mainPanelHeight).Background(lipgloss.Color(m.theme.BackgroundWindow))
}

func (m model) InputStyle() lipgloss.Style {
	return borderStyle.
		BorderForeground(lipgloss.Color(m.theme.Border)).
		Padding(2).
		Height(minimumHeight - 5)

}

func (m model) RequestTabStyle(selected bool) lipgloss.Style {
	if selected {
		return tabSelectedStyle
	}
	return tabStyle
}
