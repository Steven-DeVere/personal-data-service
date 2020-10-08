package articles

import (
	"fmt"
	"log"

	"github.com/devere-here/personal-data-service/datasources/mysql/articlesdb"
	"github.com/devere-here/personal-data-service/domain/errors"
)

const (
	queryInsertArticles = "INSERT INTO articles(title, blurb, content) VALUES(?, ?, ?);"
	queryGetAllArticles = "SELECT * FROM articles;"
	queryGetArticle     = "SELECT * FROM articles WHERE id=?"
)

// Only point in the application where you interact with the database

var (
	articlesDB = make(map[int64]*Article)
)

// Get retrieves an article from the database
func Get(articleID int64) (*Article, *errors.RestErr) {
	stmt, err := articlesdb.Client.Prepare(queryGetArticle)
	if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	article := Article{}
	result := stmt.QueryRow(articleID)
	if err := result.Scan(&article.ID, &article.Title, &article.Blurb, &article.Content); err != nil {
		return nil, errors.NewInternalServerError(fmt.Sprintf("error when trying to get article %d", articleID))
	}

	return &article, nil
}

// GetAll retrieves all articles from the database
func GetAll() (*[]Article, *errors.RestErr) {
	stmt, err := articlesdb.Client.Prepare(queryGetAllArticles)
	if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	articles := []Article{}
	rows, err := stmt.Query()
	if err != nil {
		return nil, errors.NewInternalServerError("error when trying to get articles")
	}
	defer rows.Close()

	for rows.Next() {
		article := Article{}
		if err := rows.Scan(&article.ID, &article.Title, &article.Blurb, &article.Content); err != nil {
			log.Fatal(err)
		}

		articles = append(articles, article)
	}

	return &articles, nil
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
