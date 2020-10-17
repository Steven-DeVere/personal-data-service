package model

import (
	"github.com/devere-here/personal-data-service/domain/articles"
)

// ConvertCreateInput converts CreateArticleInput and convert it to an Article
func ConvertCreateInput(articleInputs CreateArticleInput) articles.Article {
	return articles.Article{
		Title:   articleInputs.Title,
		Blurb:   *articleInputs.Blurb,
		Content: articleInputs.Content,
	}
}

// ConvertUpdateInput converts CreateArticleInput and convert it to an Article
func ConvertUpdateInput(articleInputs UpdateArticleInput) articles.Article {
	return articles.Article{
		ID:      int64(*articleInputs.ID),
		Title:   articleInputs.Title,
		Blurb:   *articleInputs.Blurb,
		Content: articleInputs.Content,
	}
}
