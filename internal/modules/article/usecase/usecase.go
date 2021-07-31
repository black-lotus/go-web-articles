package usecase

import (
	"context"
	"webarticles/internal/modules/article/domain"
)

// ArticleUsecase abstract interface
type ArticleUsecase interface {
	FindAll(ctx context.Context, filter *domain.Filter) ([]*domain.Article, error)
	FindByID(ctx context.Context, ID *string) (*domain.Article, error)
	CreateArticle(ctx context.Context, article *domain.Article) (*domain.Article, error)
}
