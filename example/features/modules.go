package features

import (
	"github.com/thyagobudal/go-base/fiberfx"
)

// GetModules returns all feature modules
func GetModules() []*fiberfx.Module {
	return []*fiberfx.Module{
		NewHealthFeature(),
		NewErrorSimulationFeature(),
		NewLoggerExampleFeature(),
	}
}
