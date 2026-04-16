package tui

import (
	"tui-task-manager/internal/core"

	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) updateInsert(msg tea.Msg) (tea.Model, tea.Cmd) {
	tasks := m.manager.GetTasks()
	if len(tasks) == 0 || m.selectedTaskIdx >= len(tasks) {
		m.mode = NormalMode
		return m, nil
	}

	task := tasks[m.selectedTaskIdx]

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			m.textInput.Blur()
			m.mode = NormalMode
			return m, nil

		case "enter":
			if m.cursor == 0 {
				newDesc := m.textInput.Value()
				updateCmd := &core.UpdateTaskCommand{
					Manager: m.manager,
					TaskID:  task.ID,
					Updates: map[string]interface{}{"description": newDesc},
				}
				m.manager.ExecuteCommand(updateCmd)
			}

			m.textInput.Blur()
			m.mode = NormalMode
			return m, nil

		// Seleção rápida (status / prioridade)
		case "1", "2", "3":
			key := msg.String()

			if m.cursor == 1 {
				var newStatus core.Status
				switch key {
				case "1":
					newStatus = core.Todo
				case "2":
					newStatus = core.InProgress
				case "3":
					newStatus = core.Done
				}

				// Debug
				// m.message = fmt.Sprintf("Tentando mudar status para: %d", newStatus)

				updateCmd := &core.UpdateTaskCommand{
					Manager: m.manager,
					TaskID:  task.ID,
					Updates: map[string]interface{}{"status": newStatus},
				}

				if err := m.manager.ExecuteCommand(updateCmd); err == nil {
					// m.message = "Status atualizado"
				}

				m.mode = NormalMode
				m.textInput.Blur()
				return m, nil
			}

			if m.cursor == 2 {
				var newPriority core.Priority
				switch key {
				case "1":
					newPriority = core.Low
				case "2":
					newPriority = core.Medium
				case "3":
					newPriority = core.High
				}

				updateCmd := &core.UpdateTaskCommand{
					Manager: m.manager,
					TaskID:  task.ID,
					Updates: map[string]interface{}{"priority": newPriority},
				}

				m.manager.ExecuteCommand(updateCmd)
				m.mode = NormalMode
				m.textInput.Blur()
				return m, nil
			}
		}
	}

	// Input de texto apenas para descrição
	if m.cursor == 0 {
		var cmd tea.Cmd
		m.textInput, cmd = m.textInput.Update(msg)
		return m, cmd
	}

	return m, nil
}
