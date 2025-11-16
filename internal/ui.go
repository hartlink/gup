package ui

import (
	"fmt"
	"gup/pkg/i18n"
	"os/exec"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	titleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#04B575")).
			PaddingBottom(1)

	infoStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#666666"))

	successStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#04B575"))

	errorStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#FF0000"))

	loadingStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FFAA00"))
)

type CommandState int

const (
	StateInitial CommandState = iota
	StateRunning
	StateSuccess
	StateError
	StateDone
)

type CommandModel struct {
	cmdName     string
	cmdArgs     []string
	description string
	state       CommandState
	output      string
	error       error
	spinner     int
	quitting    bool
}

type tickMsg time.Time

func NewCommandModel(cmdName string, cmdArgs []string, description string) CommandModel {
	return CommandModel{
		cmdName:     cmdName,
		cmdArgs:     cmdArgs,
		description: description,
		state:       StateInitial,
	}
}

func (m CommandModel) Init() tea.Cmd {
	return tea.Batch(
		m.executeCommand(),
		m.tickCmd(),
	)
}

func (m CommandModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			m.quitting = true
			return m, tea.Quit
		case "enter":
			if m.state == StateSuccess || m.state == StateError {
				m.quitting = true
				return m, tea.Quit
			}
		}

	case tickMsg:
		if m.state == StateRunning {
			m.spinner = (m.spinner + 1) % 4
			return m, m.tickCmd()
		}

	case commandFinishedMsg:
		if msg.err != nil {
			m.state = StateError
			m.error = msg.err
		} else {
			m.state = StateSuccess
			m.output = msg.output
		}
		return m, nil
	}

	return m, nil
}

func (m CommandModel) View() string {
	if m.quitting {
		return ""
	}

	s := titleStyle.Render(fmt.Sprintf("🚀 %s", i18n.T("ui.title")))
	s += "\n\n"

	switch m.state {
	case StateInitial:
		s += infoStyle.Render(fmt.Sprintf("📋 %s: %s", i18n.T("ui.preparing"), m.description))

	case StateRunning:
		spinners := []string{"⩾", "⩽", "⩻", "⩿"}
		s += loadingStyle.Render(fmt.Sprintf("%s %s: %s", spinners[m.spinner], i18n.T("ui.executing"), m.description))
		cmdDisplay := m.cmdName
		if len(m.cmdArgs) > 0 {
			cmdDisplay += " " + strings.Join(m.cmdArgs, " ")
		}
		s += "\n" + infoStyle.Render(fmt.Sprintf("%s: %s", i18n.T("ui.command"), cmdDisplay))

	case StateSuccess:
		s += successStyle.Render(fmt.Sprintf("✅ %s", i18n.T("ui.success")))
		s += "\n\n" + infoStyle.Render(i18n.T("ui.output")+":")
		s += "\n" + m.formatOutput(m.output)
		s += "\n\n" + infoStyle.Render(i18n.T("ui.continue"))

	case StateError:
		s += errorStyle.Render(fmt.Sprintf("❌ %s", i18n.T("ui.error")))
		s += "\n\n" + errorStyle.Render(fmt.Sprintf("Error: %v", m.error))
		s += "\n\n" + infoStyle.Render(i18n.T("ui.continue"))
	}

	return s + "\n"
}

func (m CommandModel) tickCmd() tea.Cmd {
	return tea.Tick(time.Millisecond*100, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}

type commandFinishedMsg struct {
	output string
	err    error
}

func (m CommandModel) executeCommand() tea.Cmd {
	return func() tea.Msg {
		m.state = StateRunning

		if m.cmdName == "" {
			return commandFinishedMsg{err: fmt.Errorf("comando vacío")}
		}

		cmd := exec.Command(m.cmdName, m.cmdArgs...)
		// Preservar el entorno actual para sudo
		cmd.Env = append(cmd.Env, "SUDO_ASKPASS=")
		output, err := cmd.CombinedOutput()

		return commandFinishedMsg{
			output: string(output),
			err:    err,
		}
	}
}

func (m CommandModel) formatOutput(output string) string {
	lines := strings.Split(output, "\n")
	var formatted []string

	for i, line := range lines {
		if i > 20 { // Limitar a 20 líneas para evitar salida muy larga
			formatted = append(formatted, i18n.T("ui.truncated"))
			break
		}
		if strings.TrimSpace(line) != "" {
			formatted = append(formatted, "  "+line)
		}
	}

	if len(formatted) == 0 {
		return "  " + i18n.T("ui.no_output")
	}

	return strings.Join(formatted, "\n")
}
