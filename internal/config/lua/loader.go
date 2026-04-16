package lua

import (
	"fmt"
	"os"
)

// ConfigData representa os dados extraídos do Lua
type ConfigData struct {
	DataPath string
	Theme    string
	// Adicione outros campos conforme necessário
}

// LoadConfigFile carrega e executa um arquivo Lua, retornando os valores extraídos
func LoadConfigFile(path string) (*ConfigData, error) {
	// Verifica se arquivo existe
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return nil, nil // Arquivo não existe não é erro
	}

	// Cria VM
	vm := NewVM()
	defer vm.Close()

	// Executa arquivo
	if err := vm.L.DoFile(path); err != nil {
		return nil, fmt.Errorf("erro executando %s: %w", path, err)
	}

	// Extrai valores
	config := &ConfigData{}

	if val, ok := vm.GetString("data_path"); ok {
		config.DataPath = val
	}

	if val, ok := vm.GetString("theme"); ok {
		config.Theme = val
	}

	return config, nil
}

// LoadConfigString carrega configuração de uma string Lua (útil para testes)
func LoadConfigString(luaCode string) (*ConfigData, error) {
	vm := NewVM()
	defer vm.Close()

	if err := vm.L.DoString(luaCode); err != nil {
		return nil, fmt.Errorf("erro executando código Lua: %w", err)
	}

	config := &ConfigData{}

	if val, ok := vm.GetString("data_path"); ok {
		config.DataPath = val
	}

	if val, ok := vm.GetString("theme"); ok {
		config.Theme = val
	}

	return config, nil
}
