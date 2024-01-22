package data

import (
	"database/sql"
	"errors"
)

var (
	ErrRecordNotFound = errors.New("models: no matching record found")
)

type Models struct {
	Articles ArticleModel
}

func NewModels(db *sql.DB) Models {
	return Models{
		Articles: ArticleModel{DB: db},
	}
}
