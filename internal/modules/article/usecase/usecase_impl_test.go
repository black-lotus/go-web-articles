package usecase

import (
	"context"
	"database/sql"
	"fmt"
	"testing"
	"time"
	"webarticles/internal/modules/article/domain"
	"webarticles/internal/modules/article/repository"
	article_repo_mock "webarticles/internal/modules/article/repository/interfaces"
	redis_mock "webarticles/pkg/redis"
	"webarticles/pkg/testhelper"

	"github.com/stretchr/testify/assert"
	mock "github.com/stretchr/testify/mock"
)

var (
	ctx              context.Context
	articleRepoMocks *article_repo_mock.MockArticleRepository
	storageMocks     *redis_mock.MockRedisStore
	usecaseTest      ArticleUsecase
)

func doMock() {
	// set context
	ctx = context.Background()

	// set repo
	db := &sql.DB{}
	repo := repository.NewRepository(db)
	articleRepoMocks = &article_repo_mock.MockArticleRepository{}
	repo.ArticleRepo = articleRepoMocks

	// set cache
	storageMocks = &redis_mock.MockRedisStore{}

	// set the target
	usecaseTest = NewArticleUsecase(repo, storageMocks)
}

func TestNewArticleUsecase(t *testing.T) {
	testName := testhelper.SetTestcaseName(1, "new usecase")

	t.Run(testName, func(t *testing.T) {
		doMock()

		// set usecase
		usecase := usecaseTest

		assert.NotNil(t, usecase)
	})
}

func Test_articleUsecaseImpl_FindAll(t *testing.T) {
	articles := []*domain.Article{
		&domain.Article{},
	}

	testCase := map[string]struct {
		wantError    bool
		findAll      []*domain.Article
		findAllError error
	}{
		testhelper.SetTestcaseName(1, "Given happy case"): {
			wantError:    false,
			findAll:      articles,
			findAllError: nil,
		},
		testhelper.SetTestcaseName(2, "Given repo error"): {
			wantError:    true,
			findAll:      nil,
			findAllError: fmt.Errorf("Data not found"),
		},
	}
	for name, test := range testCase {
		t.Run(name, func(t *testing.T) {
			doMock()
			if test.findAll != nil || test.findAllError != nil {
				articleRepoMocks.On("FindAll", mock.Anything, mock.Anything).Return(test.findAll, test.findAllError).Once()
			}

			// set usecase
			usecase := usecaseTest

			// run the usecase
			_, err := usecase.FindAll(ctx, &domain.Filter{})
			if test.wantError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			articleRepoMocks.AssertExpectations(t)
		})
	}
}

func Test_articleUsecaseImpl_FindByID(t *testing.T) {
	id := "1"
	article := domain.Article{}

	type redis struct {
		article *domain.Article
		Error   error
	}
	testCase := map[string]struct {
		wantError          bool
		ID                 *string
		redisGet, redisSet *redis
		find               *domain.Article
		findError          error
	}{
		testhelper.SetTestcaseName(1, "Given cache exist"): {
			wantError: false,
			ID:        &id,
			redisGet:  &redis{article: &article},
		},
		testhelper.SetTestcaseName(2, "Given cache empty"): {
			wantError: false,
			ID:        &id,
			redisGet:  &redis{Error: fmt.Errorf("Empty")},
			find:      &article,
			findError: nil,
			redisSet:  &redis{article: &article},
		},
		testhelper.SetTestcaseName(3, "Given repo error"): {
			wantError: true,
			ID:        &id,
			redisGet:  &redis{Error: fmt.Errorf("Empty")},
			find:      nil,
			findError: fmt.Errorf("Data not found"),
		},
	}
	for name, test := range testCase {
		t.Run(name, func(t *testing.T) {
			doMock()

			if test.redisGet != nil {
				var data string = mock.Anything
				var err error = test.redisGet.Error
				storageMocks.On("Get", mock.Anything, mock.Anything).Return(data, err).Once()
			}

			if test.find != nil || test.findError != nil {
				articleRepoMocks.On("FindByID", mock.Anything, mock.Anything).Return(test.find, test.findError).Once()
			}

			if test.redisSet != nil {
				var data string = mock.Anything
				var err error = test.redisSet.Error
				storageMocks.On("Set", mock.Anything, mock.Anything, mock.Anything, 60*time.Minute).Return(data, err).Once()
			}

			// set usecase
			usecase := usecaseTest

			// run the usecase
			_, err := usecase.FindByID(ctx, test.ID)
			if test.wantError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			articleRepoMocks.AssertExpectations(t)
		})
	}
}

func Test_articleUsecaseImpl_CreateArticle(t *testing.T) {
	article := domain.Article{}

	testCase := map[string]struct {
		wantError   bool
		dataUsecase *domain.Article
		insert      *domain.Article
		insertError error
	}{
		testhelper.SetTestcaseName(0, "Given happy case"): {
			wantError:   false,
			dataUsecase: &article,
			insert:      &article,
			insertError: nil,
		},
		testhelper.SetTestcaseName(1, "Given insert error"): {
			wantError:   true,
			dataUsecase: &article,
			insert:      nil,
			insertError: fmt.Errorf("Duplicate data"),
		},
	}
	for name, test := range testCase {
		t.Run(name, func(t *testing.T) {
			doMock()

			if test.insert != nil || test.insertError != nil {
				articleRepoMocks.On("Insert", mock.Anything, mock.Anything).Return(test.insert, test.insertError).Once()
			}

			// set usecase
			usecase := usecaseTest

			// run the usecase
			_, err := usecase.CreateArticle(ctx, test.dataUsecase)
			if test.wantError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			articleRepoMocks.AssertExpectations(t)
		})
	}
}
