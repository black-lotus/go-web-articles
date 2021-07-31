package factory

import (
	"webarticles/pkg/codebase/factory/dependency"
	"webarticles/pkg/codebase/factory/types"
)

// ServiceFactory factory
type ServiceFactory interface {
	GetDependency() dependency.Dependency
	GetModules() []ModuleFactory
	Name() types.Service
}
