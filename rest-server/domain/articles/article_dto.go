package articles

import "github.com/devere-here/personal-data-service/rest-server/domain/errors"

// Article defines the article struct
type Article struct {
	ID         int64  `json:"id"`
	Title      string `json:"title"`
	ArticleURL string `json:"articleUrl"`
	ImageURL   string `json:"imageUrl"`
}

// Validate validates the article
func (article *Article) Validate() *errors.RestErr {
	if article.ArticleURL == "" {
		return errors.NewBadRequestError("Invalid ArticleURL")
	}

	if article.ImageURL == "" {
		return errors.NewBadRequestError("Invalid ImageURL")
	}

	if article.Title == "" {
		return errors.NewBadRequestError("Invalid Title")
	}

	return nil
}
