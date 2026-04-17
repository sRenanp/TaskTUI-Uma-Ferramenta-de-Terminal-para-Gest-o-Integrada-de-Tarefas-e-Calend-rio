# Create a README.md (standard for GitHub) for the TUI Task Manager project based on the uploaded files.

readme_content = """# TUI Task Manager

Um gerenciador de tarefas robusto baseado em terminal (TUI) escrito em Go, utilizando a biblioteca Bubble Tea. O projeto foca em produtividade através de comandos rápidos e uma interface intuitiva diretamente no terminal.

## 🚀 Funcionalidades

- **Gerenciamento de Tarefas:** Adicione, edite e remova tarefas com facilidade.
- **Interface TUI:** Construído com o ecossistema [Charmbracelet](https://charm.sh/) (Bubble Tea, Lip Gloss, Bubbles).
- **Sistema de Comandos:** Suporte a comandos estilo Vim para agilidade.
- **Persistência:** Arquitetura preparada para diferentes estratégias de armazenamento (Storage Strategy).
- **Organização:** Prioridades, status (Todo, In Progress, Done) e tags.

## 🛠️ Tecnologias

- **Go (Golang)**
- [Bubble Tea](https://github.com/charmbracelet/bubbletea): Framework TUI.
- [Lip Gloss](https://github.com/charmbracelet/lipgloss): Estilização de layout no terminal.
- [UUID](https://github.com/google/uuid): Identificação única de tarefas.

## 📂 Estrutura do Projeto

- `/internal/core`: Lógica de negócio, modelos de dados (`Task`), e padrões de design como Command e Strategy.
- `/internal/tui`: Implementação da interface, estados (Normal, Insert, Command) e renderização.

## ⌨️ Atalhos e Comandos

### Navegação

- `tab` / `shift+tab`: Alternar entre painéis.
- `j` / `k`: Mover seleção para cima/baixo.
- `enter`: Abrir detalhes ou entrar no modo de edição.

### Ações

- `a`: Adicionar nova tarefa.
- `d`: Deletar tarefa selecionada.
- `:`: Entrar no modo de comando.
- `?`: Mostrar/ocultar ajuda.
- `q`: Sair da aplicação.

## 🏗️ Como Executar

1. Certifique-se de ter o Go instalado em sua máquina.
2. Clone o repositório:
   ```bash
   git clone <url-do-repositorio>
   ```
