package services

import (
	"github.com/devere-here/personal-data-service/domain/articles"
	"github.com/devere-here/personal-data-service/domain/errors"
)

// GetArticle creates an article
func GetArticle(articleID int64) (*articles.Article, *errors.RestErr) {
	article, err := articles.Get(articleID)
	if err != nil {
		return nil, err
	}

	return article, nil
}

// GetAllArticles gets all article
func GetAllArticles() (*[]articles.Article, *errors.RestErr) {
	articles, err := articles.GetAll()
	if err != nil {
		return nil, err
	}

	return articles, nil
}

// CreateArticle creates an article
func CreateArticle(article articles.Article) (*articles.Article, *errors.RestErr) {
	if err := article.Validate(); err != nil {
		return nil, err
	}

	if err := article.Save(); err != nil {
		return nil, err
	}
	return &article, nil
}

// UpdateArticle updates an article
func UpdateArticle(article articles.Article) (*articles.Article, *errors.RestErr) {
	if err := article.Validate(); err != nil {
		return nil, err
	}

	if err := article.Update(); err != nil {
		return nil, err
	}
	return &article, nil
}

// DeleteArticle deletes an article
func DeleteArticle(articleID int64) (*int64, *errors.RestErr) {
	id, err := articles.Delete(articleID)
	if err != nil {
		return nil, err
	}

	return id, nil
}
