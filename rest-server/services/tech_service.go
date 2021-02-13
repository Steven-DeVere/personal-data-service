package services

import (
	"github.com/devere-here/personal-data-service/rest-server/domain/errors"
	"github.com/devere-here/personal-data-service/rest-server/domain/tech"
)

// GetTech creates a tech
func GetTech(techID int64) (*tech.Tech, *errors.RestErr) {
	tech, err := tech.Get(techID)
	if err != nil {
		return nil, err
	}

	return tech, nil
}

// GetAllTech gets all tech
func GetAllTech() (*[]tech.Tech, *errors.RestErr) {
	tech, err := tech.GetAll()
	if err != nil {
		return nil, err
	}

	return tech, nil
}

// CreateTech creates a tech
func CreateTech(tech tech.Tech) (*tech.Tech, *errors.RestErr) {
	if err := tech.Validate(); err != nil {
		return nil, err
	}
	if err := tech.Save(); err != nil {
		return nil, err
	}

	return &tech, nil
}

// UpdateTech updates a tech
func UpdateTech(tech tech.Tech) (*tech.Tech, *errors.RestErr) {
	if err := tech.Validate(); err != nil {
		return nil, err
	}

	if err := tech.Update(); err != nil {
		return nil, err
	}
	return &tech, nil
}

// DeleteTech deletes a tech
func DeleteTech(techID int64) (*int64, *errors.RestErr) {
	id, err := tech.Delete(techID)
	if err != nil {
		return nil, err
	}

	return id, nil
}
