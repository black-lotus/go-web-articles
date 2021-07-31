package usecase

import (
	"context"
	"webarticles/internal/modules/article/domain"
	"webarticles/internal/modules/article/repository"
	"webarticles/pkg/codebase/interfaces"
)

//articleUsecaseImpl structure
type articleUsecaseImpl struct {
	repo  *repository.Repository
	cache interfaces.Store
}

// NewAuditTrailUsecase create new audit trail usecase
func NewArticleUsecase(repo *repository.Repository, cache interfaces.Store) ArticleUsecase {
	return &articleUsecaseImpl{repo: repo, cache: cache}
}

func (a *articleUsecaseImpl) FindAll(ctx context.Context, filter *domain.Filter) ([]*domain.Article, error) {
	return nil, nil
}

func (a *articleUsecaseImpl) FindByID(ctx context.Context, ID *string) (*domain.Article, error) {
	return nil, nil
}

func (a *articleUsecaseImpl) CreateArticle(ctx context.Context, article *domain.Article) (*domain.Article, error) {
	return nil, nil
}
