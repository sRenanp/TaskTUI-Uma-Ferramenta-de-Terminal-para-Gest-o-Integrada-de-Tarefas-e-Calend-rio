package core

// StorageStrategy define o contrato para persistência de tarefas
type StorageStrategy interface {
	LoadTasks() ([]*Task, error)
	SaveTasks(tasks []*Task) error
}
