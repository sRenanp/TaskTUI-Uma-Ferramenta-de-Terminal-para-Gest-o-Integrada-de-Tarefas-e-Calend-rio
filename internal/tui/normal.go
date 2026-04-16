package tui

import (
	"tui-task-manager/internal/core"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) updateNormal(msg tea.Msg) (tea.Model, tea.Cmd) {
	keyMsg, ok := msg.(tea.KeyMsg)
	if !ok {
		return m, nil
	}

	switch keyMsg.String() {
	// Navegação entre seções
	case "tab", "ctrl+l":
		m.activeSec = (m.activeSec + 1) % 3
		m.cursor = 0
	case "shift+tab", "ctrl+h":
		m.activeSec = (m.activeSec - 1 + 3) % 3
		m.cursor = 0

	// Navegação vertical
	case "j", "down":
		m.moveCursorDown()
	case "k", "up":
		m.moveCursorUp()

	// Navegação horizontal (detalhes)
	case "l", "right":
		if m.activeSec == DetailSection {
			// m.moveCursorRight()
		}
	case "h", "left":
		if m.activeSec == DetailSection {
			// m.moveCursorLeft()
		}

	// Ações
	case "enter":
		switch m.activeSec {
		case TaskListSection:
			m.selectedTaskIdx = m.cursor
			m.activeSec = DetailSection
			m.cursor = 0
		case DetailSection:
			switch m.cursor {
			case 0: // Descrição
				m.mode = InsertMode
				m.textInput.Focus()
				m.textInput.SetValue(m.manager.GetTasks()[m.selectedTaskIdx].Description)
				return m, textinput.Blink
			case 1: // Status
				m.mode = InsertMode
				return m, nil
			case 2: // Prioridade
				m.mode = InsertMode
				return m, nil
			}
		}

	case "d":
		if m.activeSec == TaskListSection {
			m.deleteCurrentTask()
			m.message = "Tarefa deletada"
			return m, clearMessageAfter()
		}

	case "a":
		m.mode = CommandMode
		m.command = ":add "

	case ":":
		m.mode = CommandMode
		m.command = ":"

	case "q":
		if m.showHelp {
			m.showHelp = false
		} else {
			return m, tea.Quit
		}
	}

	return m, nil
}

// cycleStatus alterna o status da tarefa (pode ser usado futuramente)
func (m *Model) cycleStatus(task *core.Task) {
	newStatus := (task.Status + 1) % 3
	updateCmd := &core.UpdateTaskCommand{
		Manager: m.manager,
		TaskID:  task.ID,
		Updates: map[string]interface{}{"status": core.Status(newStatus)},
	}
	m.manager.ExecuteCommand(updateCmd)
}

func (m *Model) deleteCurrentTask() {
	tasks := m.manager.GetTasks()

	if len(tasks) == 0 || m.cursor >= len(tasks) {
		return
	}

	taskID := tasks[m.cursor].ID

	delCmd := &core.DeleteTaskCommand{
		Manager: m.manager,
		TaskID:  taskID,
	}

	if err := m.manager.ExecuteCommand(delCmd); err == nil {

		if m.cursor > 0 && m.cursor == len(tasks)-1 {
			m.cursor--
		}
	} else {
		m.message = "Erro ao deletar: " + err.Error()
	}
}

func (m *Model) moveCursorDown() {
	switch m.activeSec {
	case TaskListSection:
		tasks := m.manager.GetTasks()
		if len(tasks) > 0 && m.cursor < len(tasks)-1 {
			m.cursor++
		}
	case DetailSection:
		if m.cursor < 2 {
			m.cursor++
		}
	}
}

func (m *Model) moveCursorUp() {
	switch m.activeSec {
	case TaskListSection:
		if m.cursor > 0 {
			m.cursor--
		}
	case DetailSection:
		if m.cursor > 0 {
			m.cursor--
		}
	}
}
