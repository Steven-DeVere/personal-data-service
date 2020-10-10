package graphql

import (
	"github.com/devere-here/personal-data-service/domain/articles"
	"github.com/devere-here/personal-data-service/services"
)

type queryResolver struct{ *Resolver }

func (r *queryResolver) Article(id string) (*articles.Article, error) {
	return services.GetArticle(id)
}

func (r *queryResolver) Articles() (*[]articles.Article, error) {
	return services.GetAllArticles()
}
