package services

import (
	"github.com/devere-here/personal-data-service/domain/articles"
	"github.com/devere-here/personal-data-service/domain/errors"
)

// CreateArticle creates an article
func CreateArticle(article articles.Article) (*articles.Article, *errors.RestErr) {
	return &article, nil
}
