package program

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var listStyle = lipgloss.NewStyle().Margin(1, 2)

type processModel struct {
	appState *appState
	list     list.Model
}

type actionItem struct {
	title string
	desc  string
}

func (i actionItem) Title() string {
	return i.title
}

func (i actionItem) Description() string {
	return i.desc
}

func (i actionItem) FilterValue() string {
	return i.title
}

func newProcessModel(appState *appState) StageModel {
	items := []list.Item{
		actionItem{
			title: "Summary",
			desc:  "Short summary of loaded data",
		},
		actionItem{
			title: "Review comments",
			desc:  "Ensure submitted comments answer provided question",
		},
		actionItem{
			title: "Review entries",
			desc:  "View each entered user and the number of votes they have",
		},
		actionItem{
			title: "Choose winner",
			desc:  "Randomly select a winner from the entries",
		},
	}

	m := &processModel{
		appState: appState,
		list:     list.New(items, list.NewDefaultDelegate(), 0, 0),
	}
	m.list.Title = "Action Items"
	return m
}

func (m *processModel) Init(dispatch StageDispatcher) tea.Cmd {
	m.setSize()
	return nil
}

func (m *processModel) setSize() {
	h, v := listStyle.GetFrameSize()
	m.list.SetSize(m.appState.windowWidth-h, m.appState.windowHeight-v)
}

func (m *processModel) Update(msg tea.Msg) (Stage, StageModel, tea.Cmd) {
	switch msg.(type) {
	case tea.WindowSizeMsg:
		m.setSize()
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return StageProcess, m, cmd
}

func (m *processModel) View() string {
	return listStyle.Render(m.list.View())
}
