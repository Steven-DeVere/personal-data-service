package services

import (
	"github.com/devere-here/personal-data-service/rest-server/domain/errors"
	"github.com/devere-here/personal-data-service/rest-server/domain/tech"
)

// GetTech creates a project
func GetTech(techID int64) (*tech.Tech, *errors.RestErr) {
	project, err := tech.Get(techID)
	if err != nil {
		return nil, err
	}

	return project, nil
}

// GetAllTech gets all tech
func GetAllTech() (*[]tech.Tech, *errors.RestErr) {
	tech, err := tech.GetAll()
	if err != nil {
		return nil, err
	}

	return tech, nil
}

// CreateTech creates a project
func CreateTech(project tech.Tech) (*tech.Tech, *errors.RestErr) {
	if err := project.Validate(); err != nil {
		return nil, err
	}

	if err := project.Save(); err != nil {
		return nil, err
	}
	return &project, nil
}

// UpdateTech updates a project
func UpdateTech(project tech.Tech) (*tech.Tech, *errors.RestErr) {
	if err := project.Validate(); err != nil {
		return nil, err
	}

	if err := project.Update(); err != nil {
		return nil, err
	}
	return &project, nil
}

// DeleteTech deletes a project
func DeleteTech(techID int64) (*int64, *errors.RestErr) {
	id, err := tech.Delete(techID)
	if err != nil {
		return nil, err
	}

	return id, nil
}
