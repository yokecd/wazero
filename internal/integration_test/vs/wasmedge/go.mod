module github.com/yokecd/wazero/internal/integration_test/vs/wasmedge

go 1.19

require (
	github.com/second-state/WasmEdge-go v0.12.1
	github.com/yokecd/wazero v0.0.0
)

replace github.com/yokecd/wazero => ../../../..
