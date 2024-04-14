package ui

import (
	"time"

	"github.com/google/uuid"
	"github.com/hiimnhan/post-office/theme"
	"github.com/hiimnhan/post-office/types"
)

type channelMsg struct{}

type sidebarItem struct {
	request   types.Request
	createdAt time.Time
	uuid      uuid.UUID
}
type sidebarModel struct {
	items        []sidebarItem
	cursor       int
	selected     int
	selectedItem *sidebarItem
}

type inputModel struct {
	input string
}

type requestTabModel struct {
	sidebarItem *sidebarItem
}

type mainPanelModel struct {
	requestTabs  []requestTabModel
	selectedTab  uuid.UUID
	openTabUuids map[uuid.UUID]bool
}

type model struct {
	sidebarModel    sidebarModel
	mainPanelModel  mainPanelModel
	mainPanelHeight int
	theme           *theme.Theme
	panelFocus      focusedPanel
	fullWidth       int
	fullHeight      int
}
