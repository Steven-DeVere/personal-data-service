package controllers

import (
	"net/http"
	"strconv"

	"github.com/devere-here/personal-data-service/rest-server/domain/errors"
	"github.com/devere-here/personal-data-service/rest-server/domain/projects"
	"github.com/devere-here/personal-data-service/rest-server/services"
	"github.com/gin-gonic/gin"
)

// GetProject gets a specific project
func GetProject(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		restErr := errors.NewBadRequestError("invalid id query string")

		c.JSON(restErr.Status, restErr)
		return
	}

	project, dbErr := services.GetProject(id)
	if err != nil {
		c.JSON(dbErr.Status, dbErr)
	}

	c.JSON(http.StatusOK, project)
}

// GetAllProjects gets all projects
func GetAllProjects(c *gin.Context) {
	projects, err := services.GetAllProjects()
	if err != nil {
		c.JSON(err.Status, err)
	}

	c.JSON(http.StatusOK, projects)
}

// DeleteProject deletes an article
func DeleteProject(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		restErr := errors.NewBadRequestError("invalid id query string")

		c.JSON(restErr.Status, restErr)
		return
	}

	projectID, dbErr := services.DeleteProject(id)
	if dbErr != nil {
		c.JSON(dbErr.Status, dbErr)
	}

	c.JSON(http.StatusOK, projectID)
}

// CreateProject creates a new article
func CreateProject(c *gin.Context) {
	var project projects.Project

	if err := c.ShouldBindJSON(&project); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")

		c.JSON(restErr.Status, restErr)
		return
	}

	result, saveErr := services.CreateProject(project)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}

	c.JSON(http.StatusCreated, result)
}

// UpdateProject updates an existing article
func UpdateProject(c *gin.Context) {
	var project projects.Project

	if err := c.ShouldBindJSON(&project); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")

		c.JSON(restErr.Status, restErr)
		return
	}

	if _, err := services.GetProject(project.ID); err != nil {
		c.JSON(http.StatusBadRequest, err)
	}

	result, saveErr := services.UpdateProject(project)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}

	c.JSON(http.StatusOK, result)
}
