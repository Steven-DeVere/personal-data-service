package controllers

import (
	"net/http"
	"strconv"

	"github.com/devere-here/personal-data-service/domain/articles"
	"github.com/devere-here/personal-data-service/domain/errors"
	"github.com/devere-here/personal-data-service/services"
	"github.com/gin-gonic/gin"
)

// GetArticle gets a specific article
func GetArticle(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		restErr := errors.NewBadRequestError("invalid id query string")

		c.JSON(restErr.Status, restErr)
		return
	}

	article, dbErr := articles.Get(id)
	if dbErr != nil {
		c.JSON(dbErr.Status, dbErr)
	}

	c.JSON(http.StatusOK, article)
}

// GetAllArticles gets all articles
func GetAllArticles(c *gin.Context) {
	articles, dbErr := articles.GetAll()
	if dbErr != nil {
		c.JSON(dbErr.Status, dbErr)
	}

	c.JSON(http.StatusOK, articles)
}

// DeleteArticle deletes an article
func DeleteArticle(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		restErr := errors.NewBadRequestError("invalid id query string")

		c.JSON(restErr.Status, restErr)
		return
	}

	articleID, dbErr := articles.Delete(id)
	if dbErr != nil {
		c.JSON(dbErr.Status, dbErr)
	}

	c.JSON(http.StatusOK, articleID)
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
