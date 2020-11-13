package sideprojects

import "github.com/devere-here/personal-data-service/rest-server/domain/errors"

// SideProject defines the sideProject struct
type SideProject struct {
	ID      int64  `json:"id"`
	Name   string `json:"name"`
	Blurb   string `json:"blurb"`
	Technologies   []string `json:"technologies"`
	RepoUrl string `json:"repoUrl"`
	ProjectUrl string `json:"projectUrl"`
	ImageUrl string `json:"imageUrl"`
}

// Validate validates the sideProject
func (sideProject *SideProject) Validate() *errors.RestErr {
	if sideProject.Name == "" {
		return errors.NewBadRequestError("Invalid Name")
	}

	if sideProject.Blurb == "" {
		return errors.NewBadRequestError("Invalid Blurb")
	}

	if sideProject.RepoUrl == "" {
		return errors.NewBadRequestError("Invalid RepoUrl")
	}

	if sideProject.ImageUrl == "" {
		return errors.NewBadRequestError("Invalid ImageUrl")
	}

	return nil
}
