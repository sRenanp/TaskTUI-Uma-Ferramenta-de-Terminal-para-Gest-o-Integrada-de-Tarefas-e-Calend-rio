package core

// Command define a interface para todas as ações do sistema
type Command interface {
	Execute() error
	// Undo() error // Deixando espaço para um futuro CTRL+Z
}

type AddTaskCommand struct {
	Manager *Manager
	Task    *Task
}

func (c *AddTaskCommand) Execute() error {
	return c.Manager.addTask(c.Task)
}

type MarkCompleteCommand struct {
	Manager *Manager
	TaskID  string
}

func (c *MarkCompleteCommand) Execute() error {
	return c.Manager.markTaskComplete(c.TaskID)
}

type DeleteTaskCommand struct {
	Manager *Manager
	TaskID  string
}

func (c *DeleteTaskCommand) Execute() error {
	return c.Manager.deleteTask(c.TaskID)
}
