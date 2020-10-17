package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/devere-here/personal-data-service/domain/articles"
	"github.com/devere-here/personal-data-service/graph/generated"
	"github.com/devere-here/personal-data-service/graph/model"
)

func (r *mutationResolver) CreateArticle(ctx context.Context, input model.CreateArticleInput) (*articles.Article, error) {
	url := "/articles"
	jsonInput, err := json.Marshal(input)
	if err != nil {
		log.Fatal(err)
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(jsonInput))
	if err != nil {
		log.Fatal(err)
	}

	res, err := r.Client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	article := articles.Article{}
	err = json.Unmarshal(body, &article)

	return &article, err
}

func (r *mutationResolver) UpdateArticle(ctx context.Context, input model.UpdateArticleInput) (*articles.Article, error) {
	url := "/articles"
	jsonInput, err := json.Marshal(input)
	if err != nil {
		log.Fatal(err)
	}

	req, err := http.NewRequest(http.MethodPut, url, bytes.NewReader(jsonInput))
	if err != nil {
		log.Fatal(err)
	}

	res, err := r.Client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	article := articles.Article{}
	err = json.Unmarshal(body, &article)

	return &article, err
}

func (r *mutationResolver) DeleteArticle(ctx context.Context, id int) (*int64, error) {
	url := fmt.Sprintf("/articles/%s", id)

	req, err := http.NewRequest(http.MethodPut, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	res, err := r.Client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	var i64 int64
	err = json.Unmarshal(body, &i64)

	return &i64, err
}

func (r *queryResolver) Article(ctx context.Context, id int) (*articles.Article, error) {
	url := fmt.Sprintf("/articles/%s", id)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	res, err := r.Client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	article := articles.Article{}
	err = json.Unmarshal(body, &article)

	return &article, err
}

func (r *queryResolver) Articles(ctx context.Context) (*[]articles.Article, error) {
	url := "/articles"
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	res, err := r.Client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	articles := []articles.Article{}
	err = json.Unmarshal(body, &articles)

	return &articles, err
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
