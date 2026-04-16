package tui

import (
	"tui-task-manager/internal/core"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type Mode int

const (
	NormalMode Mode = iota
	CommandMode
	InsertMode
)

type Section int

const (
	TaskListSection Section = iota
	DetailSection
	HelpSection
)

type Model struct {
	manager         *core.Manager
	mode            Mode
	activeSec       Section // Seção ativa
	command         string
	cursor          int
	selectedTaskIdx int
	textInput       textinput.Model
	width           int
	height          int
	ready           bool
	showHelp        bool
	message         string // Mensagens temporárias
	messageTimer    int
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m *Model) OnUpdate(tasks []*core.Task) {
	// Atualiza quando o manager notificar mudanças
}

func NewModel(mgr *core.Manager) Model {
	ti := textinput.New()
	ti.Placeholder = "Nova descrição..."
	ti.CharLimit = 50
	ti.Width = 30

	m := Model{
		manager:   mgr,
		mode:      NormalMode,
		activeSec: TaskListSection,
		textInput: ti,
	}
	mgr.RegisterObserver(&m)
	return m
}
