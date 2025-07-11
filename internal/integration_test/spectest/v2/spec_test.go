package v2

import (
	"context"
	"runtime"
	"testing"

	"github.com/yokecd/wazero"
	"github.com/yokecd/wazero/api"
	"github.com/yokecd/wazero/experimental/opt"
	"github.com/yokecd/wazero/internal/integration_test/spectest"
	"github.com/yokecd/wazero/internal/platform"
)

const enabledFeatures = api.CoreFeaturesV2

func TestCompiler(t *testing.T) {
	if !platform.CompilerSupported() {
		t.Skip()
	}
	spectest.Run(t, Testcases, context.Background(), wazero.NewRuntimeConfigCompiler().WithCoreFeatures(enabledFeatures))
}

func TestInterpreter(t *testing.T) {
	spectest.Run(t, Testcases, context.Background(), wazero.NewRuntimeConfigInterpreter().WithCoreFeatures(enabledFeatures))
}

func TestWazevo(t *testing.T) {
	if runtime.GOARCH != "arm64" {
		t.Skip()
	}
	c := opt.NewRuntimeConfigOptimizingCompiler().WithCoreFeatures(enabledFeatures)
	spectest.Run(t, Testcases, context.Background(), c)
}
