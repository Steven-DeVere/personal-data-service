package services

import (
	"github.com/devere-here/personal-data-service/rest-server/domain/errors"
	"github.com/devere-here/personal-data-service/rest-server/domain/projects"
)

// GetProject creates a project
func GetProject(projectID int64) (*projects.Project, *errors.RestErr) {
	project, err := projects.Get(projectID)
	if err != nil {
		return nil, err
	}

	return project, nil
}

// GetAllProjects gets all projects
func GetAllProjects() (*[]projects.Project, *errors.RestErr) {
	projects, err := projects.GetAll()
	if err != nil {
		return nil, err
	}

	return projects, nil
}

// CreateProject creates a project
func CreateProject(project projects.Project) (*projects.Project, *errors.RestErr) {
	if err := project.Validate(); err != nil {
		return nil, err
	}

	if err := project.Save(); err != nil {
		return nil, err
	}
	return &project, nil
}

// UpdateProject updates a project
func UpdateProject(project projects.Project) (*projects.Project, *errors.RestErr) {
	if err := project.Validate(); err != nil {
		return nil, err
	}

	if err := project.Update(); err != nil {
		return nil, err
	}
	return &project, nil
}

// DeleteProject deletes a project
func DeleteProject(projectID int64) (*int64, *errors.RestErr) {
	id, err := projects.Delete(projectID)
	if err != nil {
		return nil, err
	}

	return id, nil
}
