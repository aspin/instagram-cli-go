package program

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"log"
	"strings"
)

var (
	focusedStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("86"))
	noStyle      = lipgloss.NewStyle()
)

type inputModel struct {
	inputs     []textinput.Model
	focusIndex int

	appState *appState
}

func newInputModel(appState *appState) StageModel {
	m := inputModel{
		inputs:   make([]textinput.Model, 4),
		appState: appState,
	}

	for i := range m.inputs {
		t := textinput.New()
		switch i {
		case 0:
			t.Placeholder = "Instagram Username (for authentication)"
			t.Focus()
			t.PromptStyle = focusedStyle
		case 1:
			t.Placeholder = "Instagram Password (for authentication)"
			t.EchoMode = textinput.EchoPassword
			t.EchoCharacter = '*'
		case 2:
			t.Placeholder = "Instagram Username (for follower lookup)"
		case 3:
			t.Placeholder = "Instagram Post Link"
			// TODO: validate text is link
		}

		m.inputs[i] = t
	}
	return m
}

func (m inputModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m inputModel) Update(msg tea.Msg) (Stage, StageModel, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch k := msg.Type; k {
		case tea.KeyTab, tea.KeyShiftTab, tea.KeyEnter, tea.KeyUp, tea.KeyDown:
			// continue on in app: for now, quit
			if k == tea.KeyEnter && m.focusIndex == len(m.inputs) {
				return StageLoad, m, nil
			}

			if k == tea.KeyUp || k == tea.KeyShiftTab {
				m.focusIndex--
			} else {
				log.Printf("[input]: focus down")
				m.focusIndex++
			}

			if m.focusIndex > len(m.inputs) {
				m.focusIndex = 0
			} else if m.focusIndex < 0 {
				m.focusIndex = len(m.inputs)
			}

			cmds := make([]tea.Cmd, len(m.inputs))
			for i := range m.inputs {
				if i == m.focusIndex {
					cmds[i] = m.inputs[i].Focus()
					m.inputs[i].PromptStyle = focusedStyle
					continue
				}

				m.inputs[i].Blur()
				m.inputs[i].PromptStyle = noStyle
			}
			return StageInput, m, tea.Batch(cmds...)
		}
	}

	cmd = m.updateInputs(msg)
	return StageInput, m, cmd
}

func (m *inputModel) updateInputs(msg tea.Msg) tea.Cmd {
	cmds := make([]tea.Cmd, len(m.inputs))

	for i := range m.inputs {
		m.inputs[i], cmds[i] = m.inputs[i].Update(msg)
	}

	return tea.Batch(cmds...)
}

func (m inputModel) View() string {
	var b strings.Builder

	for i := range m.inputs {
		b.WriteString(m.inputs[i].View())
		b.WriteRune('\n')
	}

	button := "\n[ Submit ]\n"
	if m.focusIndex == len(m.inputs) {
		button = focusedStyle.Render(button)
	}
	b.WriteString(button)
	return b.String()
}
