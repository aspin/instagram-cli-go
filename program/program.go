package program

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"strings"
)

type errMsg error

type model struct {
	inputs     []textinput.Model
	focusIndex int
}

func New() *tea.Program {
	m := model{
		inputs: make([]textinput.Model, 4),
	}

	for i := range m.inputs {
		t := textinput.New()
		switch i {
		case 0:
			t.Placeholder = "Instagram Username (for authentication)"
			t.Focus()
		case 1:
			t.Placeholder = "Instagram Password (for authentication)"
			t.EchoMode = textinput.EchoPassword
			t.EchoCharacter = '*'
		case 2:
			t.Placeholder = "Instagram Username (for follower-lookup)"
		case 3:
			t.Placeholder = "Instagram Post Link"
			// TODO: validate = link
		}

		m.inputs[i] = t
	}

	p := tea.NewProgram(m)
	return p
}

func (m model) Init() tea.Cmd {
	return textinput.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		case tea.KeyTab, tea.KeyShiftTab, tea.KeyEnd, tea.KeyUp, tea.KeyDown:

		}
	}

	cmd = m.updateInputs(msg)
	return m, cmd
}

func (m *model) updateInputs(msg tea.Msg) tea.Cmd {
	cmds := make([]tea.Cmd, len(m.inputs))

	for i := range m.inputs {
		m.inputs[i], cmds[i] = m.inputs[i].Update(msg)
	}

	return tea.Batch(cmds...)
}

func (m model) View() string {
	var b strings.Builder

	for i := range m.inputs {
		b.WriteString(m.inputs[i].View())
		b.WriteRune('\n')
	}

	b.WriteString("\n[ Submit ]\n")
	return b.String()
}
