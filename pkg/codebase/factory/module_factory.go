package factory

import (
	"webarticles/pkg/codebase/factory/types"
	"webarticles/pkg/codebase/interfaces"
)

// ModuleFactory factory
type ModuleFactory interface {
	RestHandler() interfaces.EchoRestHandler
	Name() types.Module
}
