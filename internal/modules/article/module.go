package article

import (
	"webarticles/internal/modules/article/delivery/resthandler"
	"webarticles/internal/modules/article/repository"
	"webarticles/internal/modules/article/usecase"
	"webarticles/pkg/codebase/factory/dependency"
	"webarticles/pkg/codebase/factory/types"
	"webarticles/pkg/codebase/interfaces"
)

const (
	// Name module name
	Name types.Module = "Article"
)

// Module model
type Module struct {
	restHandler interfaces.EchoRestHandler
}

// NewModule module constructor
func NewModule(deps dependency.Dependency) *Module {
	repo := repository.NewRepository(deps.GetSQLDatabase().GetSQLDB())
	uc := usecase.NewArticleUsecase(repo, deps.GetRedisPool().Store())

	var mod Module
	mod.restHandler = resthandler.NewRestHandler(
		uc,
	)

	return &mod
}

// RestHandler method
func (m *Module) RestHandler() interfaces.EchoRestHandler {
	return m.restHandler
}

// Name get module name
func (m *Module) Name() types.Module {
	return Name
}
