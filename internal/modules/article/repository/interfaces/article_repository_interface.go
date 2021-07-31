package interfaces

import (
	"context"
	"webarticles/internal/modules/article/domain"
)

// ArticleRepository abstract interface
type ArticleRepository interface {
	FindAll(ctx context.Context, filter *domain.Filter) ([]*domain.Article, error)
	FindByID(tx context.Context, ID *string) (*domain.Article, error)
	Save(ctx context.Context, data *domain.Article) (*domain.Article, error)
	Insert(ctx context.Context, newData *domain.Article) (*domain.Article, error)
}
