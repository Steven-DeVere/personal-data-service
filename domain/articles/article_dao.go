package articles

import (
	"fmt"

	"github.com/devere-here/personal-data-service/datasources/mysql/articles_db"
	"github.com/devere-here/personal-data-service/domain/errors"
)

// Only point in the application where you interact with the database

var (
	articlesDB = make(map[int64]*Article)
)

// Get retrieves an article from the database
func Get(articleID int64) (*Article, *errors.RestErr) {
	if err := articles_db.Client.Ping(); err != nil {
		panic(err)
	}

	result := articlesDB[articleID]
	if result == nil {
		return nil, errors.NewNotFoundError(fmt.Sprintf("article %d not found", articleID))
	}

	return result, nil
}

// GetAll retrieves all articles from the database
func GetAll() ([]Article, *errors.RestErr) {
	var articles []Article
	for _, article := range articlesDB {
		if article == nil {
			return nil, errors.NewInternalDBError("Internal DB error, DB contains nil value for article")
		}
		articles = append(articles, *article)
	}

	return articles, nil
}

// Update updates an article from the database
func (article *Article) Update() *errors.RestErr {
	result := articlesDB[article.ID]
	if result == nil {
		return errors.NewNotFoundError(fmt.Sprintf("article %d not found", article.ID))
	}

	articlesDB[article.ID] = article
	return nil
}

// GetAllByCategories retrieves all articles from the database that have a one of the provided categories
func GetAllByCategories(categories []string) ([]Article, *errors.RestErr) {
	var articles []Article
	for _, article := range articlesDB {
		if article == nil {
			return nil, errors.NewInternalDBError("Internal DB error, DB contains nil value for article")
		}

		for _, dbCategory := range article.Categories {
			for _, category := range categories {
				if category == dbCategory {
					articles = append(articles, *article)
					break
				}
			}
		}
	}

	return articles, nil
}

// Save saves an article to the database
func (article *Article) Save() *errors.RestErr {
	current := articlesDB[article.ID]
	if current != nil {
		if current.Title == article.Title {
			return errors.NewBadRequestError(fmt.Sprintf("title %s already exists", article.Title))
		}

		return errors.NewBadRequestError(fmt.Sprintf("article %d already exists", article.ID))
	}

	articlesDB[article.ID] = article
	return nil
}

// Delete removes an article from the database
func Delete(articleID int64) (*Article, *errors.RestErr) {
	current := articlesDB[articleID]
	if current == nil {
		return nil, errors.NewBadRequestError(fmt.Sprintf("article %d does not exist", articleID))
	}

	delete(articlesDB, articleID)
	return current, nil
}
