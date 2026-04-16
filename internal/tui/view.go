package tui

import (
	"fmt"
	"strings"
	"tui-task-manager/internal/core"

	"github.com/charmbracelet/lipgloss"
)

var (
	// Estilos para os diferentes painéis
	appStyle = lipgloss.NewStyle().
			Padding(0, 1)

	titleStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("205")).
			Background(lipgloss.Color("236")).
			Bold(true).
			Padding(0, 2).
			Width(100)

	taskListStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("63")).
			Padding(0, 1).
			Width(40)

	detailStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("63")).
			Padding(0, 1).
			Width(40)

	helpStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("240")).
			Padding(0, 1).
			Width(80)

	statusStyle = lipgloss.NewStyle().
			Background(lipgloss.Color("62")).
			Foreground(lipgloss.Color("230")).
			Padding(0, 1)

	cmdStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("82")).
			Background(lipgloss.Color("236")).
			Padding(0, 1)

	cursorStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("205")).
			Bold(true)

	activeSectionStyle = lipgloss.NewStyle().
				BorderForeground(lipgloss.Color("205")).
				Border(lipgloss.DoubleBorder())

	messageStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("82")).
			Background(lipgloss.Color("236")).
			Padding(0, 2)

	errorStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("196")).
			Background(lipgloss.Color("236")).
			Padding(0, 2)
)

func (m Model) View() string {
	if !m.ready {
		return "Carregando..."
	}

	// Se help estiver ativo, mostra tela de ajuda
	if m.showHelp {
		return m.helpView()
	}

	// Cabeçalho
	title := titleStyle.Render("📅 TASK-TUI - GESTÃO INTEGRADA")

	// Corpo principal com dois painéis
	var mainContent string

	// Calcula larguras baseado no tamanho da tela
	panelWidth := (m.width - 6) / 2

	// Atualiza estilos com largura dinâmica
	taskListStyle = taskListStyle.Width(panelWidth - 4)
	detailStyle = detailStyle.Width(panelWidth - 4)

	// Aplica estilo de seção ativa
	listBorder := taskListStyle
	detailBorder := detailStyle

	if m.activeSec == TaskListSection {
		listBorder = activeSectionStyle.Width(panelWidth - 4)
	}
	if m.activeSec == DetailSection {
		detailBorder = activeSectionStyle.Width(panelWidth - 4)
	}

	// Renderiza lista de tarefas
	taskList := listBorder.Render(m.renderTaskList())

	// Renderiza detalhes da tarefa selecionada
	taskDetail := detailBorder.Render(m.renderTaskDetail())

	// Junta os painéis lado a lado
	mainContent = lipgloss.JoinHorizontal(
		lipgloss.Top,
		taskList,
		strings.Repeat(" ", 2),
		taskDetail,
	)

	// Status bar ou comando
	statusBar := m.renderStatusBar()

	// Mensagem temporária
	messageBar := m.renderMessage()

	// Junta tudo
	content := fmt.Sprintf("%s\n\n%s\n\n%s\n%s",
		title,
		mainContent,
		statusBar,
		messageBar,
	)

	return appStyle.Width(m.width).Height(m.height).Render(content)
}

func (m Model) renderTaskList() string {
	var s strings.Builder

	tasks := m.manager.GetTasks()

	if len(tasks) == 0 {
		return " Nenhuma tarefa encontrada.\n\n Pressione 'a' para adicionar."
	}

	// Calcula quantas tarefas cabem no painel
	maxItems := (m.height - 10) / 2
	startIdx := 0
	endIdx := len(tasks)

	// Scroll se necessário
	if len(tasks) > maxItems {
		if m.cursor >= startIdx+maxItems {
			startIdx = m.cursor - maxItems + 1
		} else if m.cursor < startIdx {
			startIdx = m.cursor
		}
		endIdx = min(startIdx+maxItems, len(tasks))
	}

	for i := startIdx; i < endIdx; i++ {
		task := tasks[i]

		// Indicador de cursor
		cursor := "  "
		if m.cursor == i && m.activeSec == TaskListSection {
			cursor = cursorStyle.Render("▶ ")
		} else if m.cursor == i {
			cursor = "  "
		}

		// Status visual
		statusSymbol := "○"
		switch task.Status {
		case core.InProgress:
			statusSymbol = "◔"
		case core.Done:
			statusSymbol = "●"
		}

		// Prioridade visual
		priorityColor := "250" // cinza
		switch task.Priority {
		case core.Medium:
			priorityColor = "214" // laranja
		case core.High:
			priorityColor = "196" // vermelho
		}

		priorityStyle := lipgloss.NewStyle().
			Foreground(lipgloss.Color(priorityColor))

		// Formata a linha
		line := fmt.Sprintf("%s %s %s %s",
			cursor,
			statusSymbol,
			task.Description,
			priorityStyle.Render("●"),
		)

		// Trunca se necessário
		if lipgloss.Width(line) > taskListStyle.GetWidth()-4 {
			desc := task.Description
			maxDesc := taskListStyle.GetWidth() - 10
			if len(desc) > maxDesc {
				desc = desc[:maxDesc-3] + "..."
			}
			line = fmt.Sprintf("%s %s %s %s",
				cursor,
				statusSymbol,
				desc,
				priorityStyle.Render("●"),
			)
		}

		s.WriteString(line)
		s.WriteString("\n")
	}

	return s.String()
}

