package usecase

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
	"webarticles/internal/modules/article/domain"
	"webarticles/internal/modules/article/repository"
	"webarticles/pkg/codebase/interfaces"

	config "webarticles/configs"
)

const (
	cachePattern string = "%s_%s_%s"
	tableArticle string = "article"
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
	return a.repo.ArticleRepo.FindAll(ctx, filter)
}

func (a *articleUsecaseImpl) FindByID(ctx context.Context, ID *string) (*domain.Article, error) {
	// check ID from cache redis. if found return!
	serviceName := config.GetEnv().ServicePath
	cacheKey := fmt.Sprintf(cachePattern, serviceName, tableArticle, *ID)
	if cacheData, err := a.cache.Get(ctx, cacheKey); err == nil {
		var article domain.Article
		json.Unmarshal([]byte(cacheData), &article)
		return &article, nil
	}

	article, err := a.repo.ArticleRepo.FindByID(ctx, ID)
	if err != nil {
		return nil, err
	}

	// set to redis
	b, _ := json.Marshal(article)
	a.cache.Set(ctx, cacheKey, b, 60*time.Minute)

	return article, nil
}

func (a *articleUsecaseImpl) CreateArticle(ctx context.Context, article *domain.Article) (*domain.Article, error) {
	article.Created = time.Now()
	return a.repo.ArticleRepo.Insert(ctx, article)
}
