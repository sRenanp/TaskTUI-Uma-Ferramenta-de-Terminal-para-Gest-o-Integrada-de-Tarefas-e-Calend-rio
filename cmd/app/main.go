package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"tui-task-manager/internal/config"
	"tui-task-manager/internal/core"
	"tui-task-manager/internal/storage"
	"tui-task-manager/internal/tui"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
)

// initializeApp centraliza a criação do manager e storage para evitar repetição
func initializeApp() (*core.Manager, error) {
	cfg, err := config.Load()
	if err != nil {
		return nil, err
	}

	// Cria o storage usando o caminho da configuração
	store := storage.NewJSONStore(cfg.DataPath)

	// Inicializa o manager com a estratégia de storage
	return core.NewManager(store), nil
}

func main() {
	// Comando raiz: Executa a TUI se nenhum subcomando for passado
	var rootCmd = &cobra.Command{
		Use:   "tui-task-manager",
		Short: "Gestor de tarefas com interface TUI e CLI",
		Run: func(cmd *cobra.Command, args []string) {
			manager, err := initializeApp()
			if err != nil {
				log.Fatal(err)
			}

			// Inicia a TUI
			model := tui.NewModel(manager)
			program := tea.NewProgram(model, tea.WithAltScreen())

			if _, err := program.Run(); err != nil {
				log.Fatalf("Erro ao rodar a TUI: %v", err)
			}
		},
	}

	// Subcomando: Adicionar tarefa via CLI
	var addCmd = &cobra.Command{
		Use:   "add [descrição]",
		Short: "Adiciona uma nova tarefa via linha de comando",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			manager, err := initializeApp()
			if err != nil {
				log.Fatal(err)
			}

			description := strings.Join(args, " ")
			task := &core.Task{
				Description: description,
			}

			// Utiliza o padrão Command para a ação
			addCommand := &core.AddTaskCommand{
				Manager: manager,
				Task:    task,
			}

			if err := manager.ExecuteCommand(addCommand); err != nil {
				fmt.Printf("Erro ao adicionar tarefa: %v\n", err)
				os.Exit(1)
			}

			fmt.Printf("✅ Tarefa '%s' adicionada com sucesso!\n", description)
		},
	}

	// Adiciona os subcomandos ao comando raiz
	rootCmd.AddCommand(addCmd)

	// Executa a aplicação
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
