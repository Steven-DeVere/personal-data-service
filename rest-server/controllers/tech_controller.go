package controllers

import (
	"net/http"
	"strconv"

	"github.com/devere-here/personal-data-service/rest-server/domain/errors"
	"github.com/devere-here/personal-data-service/rest-server/domain/tech"
	"github.com/devere-here/personal-data-service/rest-server/services"
	"github.com/gin-gonic/gin"
)

// GetTech gets a specific tech
func GetTech(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		restErr := errors.NewBadRequestError("invalid id query string")

		c.JSON(restErr.Status, restErr)
		return
	}

	tech, dbErr := services.GetTech(id)
	if err != nil {
		c.JSON(dbErr.Status, dbErr)
	}

	c.JSON(http.StatusOK, tech)
}

// GetAllTech gets all tech
func GetAllTech(c *gin.Context) {
	allTech, err := services.GetAllTech()
	if err != nil {
		c.JSON(err.Status, err)
	}

	c.JSON(http.StatusOK, allTech)
}

// DeleteTech deletes an article
func DeleteTech(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		restErr := errors.NewBadRequestError("invalid id query string")

		c.JSON(restErr.Status, restErr)
		return
	}

	techID, dbErr := services.DeleteTech(id)
	if dbErr != nil {
		c.JSON(dbErr.Status, dbErr)
	}

	c.JSON(http.StatusOK, techID)
}

// CreateTech creates a new article
func CreateTech(c *gin.Context) {
	var tech tech.Tech

	if err := c.ShouldBindJSON(&tech); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")

		c.JSON(restErr.Status, restErr)
		return
	}

	result, saveErr := services.CreateTech(tech)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}

	c.JSON(http.StatusCreated, result)
}

// UpdateTech updates an existing article
func UpdateTech(c *gin.Context) {
	var tech tech.Tech

	if err := c.ShouldBindJSON(&tech); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")

		c.JSON(restErr.Status, restErr)
		return
	}

	if _, err := services.GetTech(tech.ID); err != nil {
		c.JSON(http.StatusBadRequest, err)
	}

	result, saveErr := services.UpdateTech(tech)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}

	c.JSON(http.StatusOK, result)
}
