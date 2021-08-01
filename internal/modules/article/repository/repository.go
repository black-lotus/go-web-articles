package repository

import (
	"database/sql"
	"webarticles/internal/modules/article/repository/interfaces"
	"webarticles/internal/modules/article/repository/mysql"
)

// Repository parent
type Repository struct {
	db          *sql.DB
	ArticleRepo interfaces.ArticleRepository
}

// NewRepository create new repository
func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		db:          db,
		ArticleRepo: mysql.NewArticleRepo(db),
	}
}
