package services

import (
	"github.com/devere-here/personal-data-service/rest-server/domain/errors"
	"github.com/devere-here/personal-data-service/rest-server/domain/sideprojects"
)

// GetSideProject creates a side project
func GetSideProject(sideProjectID int64) (*sideprojects.SideProject, *errors.RestErr) {
	sideProject, err := sideprojects.Get(sideProjectID)
	if err != nil {
		return nil, err
	}

	return sideProject, nil
}

// GetAllSideProjects gets all sideProjects
func GetAllSideProjects() (*[]sideprojects.SideProject, *errors.RestErr) {
	sideProjects, err := sideprojects.GetAll()
	if err != nil {
		return nil, err
	}

	return sideProjects, nil
}

// CreateSideProject creates a sideProject
func CreateSideProject(sideProject sideprojects.SideProject) (*sideprojects.SideProject, *errors.RestErr) {
	if err := sideProject.Validate(); err != nil {
		return nil, err
	}

	if err := sideProject.Save(); err != nil {
		return nil, err
	}
	return &sideProject, nil
}

// UpdateSideProject updates a sideProject
func UpdateSideProject(sideProject sideprojects.SideProject) (*sideprojects.SideProject, *errors.RestErr) {
	if err := sideProject.Validate(); err != nil {
		return nil, err
	}

	if err := sideProject.Update(); err != nil {
		return nil, err
	}
	return &sideProject, nil
}

// DeleteSideProject deletes a sideProjecgt
func DeleteSideProject(sideProjectID int64) (*int64, *errors.RestErr) {
	id, err := sideprojects.Delete(sideProjectID)
	if err != nil {
		return nil, err
	}

	return id, nil
}
