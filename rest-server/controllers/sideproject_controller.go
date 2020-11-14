package controllers

import (
	"net/http"
	"strconv"

	"github.com/devere-here/personal-data-service/rest-server/domain/errors"
	"github.com/devere-here/personal-data-service/rest-server/domain/sideprojects"
	"github.com/devere-here/personal-data-service/rest-server/services"
	"github.com/gin-gonic/gin"
)

// GetSideProject gets a specific project
func GetSideProject(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		restErr := errors.NewBadRequestError("invalid id query string")

		c.JSON(restErr.Status, restErr)
		return
	}

	project, dbErr := services.GetSideProject(id)
	if err != nil {
		c.JSON(dbErr.Status, dbErr)
	}

	c.JSON(http.StatusOK, project)
}

// GetAllSideProjects gets all projects
func GetAllSideProjects(c *gin.Context) {
	projects, err := services.GetAllSideProjects()
	if err != nil {
		c.JSON(err.Status, err)
	}

	c.JSON(http.StatusOK, projects)
}

// DeleteSideProject deletes an article
func DeleteSideProject(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		restErr := errors.NewBadRequestError("invalid id query string")

		c.JSON(restErr.Status, restErr)
		return
	}

	sideProjectID, dbErr := services.DeleteSideProject(id)
	if dbErr != nil {
		c.JSON(dbErr.Status, dbErr)
	}

	c.JSON(http.StatusOK, sideProjectID)
}

// CreateSideProject creates a new article
func CreateSideProject(c *gin.Context) {
	var project sideprojects.SideProject

	if err := c.ShouldBindJSON(&project); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")

		c.JSON(restErr.Status, restErr)
		return
	}

	result, saveErr := services.CreateSideProject(project)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}

	c.JSON(http.StatusCreated, result)
}

// UpdateSideProject updates an existing article
func UpdateSideProject(c *gin.Context) {
	var project sideprojects.SideProject

	if err := c.ShouldBindJSON(&project); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")

		c.JSON(restErr.Status, restErr)
		return
	}

	if _, err := services.GetSideProject(project.ID); err != nil {
		c.JSON(http.StatusBadRequest, err)
	}

	result, saveErr := services.UpdateSideProject(project)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}

	c.JSON(http.StatusOK, result)
}
