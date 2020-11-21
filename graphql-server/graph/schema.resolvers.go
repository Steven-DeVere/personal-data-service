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

	"github.com/devere-here/personal-data-service/graphql-server/graph/generated"
	"github.com/devere-here/personal-data-service/graphql-server/graph/model"
)

func (r *mutationResolver) CreateArticle(ctx context.Context, input model.CreateArticleInput) (*model.Article, error) {
	url := "http://localhost:3000/articles"
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

	article := model.Article{}
	err = json.Unmarshal(body, &article)

	return &article, err
}

func (r *mutationResolver) UpdateArticle(ctx context.Context, input model.UpdateArticleInput) (*model.Article, error) {
	url := "http://localhost:3000/articles"
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

	article := model.Article{}
	err = json.Unmarshal(body, &article)

	return &article, err
}

func (r *mutationResolver) DeleteArticle(ctx context.Context, id int) (int, error) {
	url := fmt.Sprintf("http://localhost:3000/articles/%d", id)

	req, err := http.NewRequest(http.MethodDelete, url, nil)
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

	return int(i64), err
}

func (r *mutationResolver) CreateProject(ctx context.Context, input model.CreateProjectInput) (*model.Project, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateProject(ctx context.Context, input model.UpdateProjectInput) (*model.Project, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteProject(ctx context.Context, id int) (int, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Article(ctx context.Context, id int) (*model.Article, error) {
	url := fmt.Sprintf("http://localhost:3000/articles/%d", id)
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

	article := model.Article{}
	err = json.Unmarshal(body, &article)

	return &article, err
}

func (r *queryResolver) Articles(ctx context.Context) ([]*model.Article, error) {
	url := "http://localhost:3000/articles"
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

	articles := []*model.Article{}
	err = json.Unmarshal(body, &articles)

	return articles, err
}

func (r *queryResolver) Project(ctx context.Context, id int) (*model.Project, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Projects(ctx context.Context) ([]*model.Project, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
