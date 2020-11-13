// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Article struct {
	ID      *int    `json:"id"`
	Title   *string `json:"title"`
	Blurb   *string `json:"blurb"`
	Content *string `json:"Content"`
}

type CreateArticleInput struct {
	Title   string  `json:"title"`
	Blurb   *string `json:"blurb"`
	Content string  `json:"Content"`
}

type UpdateArticleInput struct {
	ID      *int    `json:"id"`
	Title   string  `json:"title"`
	Blurb   *string `json:"blurb"`
	Content string  `json:"Content"`
}