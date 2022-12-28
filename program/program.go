package program

import (
	tea "github.com/charmbracelet/bubbletea"
	"log"
	"strings"
)

type appState struct {
	authUsername   string
	authPassword   string
	targetUsername string
	targetPostURL  string
}

type appModel struct {
	stage  Stage
	models map[Stage]StageModel
	state  appState
}

func New() *tea.Program {
	m := appModel{
		stage: StageInput,
	}

	models := map[Stage]StageModel{
		StageInput: newInputModel(&m.state),
		StageLoad:  newLoadModel(&m.state),
	}
	m.models = models

	p := tea.NewProgram(m)
	return p
}

func (m appModel) Init() tea.Cmd {
	return m.models[StageInput].Init()
}

func (m appModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		}
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
	return m, nextModel.Init()
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
