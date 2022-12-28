package program

import (
	"fmt"
	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
	"strings"
	"time"
)

const (
	maxWidth = 80
)

type loadModel struct {
	progress progress.Model
	appState *appState
}

type tickMsg time.Time

func newLoadModel(appState *appState) StageModel {
	m := loadModel{
		progress: progress.New(progress.WithDefaultGradient()),
		appState: appState,
	}
	return m
}

func (m loadModel) Init() tea.Cmd {
	return tickCmd()
}

func (m loadModel) Update(msg tea.Msg) (Stage, StageModel, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.progress.Width = msg.Width - 4
		if m.progress.Width > maxWidth {
			m.progress.Width = maxWidth
		}
		return StageLoad, m, nil

	case tickMsg:
		if m.progress.Percent() == 1.0 {
			return StageExit, m, tea.Quit
		}

		cmd := m.progress.IncrPercent(0.25)
		return StageLoad, m, tea.Batch(tickCmd(), cmd)

	case progress.FrameMsg:
		progressModel, cmd := m.progress.Update(msg)
		m.progress = progressModel.(progress.Model)
		return StageLoad, m, cmd

	default:
		return StageLoad, m, nil
	}
}

func (m loadModel) View() string {
	var b strings.Builder
	_, _ = fmt.Fprintf(&b, "Loading post (%v) and user %v...\n", m.appState.targetPostURL, m.appState.targetUsername)
	b.WriteString(m.progress.View())
	b.WriteString("\n\n")

	return b.String()
}

func tickCmd() tea.Cmd {
	return tea.Tick(time.Second*1, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}
