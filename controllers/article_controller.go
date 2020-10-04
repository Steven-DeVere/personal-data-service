package controllers

import (
	"net/http"

	"github.com/devere-here/personal-data-service/domain/articles"
	"github.com/devere-here/personal-data-service/domain/errors"
	"github.com/devere-here/personal-data-service/services"
	"github.com/gin-gonic/gin"
)

// GetArticle gets a specific article
func GetArticle(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me!")
}

// GetAllArticles gets all articles
func GetAllArticles(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me!")
}

// SearchArticle searches for an article
func SearchArticle(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me!")
}

// CreateArticle creates a new article
func CreateArticle(c *gin.Context) {
	var article articles.Article

	if err := c.ShouldBindJSON(&article); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")

		c.JSON(restErr.Status, restErr)
		return
	}

	result, saveErr := services.CreateArticle(article)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}

	c.JSON(http.StatusCreated, result)
}
