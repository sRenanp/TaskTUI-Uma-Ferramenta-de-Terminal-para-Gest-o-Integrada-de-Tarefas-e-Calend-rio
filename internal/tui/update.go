package tui

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type clearMessageMsg struct{}

func clearMessageAfter() tea.Cmd {
	return tea.Tick(2*time.Second, func(t time.Time) tea.Msg {
		return clearMessageMsg{}
	})
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// Mensagens globais
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		m.ready = true
		return m, nil

	case clearMessageMsg:
		m.message = ""
		m.messageTimer = 0
		return m, nil

	case tea.KeyMsg:
		// Ctrl+C sempre sai
		if msg.String() == "ctrl+c" {
			return m, tea.Quit
		}
		// Alternar ajuda global
		if msg.String() == "?" {
			m.showHelp = !m.showHelp
			return m, nil
		}
	}

	// Delegar para o modo atual
	switch m.mode {
	case InsertMode:
		return m.updateInsert(msg)
	case CommandMode:
		return m.updateCommand(msg)
	case NormalMode:
		return m.updateNormal(msg)
	}

	return m, nil
}
