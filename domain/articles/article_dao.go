package articles

import (
	"fmt"

	"github.com/devere-here/personal-data-service/datasources/mysql/articlesdb"
	"github.com/devere-here/personal-data-service/domain/errors"
)

const (
	queryInsertArticles = "INSERT INTO articles(title, blurb, content) VALUES(?, ?, ?);"
)

// Only point in the application where you interact with the database

var (
	articlesDB = make(map[int64]*Article)
)

// Get retrieves an article from the database
func Get(articleID int64) (*Article, *errors.RestErr) {
	if err := articlesdb.Client.Ping(); err != nil {
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
			return nil, errors.NewInternalServerError("Internal DB error, DB contains nil value for article")
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

// Save saves an article to the database
func (article *Article) Save() *errors.RestErr {
	stmt, err := articlesdb.Client.Prepare(queryInsertArticles)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	insertResult, err := stmt.Exec(article.Title, article.Blurb, article.Content)
	if err != nil {
		return errors.NewInternalServerError(fmt.Sprintf("error when trying to save user: %s", err.Error()))
	}

	articleID, err := insertResult.LastInsertId()
	if err != nil {
		return errors.NewInternalServerError(fmt.Sprintf("error when trying to save user: %s", err.Error()))
	}
	article.ID = articleID

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
