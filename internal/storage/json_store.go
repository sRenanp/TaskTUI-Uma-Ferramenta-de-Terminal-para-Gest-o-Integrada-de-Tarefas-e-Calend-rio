package storage

import (
	"encoding/json"
	"os"
	"path/filepath"
	"tui-task-manager/internal/core"
)

type JSONStore struct {
	filePath string
}

func NewJSONStore(path string) *JSONStore {
	absPath, _ := filepath.Abs(path)
	return &JSONStore{filePath: absPath}
}

func (s *JSONStore) LoadTasks() ([]*core.Task, error) {
	if _, err := os.Stat(s.filePath); os.IsNotExist(err) {
		return []*core.Task{}, nil
	}

	data, err := os.ReadFile(s.filePath)
	if err != nil {
		return nil, err
	}

	if len(data) == 0 {
		return []*core.Task{}, nil
	}

	var tasks []*core.Task
	if err := json.Unmarshal(data, &tasks); err != nil {
		return nil, err
	}

	return tasks, nil
}

func (s *JSONStore) SaveTasks(tasks []*core.Task) error {
	dir := filepath.Dir(s.filePath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	// MarshalIndent facilita a leitura humana se o usuário quiser editar o JSON manualmente
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(s.filePath, data, 0644)
}
