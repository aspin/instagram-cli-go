package program

import (
	tea "github.com/charmbracelet/bubbletea"
	"instagram-cli-go/instagram"
	"log"
	"strings"
)

type appState struct {
	err error

	windowWidth  int
	windowHeight int

	authUsername   string
	authPassword   string
	targetUsername string
	targetPostURL  string

	followers    instagram.UserSet
	post         instagram.Media
	postLikers   instagram.UserSet
	postComments []instagram.Comment
}

type appModel struct {
	stage    Stage
	models   map[Stage]StageModel
	state    *appState
	dispatch StageDispatcher
}

func New() *tea.Program {
	m := &appModel{
		stage: StageInput,
		state: &appState{},
	}

	models := map[Stage]StageModel{
		StageInput:   newInputModel(m.state),
		StageLoad:    newLoadModel(m.state),
		StageProcess: newProcessModel(m.state),
		StageError:   newErrorModel(m.state),
	}
	m.models = models

	p := tea.NewProgram(m, tea.WithAltScreen())
	m.dispatch = p.Send
	return p
}

func (m appModel) Init() tea.Cmd {
	return m.models[StageInput].Init(m.dispatch)
}

func (m appModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		}
	case tea.WindowSizeMsg:
		m.state.windowWidth = msg.Width
		m.state.windowHeight = msg.Height
	}

	// process current stage updates
	model, ok := m.models[m.stage]
	if !ok {
		log.Printf("error[update]: could not find model for stage %v", m.stage)
		return m, tea.Quit
	}

	nextStage, nextModel, nextCmd := model.Update(msg)
	if nextStage == StageExit {
		return m, tea.Quit
	}
	m.models[m.stage] = nextModel

	// no stage transition; continue
	if m.stage == nextStage {
		return m, nextCmd
	}

	// stage transition: move onto initializing next model
	nextModel, ok = m.models[nextStage]
	if !ok {
		log.Printf("error[update]: could not find model for next stage %v", m.stage)
		return m, tea.Quit
	}
	m.stage = nextStage
	return m, nextModel.Init(m.dispatch)
}

func (m appModel) View() string {
	var b strings.Builder
	b.WriteString("Instagram Giveaway CLI Application\n\n")

	model, ok := m.models[m.stage]
	if !ok {
		log.Printf("error[view]: could not find model for stage %v", m.stage)
		return ""
	}

	b.WriteString(model.View())
	return b.String()
}
