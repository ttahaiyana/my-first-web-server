package storage

import (
	"fmt"
	"log"

	"github.com/ttahaiyana/my-first-web-server/internal/app/myfirstwebserver/models"
)

type ArticleRepository struct {
	storage *Storage
}

var (
	tableArticle string = "articles"
)

func (ar *ArticleRepository) Create(a *models.Article) (*models.Article, error) {
	query := fmt.Sprintf("INSERT INTO %s (title, author, content) VALUES ($1, $2, $3) RETURNING id", tableArticle)
	if err := ar.storage.db.QueryRow(query, a.Title, a.Author, a.Content).Scan(&a.ID); err != nil {
		return nil, err
	}
	return a, nil
}

func (ar *ArticleRepository) Delete(id int) (*models.Article, error) {
	return nil, nil
}

func (ar *ArticleRepository) SelectAll() ([]*models.Article, error) {
	query := fmt.Sprintf("SELECT * FROM %s", tableArticle)
	rows, err := ar.storage.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	articles := make([]*models.Article, 0)
	for rows.Next() {
		a := models.Article{}
		err := rows.Scan(&a.Title, &a.Author, &a.Content)
		if err != nil {
			log.Println(err)
			continue
		}
		articles = append(articles, &a)
	}
	return articles, nil
}

func (ar *ArticleRepository) FindByID(id int) (*models.Article, bool, error) {
	var founded bool
	var foundedArticle *models.Article
	articles, err := ar.SelectAll()
	if err != nil {
		return foundedArticle, founded, err
	}

	for _, a := range articles {
		if a.ID == id {
			foundedArticle = a
			founded = true
			break
		}
	}
	return foundedArticle, founded, nil
}
