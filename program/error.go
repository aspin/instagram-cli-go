package program

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"strings"
)

var (
	errorStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("197"))
	helpStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("#626262"))
)

type errorModel struct {
	appState *appState
}

func newErrorModel(appState *appState) StageModel {
	return errorModel{appState: appState}
}

func (e errorModel) Init(dispatch StageDispatcher) tea.Cmd {
	return nil
}

func (e errorModel) Update(msg tea.Msg) (Stage, StageModel, tea.Cmd) {
	return StageError, e, nil
}

func (e errorModel) View() string {
	var b strings.Builder

	b.WriteString(errorStyle.Render("Encountered fatal error:"))
	b.WriteRune('\n')
	b.WriteRune('\n')

	b.WriteString(noStyle.Render(e.appState.err.Error()))
	b.WriteRune('\n')
	b.WriteRune('\n')

	b.WriteString(helpStyle.Render("(ctrl+c or esc to exit)"))
	return b.String()
}
