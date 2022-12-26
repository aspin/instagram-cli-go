package program

import tea "github.com/charmbracelet/bubbletea"

type Stage int

type StageModel interface {
	Init() tea.Cmd

	Update(tea.Msg) (Stage, StageModel, tea.Cmd)

	View() string
}

var (
	StageExit  Stage = 0
	StageInput Stage = 1
	StageLoad  Stage = 2
)
