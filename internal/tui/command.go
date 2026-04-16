package tui

import (
	"strings"
	"tui-task-manager/internal/core"

	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) updateCommand(msg tea.Msg) (tea.Model, tea.Cmd) {
	keyMsg, ok := msg.(tea.KeyMsg)
	if !ok {
		return m, nil
	}

	switch keyMsg.String() {
	case "enter":
		return m.executeCommand()
	case "esc":
		m.mode = NormalMode
		m.command = ""
	case "backspace":
		if len(m.command) > 1 {
			m.command = m.command[:len(m.command)-1]
		} else {
			m.mode = NormalMode
			m.command = ""
		}
	default:
		if len(keyMsg.String()) == 1 {
			m.command += keyMsg.String()
		}
	}
	return m, nil
}

func (m Model) executeCommand() (tea.Model, tea.Cmd) {
	cmd := strings.TrimSpace(m.command)

	switch {
	case cmd == ":q":
		return m, tea.Quit
	case cmd == ":w":
		// m.manager.Save()
		m.message = "Salvo com sucesso"
		m.mode = NormalMode
		return m, clearMessageAfter()

	case strings.HasPrefix(cmd, ":add "):
		desc := strings.TrimPrefix(cmd, ":add ")
		if desc != "" {
			task := &core.Task{
				Description: desc,
			}
			addCmd := &core.AddTaskCommand{
				Manager: m.manager,
				Task:    task,
			}

			if err := m.manager.ExecuteCommand(addCmd); err == nil {
				m.message = "Tarefa adicionada"
			} else {
				m.message = "Erro: " + err.Error()
			}
		}
		m.mode = NormalMode
		return m, clearMessageAfter()

	case strings.HasPrefix(cmd, ":delete"):
		m.deleteCurrentTask()
		m.mode = NormalMode
		return m, clearMessageAfter()
	}

	m.mode = NormalMode
	m.command = ""
	return m, nil
}
