package projects

import "github.com/devere-here/personal-data-service/rest-server/domain/errors"

// Project defines the project struct
type Project struct {
	ID         int64  `json:"id"`
	Name       string `json:"name"`
	Blurb      string `json:"blurb"`
	RepoURL    string `json:"repoUrl"`
	ProjectURL string `json:"projectUrl"`
	ImageURL   string `json:"imageUrl"`
}

// Validate validates the project
func (project *Project) Validate() *errors.RestErr {
	if project.Name == "" {
		return errors.NewBadRequestError("Invalid Name")
	}

	if project.Blurb == "" {
		return errors.NewBadRequestError("Invalid Blurb")
	}

	if project.RepoURL == "" {
		return errors.NewBadRequestError("Invalid RepoURL")
	}

	if project.ImageURL == "" {
		return errors.NewBadRequestError("Invalid ImageURL")
	}

	return nil
}
