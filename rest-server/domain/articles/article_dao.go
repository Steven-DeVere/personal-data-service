package articles

import (
	"fmt"
	"log"

	"github.com/devere-here/personal-data-service/rest-server/datasources/mysql/articlesdb"
	"github.com/devere-here/personal-data-service/rest-server/domain/errors"
)

const (
	queryInsertArticles = "INSERT INTO articles(title, articleUrl, imageUrl) VALUES(?, ?, ?);"
	queryGetAllArticles = "SELECT * FROM articles;"
	queryGetArticle     = "SELECT * FROM articles WHERE id=?;"
	queryRemoveArticle  = "DELETE FROM articles WHERE id=?;"
	queryUpdateArticle  = "UPDATE articles SET title=?, articleUrl=?, imageUrl=? WHERE id=?;"
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
	if err := result.Scan(&article.ID, &article.Title, &article.ArticleURL, &article.ImageURL); err != nil {
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
		if err := rows.Scan(&article.ID, &article.Title, &article.ArticleURL, &article.ImageURL); err != nil {
			log.Fatal(err)
		}

		articles = append(articles, article)
	}

	return &articles, nil
}

// Save saves an article to the database
func (article *Article) Save() *errors.RestErr {
	stmt, err := articlesdb.Client.Prepare(queryInsertArticles)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	insertResult, err := stmt.Exec(article.Title, article.ArticleURL, article.ImageURL)
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
func Delete(articleID int64) (*int64, *errors.RestErr) {
	stmt, err := articlesdb.Client.Prepare(queryRemoveArticle)
	if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	_, execErr := stmt.Exec(articleID)
	if execErr != nil {
		return nil, errors.NewInternalServerError(fmt.Sprintf("error when trying to delete article: %s", err.Error()))
	}

	return &articleID, nil
}

// Update updates the value of an article in the database
func (article *Article) Update() *errors.RestErr {
	stmt, err := articlesdb.Client.Prepare(queryUpdateArticle)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	_, err = stmt.Exec(article.Title, article.ArticleURL, article.ImageURL, article.ID)
	if err != nil {
		return errors.NewInternalServerError(fmt.Sprintf("error when trying to update article: %s", err.Error()))
	}

	return nil
}
