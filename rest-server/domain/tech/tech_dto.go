package tech

import "github.com/devere-here/personal-data-service/rest-server/domain/errors"

// Tech defines the tech struct
type Tech struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Type        string `json:"type"`
	Proficiency string `json:"proficiency"`
}

// Validate validates the tech item
func (tech *Tech) Validate() *errors.RestErr {
	if tech.Name == "" {
		return errors.NewBadRequestError("Invalid Name")
	}

	if tech.Type == "" {
		return errors.NewBadRequestError("Invalid Type")
	}

	if tech.Proficiency == "" {
		return errors.NewBadRequestError("Invalid Proficiency")
	}

	return nil
}
