package ui

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/log"
	"github.com/google/uuid"
	"github.com/hiimnhan/post-office/theme"
	"github.com/hiimnhan/post-office/types"
	"github.com/hiimnhan/post-office/util"
)

type channelMsg struct{}

var channel chan channelMsg = make(chan channelMsg)

type sidebarItem struct {
	request types.Request
}
type sidebarModel struct {
	items    []sidebarItem
	selected int
}

type model struct {
	sidebarModel    sidebarModel
	mainPanelHeight int
	theme           *theme.Theme
	panelFocus      focusedPanel
	fullWidth       int
	fullHeight      int
}

func listenToChannelMsg(msg chan channelMsg) tea.Cmd {
	return func() tea.Msg {
		select {
		case m := <-msg:
			return m
		default:
			return nil
		}
	}
}

func (m model) Init() tea.Cmd {
	return tea.Batch(
		tea.SetWindowTitle("Post Office"),
		listenToChannelMsg(channel),
	)
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.fullWidth = msg.Width
		m.fullHeight = msg.Height
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, util.Keys.Quit):
			return m, tea.Quit
		case key.Matches(msg, util.Keys.Up):
			if m.panelFocus == sidebarPanelFocused {
				m = ControllerSidebarUp(m)
			}
		case key.Matches(msg, util.Keys.Down):
			if m.panelFocus == sidebarPanelFocused {
				m = ControllerSidebarDown(m)
			}
		}
	}

	return m, cmd
}

func (m model) View() string {
	sidebar := m.NewSidebarModel()

	return sidebar
}

func NewModel(cfgFile string) model {
	theme, err := theme.LoadThemeConfig(cfgFile)
	log.Debug("Model theme", theme)
	if err != nil {
		log.Fatal("Cannot load theme config", err)
	}
	// test
	methods := []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"}
	sidebarItems := []sidebarItem{}
	var item sidebarItem

	for i := 0; i < 10; i++ {
		httpRequest, _ := http.NewRequest(methods[rand.Intn(len(methods))], "http://localhost:8080", nil)
		item = sidebarItem{
			request: types.Request{
				Request:   httpRequest,
				UUID:      uuid.New().String(),
				CreatedAt: time.Now(),
			},
		}

		sidebarItems = append(sidebarItems, item)
	}
	return model{
		theme: theme,
		sidebarModel: sidebarModel{
			items:    sidebarItems,
			selected: 0,
		},
		panelFocus: sidebarPanelFocused,
	}
}