func (m Model) renderTaskDetail() string {
	tasks := m.manager.GetTasks()
	if len(tasks) == 0 || m.selectedTaskIdx >= len(tasks) {
		return " Selecione uma tarefa para ver detalhes"
	}

	task := tasks[m.selectedTaskIdx]

	var s strings.Builder
	s.WriteString(fmt.Sprintf(" 📝 Tarefa #%s\n\n", task.ID))

	// Descrição
	if m.activeSec == DetailSection && m.cursor == 0 {
		if m.mode == InsertMode {
			s.WriteString("   Descrição: " + m.textInput.View() + "\n")
		} else {
			s.WriteString(cursorStyle.Render("▶ ") + "Descrição: " + task.Description + "\n")
		}
	} else {
		s.WriteString("   Descrição: " + task.Description + "\n")
	}

	// Status
	statusCursor := "  "
	if m.activeSec == DetailSection && m.cursor == 1 {
		statusCursor = cursorStyle.Render("▶ ")
	}

	statusStr := "A Fazer"
	switch task.Status {
	case core.InProgress:
		statusStr = "Em Andamento"
	case core.Done:
		statusStr = "Concluído"
	}
	s.WriteString(fmt.Sprintf("%s Status: %s\n", statusCursor, statusStr))

	if m.mode == InsertMode && m.cursor == 1 {
		s.WriteString("\n [1] A Fazer  [2] Em Andamento  [3] Concluído")
	}

	if m.mode == InsertMode && m.cursor == 2 {
		s.WriteString("\n [1] Baixa  [2] Média  [3] Alta")
	}

	// Prioridade
	priorityCursor := "  "
	if m.activeSec == DetailSection && m.cursor == 2 {
		priorityCursor = cursorStyle.Render("▶ ")
	}

	priorityStr := "Baixa"
	switch task.Priority {
	case core.Medium:
		priorityStr = "Média"
	case core.High:
		priorityStr = "Alta"
	}
	s.WriteString(fmt.Sprintf("%s Prioridade: %s\n", priorityCursor, priorityStr))

	// Data
	if !task.EndDate.IsZero() {
		s.WriteString(fmt.Sprintf("\n 📅 Data: %s\n", task.EndDate.Format("02/01/2006")))
	}

	// Indicador de modo insert
	if m.mode == InsertMode && m.activeSec == DetailSection {
		s.WriteString("\n -- INSERT --")
	}

	return s.String()
}

func (m Model) renderStatusBar() string {
	var mode string
	var help string

	switch m.mode {
	case NormalMode:
		mode = " NORMAL "
		help = "?:help  tab:next  j/k:move  a:add  q:quit"
	case CommandMode:
		return cmdStyle.Render(":" + m.command)
	case InsertMode:
		mode = " INSERT "
		help = "esc:back  enter:save"
	}

	// Mostra seção ativa
	section := "Lista"
	switch m.activeSec {
	case TaskListSection:
		section = "Lista"
	case DetailSection:
		section = "Detalhes"
	case HelpSection:
		section = "Ajuda"
	}

	left := statusStyle.Render(mode) + "  " + help
	right := fmt.Sprintf("Seção: %s ", section)

	// Calcula espaçamento
	spaces := m.width - lipgloss.Width(left) - lipgloss.Width(right) - 4

	return left + strings.Repeat(" ", max(0, spaces)) + right
}

func (m Model) renderMessage() string {
	if m.message == "" {
		return ""
	}

	if strings.Contains(m.message, "Erro") {
		return errorStyle.Render(m.message)
	}
	return messageStyle.Render(m.message)
}

func (m Model) helpView() string {
	help := lipgloss.JoinVertical(lipgloss.Center,
		"",
		"📖 AJUDA - TASK-TUI",
		"",
		"NAVEGAÇÃO:",
		"  tab/shift+tab  - Alternar entre painéis",
		"  j/k            - Mover para cima/baixo",
		"  h/l            - Navegação horizontal (nos detalhes)",
		"  enter          - Abrir detalhes / entrar em modo insert",
		"  esc            - Voltar ao modo normal",
		"",
		"AÇÕES:",
		"  a              - Adicionar nova tarefa",
		"  d              - Deletar tarefa atual",
		"  :              - Modo comando",
		"  ?              - Mostrar/ocultar esta ajuda",
		"  q              - Sair",
		"",
		"COMANDOS:",
		"  :q             - Sair",
		"  :w             - Salvar",
		"  :add <desc>    - Adicionar tarefa",
		"  :delete        - Deletar tarefa atual",
		"",
		"Pressione 'q' ou ESC para voltar",
		"",
	)

	return helpStyle.Width(m.width - 4).Height(m.height - 4).Render(help)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
