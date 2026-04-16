package config

import (
	"os"
	"path/filepath"

	"tui-task-manager/internal/config/lua" 
)

type Config struct {
	DataPath string
	Theme    string
}

func Load() (*Config, error) {
	// Defaults
	cfg := &Config{
		DataPath: "data/tasks.json",
		Theme:    "dark",
	}

	// Tenta carregar configuração Lua
	luaConfig, err := lua.LoadConfigFile("configs/config.lua")
	if err != nil {
		// Log do erro mas continua com defaults
		// logger.Warn("failed to load lua config", "error", err)
	} else if luaConfig != nil {
		// Sobrescreve defaults com valores do Lua (se existirem)
		if luaConfig.DataPath != "" {
			cfg.DataPath = luaConfig.DataPath
		}
		if luaConfig.Theme != "" {
			cfg.Theme = luaConfig.Theme
		}
	}

	// Garante que a pasta de dados existe
	dir := filepath.Dir(cfg.DataPath)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return nil, err
		}
	}

	return cfg, nil
}

// LoadFromString útil para testes
func LoadFromString(luaCode string) (*Config, error) {
	cfg := &Config{
		DataPath: "data/tasks.json",
		Theme:    "dark",
	}

	luaConfig, err := lua.LoadConfigString(luaCode)
	if err != nil {
		return nil, err
	}

	if luaConfig.DataPath != "" {
		cfg.DataPath = luaConfig.DataPath
	}
	if luaConfig.Theme != "" {
		cfg.Theme = luaConfig.Theme
	}

	return cfg, nil
}
