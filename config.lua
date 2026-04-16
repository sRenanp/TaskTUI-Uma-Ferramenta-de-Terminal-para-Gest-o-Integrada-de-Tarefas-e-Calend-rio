-- Configuração do usuário
data_path = "data/tasks.json"
theme = "dark"

-- Pode usar lógica condicional
if os.getenv("DEV_MODE") == "1" then
    theme = "light"
    -- Ativa debug
    debug_mode = true
end

-- Pode calcular caminhos dinâmicos
if os.getenv("XDG_DATA_HOME") then
    data_path = os.getenv("XDG_DATA_HOME") .. "/myapp/tasks.json"
end