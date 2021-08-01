package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"webarticles/internal/modules/article/domain"
	"webarticles/internal/modules/article/repository/interfaces"
	"webarticles/pkg/logger"
)

type articleRepoMysql struct {
	db *sql.DB
}

// NewArticleRepo create new article repository mysql
func NewArticleRepo(db *sql.DB) interfaces.ArticleRepository {
	return &articleRepoMysql{db}
}

func (a *articleRepoMysql) FindAll(ctx context.Context, filter *domain.Filter) (result []*domain.Article, err error) {
	var sb strings.Builder
	if filter.Query != "" {
		sb.WriteString(fmt.Sprintf(`AND title LIKE '%%%s%%' OR body LIKE '%%%s%%'`, filter.Query, filter.Query))
	}

	if filter.Author != "" {
		sb.WriteString(fmt.Sprintf(`AND author = '%s'`, filter.Author))
	}

	query := fmt.Sprintf("SELECT id, author, title, body, is_deleted, created "+
		"FROM article WHERE is_deleted=%v %s ORDER BY id DESC",
		0, sb.String())

	var stmt *sql.Stmt
	stmt, err = a.db.Prepare(query)
	if err != nil {
		logger.LogRed(err.Error())
		return
	}
	defer stmt.Close()

	var rows *sql.Rows
	rows, err = stmt.QueryContext(ctx)
	if err != nil {
		logger.LogRed(err.Error())
		return
	}

	defer rows.Close()

	for rows.Next() {
		article := domain.Article{}
		if err = rows.Scan(
			&article.ID,
			&article.Author,
			&article.Title,
			&article.Body,
			&article.IsDeleted,
			&article.Created,
		); err != nil {
			logger.LogRed(err.Error())
			return
		}
		result = append(result, &article)
	}

	return result, nil
}

func (a *articleRepoMysql) FindByID(ctx context.Context, ID *string) (*domain.Article, error) {
	query := "SELECT id, author, title, body, is_deleted, created FROM article WHERE is_deleted=? AND id=? LIMIT 1"
	stmt, err := a.db.Prepare(query)
	if err != nil {
		logger.LogRed(err.Error())
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.QueryContext(ctx, 0, ID)
	if err != nil {
		logger.LogRed(err.Error())
		return nil, err
	}

	defer rows.Close()
	if !rows.Next() {
		logger.LogRed(err.Error())
		return nil, err
	}

	article := domain.Article{}
	if err = rows.Scan(
		&article.ID,
		&article.Author,
		&article.Title,
		&article.Body,
		&article.IsDeleted,
		&article.Created,
	); err != nil {
		logger.LogRed(err.Error())
		return nil, err
	}

	return &article, nil
}

func (a *articleRepoMysql) Save(ctx context.Context, data *domain.Article) (*domain.Article, error) {
	query := "UPDATE article SET author=?, title=?, body=?, is_deleted=?, created=? WHERE id=?"
	stmt, err := a.db.Prepare(query)
	if err != nil {
		logger.LogRed(err.Error())
		return nil, err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, data.Author, data.Title, data.Body, data.IsDeleted, data.Created, data.ID)
	if err != nil {
		logger.LogRed(err.Error())
		return nil, err
	}

	return data, nil
}

func (a *articleRepoMysql) Insert(ctx context.Context, newData *domain.Article) (*domain.Article, error) {
	query := "INSERT INTO article (author, title, body, is_deleted, created) VALUES (?, ?, ?, ?, ?)"
	stmt, err := a.db.Prepare(query)
	if err != nil {
		logger.LogRed(err.Error())
		return nil, err
	}
	defer stmt.Close()

	res, err := stmt.ExecContext(ctx, newData.Author, newData.Title, newData.Body, newData.IsDeleted, newData.Created)
	if err != nil {
		logger.LogRed(err.Error())
		return nil, err
	}

	lastInsertID, err := res.LastInsertId()
	if err != nil {
		logger.LogRed(err.Error())
		return nil, err
	}

	newData.ID = lastInsertID

	return newData, nil
}
