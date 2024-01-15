package data

import (
	"database/sql"
	"time"
)

type Article struct {
	ID        int64     `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Version   int32     `json:"version"`
	CreatedAt time.Time `json:"created_at"`
}

type ArticleModel struct {
	DB *sql.DB
}

func (a ArticleModel) Insert(article *Article) error {
	stm := `
		INSERT INTO articles (title, content)
		VALUES ($1, $2)
		RETURNING id, created_at, version
	`

	return a.DB.QueryRow(stm, article.Title, article.Content).Scan(&article.ID, &article.CreatedAt, &article.Version)
}

func (a ArticleModel) Get(id int64) (*Article, error) {
	return &Article{}, nil
}

func (a ArticleModel) Update(article *ArticleModel) error {
	return nil
}

func (a ArticleModel) Delete(id int64) error {
	return nil
}
