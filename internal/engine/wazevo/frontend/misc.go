package frontend

import (
	"github.com/yokecd/wazero/internal/engine/wazevo/ssa"
	"github.com/yokecd/wazero/internal/wasm"
)

func FunctionIndexToFuncRef(idx wasm.Index) ssa.FuncRef {
	return ssa.FuncRef(idx)
}
