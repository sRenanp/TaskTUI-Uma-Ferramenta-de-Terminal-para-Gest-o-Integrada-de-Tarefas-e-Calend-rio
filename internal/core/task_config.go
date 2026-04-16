package core

// TaskConfig será implementado depois da primeira versão
type TaskConfig struct {
	// Vazio
}

// DefaultTaskConfig retorna configuração padrão
func DefaultTaskConfig() *TaskConfig {
	return &TaskConfig{}
}
