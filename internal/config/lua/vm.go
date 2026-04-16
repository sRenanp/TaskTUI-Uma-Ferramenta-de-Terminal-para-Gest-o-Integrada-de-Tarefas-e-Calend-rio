package lua

import (
	lua "github.com/yuin/gopher-lua"
)

// VM encapsula a instância Lua para reúso
type VM struct {
	L *lua.LState
}

// NewVM cria uma nova máquina virtual Lua com os módulos padrão carregados
func NewVM() *VM {
	L := lua.NewState()

	// Carrega módulos úteis para configuração
	modules := []struct {
		name string
		f    lua.LGFunction
	}{
		{lua.LoadLibName, lua.OpenPackage},
		{lua.BaseLibName, lua.OpenBase},
		{lua.TabLibName, lua.OpenTable},
		{lua.StringLibName, lua.OpenString},
		{lua.MathLibName, lua.OpenMath},
		{lua.OsLibName, lua.OpenOs}, // Para os.getenv
	}

	for _, m := range modules {
		if err := L.CallByParam(lua.P{
			Fn:      L.NewFunction(m.f),
			NRet:    0,
			Protect: true,
		}, lua.LString(m.name)); err != nil {
			// Log se quiser, mas não falha na criação
		}
	}

	return &VM{L: L}
}

// Close libera recursos da VM
func (vm *VM) Close() {
	vm.L.Close()
}

// GetString obtém uma string global do estado Lua
func (vm *VM) GetString(name string) (string, bool) {
	val := vm.L.GetGlobal(name)
	if val == lua.LNil {
		return "", false
	}
	if str, ok := val.(lua.LString); ok {
		return string(str), true
	}
	return "", false
}

// GetTable obtém uma tabela global do estado Lua
func (vm *VM) GetTable(name string) (*lua.LTable, bool) {
	val := vm.L.GetGlobal(name)
	if val == lua.LNil {
		return nil, false
	}
	if tbl, ok := val.(*lua.LTable); ok {
		return tbl, true
	}
	return nil, false
}
