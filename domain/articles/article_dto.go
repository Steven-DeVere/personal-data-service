package articles

import "github.com/devere-here/personal-data-service/domain/errors"

// Article defines the article struct
type Article struct {
	ID         int64    `json:"id"`
	Title      string   `json:"title"`
	Blurb      string   `json:"blurb"`
	Content    string   `json:"content"`
	Categories []string `json:"categories"`
}

// Validate validates the article
func (article *Article) Validate() *errors.RestErr {
	if article.Content == "" {
		return errors.NewBadRequestError("Invalid Content")
	}

	if article.Title == "" {
		return errors.NewBadRequestError("Invalid Title")
	}

	return nil
}
