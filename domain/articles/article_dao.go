package articles

import (
	"fmt"

	"github.com/devere-here/personal-data-service/domain/errors"
)

// Only point in the application where you interact with the database

var (
	articlesDB = make(map[int64]*Article)
)

// Get retrieves an article from the database
func (article *Article) Get() *errors.RestErr {
	result := articlesDB[article.ID]
	if result == nil {
		return errors.NewNotFoundError(fmt.Sprintf("article %d not found", article.ID))
	}
	return nil
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
