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
	url := "http://localhost:3000/projects"
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

	project := model.Project{}
	err = json.Unmarshal(body, &project)

	return &project, err
}

func (r *mutationResolver) UpdateProject(ctx context.Context, input model.UpdateProjectInput) (*model.Project, error) {
	url := "http://localhost:3000/projects"
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

	project := model.Project{}
	err = json.Unmarshal(body, &project)

	return &project, err
}

func (r *mutationResolver) DeleteProject(ctx context.Context, id int) (int, error) {
	url := fmt.Sprintf("http://localhost:3000/projects/%d", id)

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

func (r *mutationResolver) CreateTech(ctx context.Context, input model.CreateTechInput) (*model.Tech, error) {
	url := "http://localhost:3000/tech"
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

	tech := model.Tech{}
	err = json.Unmarshal(body, &tech)

	return &tech, err
}

func (r *mutationResolver) UpdateTech(ctx context.Context, input model.UpdateTechInput) (*model.Tech, error) {
	url := "http://localhost:3000/tech"
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

	tech := model.Tech{}
	err = json.Unmarshal(body, &tech)

	return &tech, err
}

func (r *mutationResolver) DeleteTech(ctx context.Context, id int) (int, error) {
	url := fmt.Sprintf("http://localhost:3000/tech/%d", id)

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
	url := fmt.Sprintf("http://localhost:3000/projects/%d", id)
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

	project := model.Project{}
	err = json.Unmarshal(body, &project)

	return &project, err
}

func (r *queryResolver) Projects(ctx context.Context) ([]*model.Project, error) {
	url := "http://localhost:3000/projects"
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

	projects := []*model.Project{}
	err = json.Unmarshal(body, &projects)

	return projects, err
}

func (r *queryResolver) Tech(ctx context.Context, id int) (*model.Tech, error) {
	url := fmt.Sprintf("http://localhost:3000/tech/%d", id)
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

	tech := model.Tech{}
	err = json.Unmarshal(body, &tech)

	return &tech, err
}

func (r *queryResolver) Alltech(ctx context.Context) ([]*model.Tech, error) {
	url := "http://localhost:3000/tech"
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

	tech := []*model.Tech{}
	err = json.Unmarshal(body, &tech)

	return tech, err
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
