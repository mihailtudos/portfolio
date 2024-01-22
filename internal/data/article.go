package data

import (
	"database/sql"
	"errors"
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
	if id < 1 {
		return nil, ErrRecordNotFound
	}

	stm := `
		SELECT title, content, version, created_at
		FROM articles
		WHERE id = $1
	`

	var article Article
	err := a.DB.QueryRow(stm, id).Scan(&article.Title, &article.Content, &article.Version, &article.CreatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return &Article{}, ErrRecordNotFound
		} else {
			return &Article{}, err
		}

	}

	return &article, nil
}

func (a ArticleModel) GetAll() ([]Article, error) {
	stm := `
		SELECT id, title, content, version, created_at
		FROM articles
		ORDER BY id DESC LIMIT 10
	`

	rows, err := a.DB.Query(stm)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var articles []Article

	for rows.Next() {
		var article Article
		err := rows.Scan(&article.ID, &article.Title, &article.Content, &article.Version, &article.CreatedAt)
		if err != nil {
			return nil, err
		}

		articles = append(articles, article)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return articles, nil
}

func (a ArticleModel) Update(article *Article) error {
	return nil
}

func (a ArticleModel) Delete(id int64) error {
	if id < 0 {
		return ErrRecordNotFound
	}

	stm := `
		DELETE * 
		FROM articles
		WHERE id = $1
	`

	_, err := a.DB.Exec(stm, id)
	if err != nil {
		return ErrRecordNotFound
	}

	return nil
}
