package mysql

import (
	"context"
	"database/sql"
	"webarticles/internal/modules/article/domain"
	"webarticles/internal/modules/article/repository/interfaces"
)

type articleRepoMysql struct {
	db *sql.DB
}

// NewArticleRepo create new rule repository mysql
func NewArticleRepo(db *sql.DB) interfaces.ArticleRepository {
	return &articleRepoMysql{db}
}

func (a *articleRepoMysql) FindAll(ctx context.Context, filter *domain.Filter) ([]*domain.Article, error) {
	return nil, nil
}

func (a *articleRepoMysql) FindByID(tx context.Context, ID *string) (*domain.Article, error) {
	return nil, nil
}

func (a *articleRepoMysql) Save(ctx context.Context, data *domain.Article) (*domain.Article, error) {
	return nil, nil
}

func (a *articleRepoMysql) Insert(ctx context.Context, newData *domain.Article) (*domain.Article, error) {
	return nil, nil
}
