package program

import (
	"fmt"
	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
	"instagram-cli-go/instagram"
	"strings"
	"time"
)

const (
	maxWidth = 80
)

type loadModel struct {
	progress         progress.Model
	appState         *appState
	instagramService instagram.Service
}

type progressMsg struct {
	completed float64
}

type failureMsg struct {
	err error
}

func newLoadModel(appState *appState) StageModel {
	m := loadModel{
		progress:         progress.New(progress.WithDefaultGradient()),
		appState:         appState,
		instagramService: instagram.NewService(appState.authUsername, appState.authPassword),
	}
	return m
}

func (m loadModel) Init(dispatch StageDispatcher) tea.Cmd {
	go func() {
		time.Sleep(50 * time.Millisecond)
		var err error
		m.appState.followers, err = m.instagramService.Followers(m.appState.targetUsername)
		if err != nil {
			dispatch(failureMsg{err: fmt.Errorf("could not load followers: %w", err)})
			return
		}
		dispatch(progressMsg{completed: 0.1})
	}()
	go func() {
		time.Sleep(100 * time.Millisecond)
		var err error
		m.appState.post, err = m.instagramService.Post(m.appState.targetPostURL)
		if err != nil {
			dispatch(failureMsg{err: fmt.Errorf("could not load post: %w", err)})
			return
		}
		dispatch(progressMsg{completed: 0.1})
	}()
	go func() {
		time.Sleep(1 * time.Second)
		var err error
		m.appState.postLikers, err = m.instagramService.PostLikers(m.appState.targetPostURL)
		if err != nil {
			dispatch(failureMsg{err: fmt.Errorf("could not load post likers: %w", err)})
			return
		}
		dispatch(progressMsg{completed: 0.6})

		go func() {
			time.Sleep(200 * time.Millisecond)
			dispatch(progressMsg{completed: 1.0})
		}()
	}()
	go func() {
		time.Sleep(400 * time.Millisecond)
		var err error
		m.appState.postComments, err = m.instagramService.PostComments(m.appState.targetPostURL)
		if err != nil {
			dispatch(failureMsg{err: fmt.Errorf("could not load comments: %w", err)})
			return
		}
		dispatch(progressMsg{completed: 0.2})
	}()
	return nil
}

func (m loadModel) Update(msg tea.Msg) (Stage, StageModel, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.progress.Width = msg.Width - 4
		if m.progress.Width > maxWidth {
			m.progress.Width = maxWidth
		}
		return StageLoad, m, nil

	case progressMsg:
		if m.progress.Percent() == 1.0 {
			return StageProcess, nil, nil
		}

		cmd := m.progress.IncrPercent(msg.completed)
		return StageLoad, m, cmd

	case failureMsg:
		m.appState.err = msg.err
		return StageError, nil, nil

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
