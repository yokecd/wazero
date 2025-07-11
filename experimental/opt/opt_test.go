package opt_test

import (
	"context"
	"runtime"
	"testing"

	"github.com/yokecd/wazero"
	"github.com/yokecd/wazero/experimental/opt"
	"github.com/yokecd/wazero/internal/testing/require"
)

func TestUseOptimizingCompiler(t *testing.T) {
	if runtime.GOARCH != "arm64" {
		return
	}
	c := opt.NewRuntimeConfigOptimizingCompiler()
	r := wazero.NewRuntimeWithConfig(context.Background(), c)
	require.NoError(t, r.Close(context.Background()))
}
