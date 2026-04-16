// Um passo importante, implementar configurações de task depois

package core

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

// Observer interface para notificar a TUI sobre mudanças
type Observer interface {
	OnUpdate(tasks []*Task)
}

type Manager struct {
	tasks     []*Task
	storage   StorageStrategy
	observers []Observer
	config    *TaskConfig
}

func NewManager(storage StorageStrategy) *Manager {
	tasks, _ := storage.LoadTasks()
	return &Manager{
		tasks:     tasks,
		storage:   storage,
		observers: []Observer{},
	}
}

// ExecuteCommand executa qualquer comando
func (m *Manager) ExecuteCommand(cmd Command) error {
	return cmd.Execute()
}

type UpdateTaskCommand struct {
	Manager *Manager
	TaskID  string
	Updates map[string]interface{}
}

func (c *UpdateTaskCommand) Execute() error {
	return c.Manager.updateTask(c.TaskID, c.Updates)
}

// Placeholder para setar configurações
func (m *Manager) SetConfig(config *TaskConfig) {
	// Vazio por enquanto - não faz nada
}

// addTask adiciona uma nova tarefa
func (m *Manager) addTask(t *Task) error {
	if t.Description == "" {
		return errors.New("descrição vazia")
	}

	// Se não tiver ID, gera um novo
	if t.ID == "" {
		t.ID = uuid.New().String()
	}

	now := time.Now()
	t.CreatedAt = now
	t.UpdatedAt = now

	m.tasks = append(m.tasks, t)
	m.notify()
	return nil
}

// markTaskComplete marca uma tarefa como concluída
func (m *Manager) markTaskComplete(id string) error {
	for _, t := range m.tasks {
		if t.ID == id {
			t.Status = Done
			t.UpdatedAt = time.Now()
			m.notify()
			return nil
		}
	}
	return errors.New("tarefa não encontrada")
}

// updateTask atualiza uma tarefa existente
func (m *Manager) updateTask(id string, updates map[string]interface{}) error {
	for _, t := range m.tasks {
		if t.ID == id {
			if desc, ok := updates["description"]; ok {
				t.Description = desc.(string)
			}
			if priority, ok := updates["priority"]; ok {
				t.Priority = priority.(Priority)
			}
			if deadline, ok := updates["deadline"]; ok {
				t.EndDate = deadline.(time.Time)
			}
			if tags, ok := updates["tags"]; ok {
				t.Tags = tags.([]string)
			}

			t.UpdatedAt = time.Now()
			m.notify()
			return nil
		}
	}
	return errors.New("tarefa não encontrada")
}

// deleteTask remove uma tarefa
func (m *Manager) deleteTask(id string) error {
	for i, t := range m.tasks {
		if t.ID == id {
			m.tasks = append(m.tasks[:i], m.tasks[i+1:]...)
			m.notify()
			return nil
		}
	}
	return errors.New("tarefa não encontrada")
}

// ============ MÉTODOS DE CONSULTA ============

// GetTasks retorna todas as tarefas
func (m *Manager) GetTasks() []*Task {
	return m.tasks
}

// FindTask busca uma tarefa por ID
func (m *Manager) FindTask(id string) *Task {
	for _, t := range m.tasks {
		if t.ID == id {
			return t
		}
	}
	return nil
}

// ============ OBSERVER ============

func (m *Manager) RegisterObserver(o Observer) {
	m.observers = append(m.observers, o)
}

func (m *Manager) notify() {
	m.storage.SaveTasks(m.tasks)
	for _, o := range m.observers {
		o.OnUpdate(m.tasks)
	}
}
