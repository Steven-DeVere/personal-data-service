package projects

import (
	"fmt"
	"log"

	"github.com/devere-here/personal-data-service/rest-server/datasources/mysql/projectsdb"
	"github.com/devere-here/personal-data-service/rest-server/domain/errors"
)

const (
	queryInsertProjects = "INSERT INTO projects(name, blurb, repoUrl, projectUrl, imageUrl) VALUES(?, ?, ?, ?, ?);"
	queryGetAllProjects = "SELECT * FROM projects;"
	queryGetProject     = "SELECT * FROM projects WHERE id=?;"
	queryRemoveProject  = "DELETE FROM projects WHERE id=?;"
	queryUpdateProject  = "UPDATE projects SET name=?, blurb=?, repoUrl=?, projectUrl=?, imageUrl=? WHERE id=?;"
)

// Only point in the application where you interact with the database

var (
	projectsDB = make(map[int64]*Project)
)

// Get retrieves a project from the database
func Get(projectID int64) (*Project, *errors.RestErr) {
	stmt, err := projectsdb.Client.Prepare(queryGetProject)
	if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	project := Project{}
	result := stmt.QueryRow(projectID)
	if err := result.Scan(&project.ID, &project.Name, &project.Blurb, &project.RepoURL, &project.ProjectURL, &project.ImageURL); err != nil {
		return nil, errors.NewInternalServerError(fmt.Sprintf("error when trying to get project %d", projectID))
	}

	return &project, nil
}

// GetAll retrieves all projects from the database
func GetAll() (*[]Project, *errors.RestErr) {
	stmt, err := projectsdb.Client.Prepare(queryGetAllProjects)
	if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	projects := []Project{}
	rows, err := stmt.Query()
	if err != nil {
		return nil, errors.NewInternalServerError("error when trying to get projects")
	}
	defer rows.Close()

	for rows.Next() {
		project := Project{}
		if err := rows.Scan(&project.ID, &project.Name, &project.Blurb, &project.RepoURL, &project.ProjectURL, &project.ImageURL); err != nil {
			log.Fatal(err)
		}

		projects = append(projects, project)
	}

	return &projects, nil
}

// Save saves a project to the database
func (project *Project) Save() *errors.RestErr {
	stmt, err := projectsdb.Client.Prepare(queryInsertProjects)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	insertResult, err := stmt.Exec(project.Name, project.Blurb, project.RepoURL, project.ProjectURL, project.ImageURL)
	if err != nil {
		return errors.NewInternalServerError(fmt.Sprintf("error when trying to save project: %s", err.Error()))
	}

	projectID, err := insertResult.LastInsertId()
	if err != nil {
		return errors.NewInternalServerError(fmt.Sprintf("error when trying to save user: %s", err.Error()))
	}
	project.ID = projectID

	return nil
}

// Delete removes a project from the database
func Delete(projectID int64) (*int64, *errors.RestErr) {
	stmt, err := projectsdb.Client.Prepare(queryRemoveProject)
	if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	_, execErr := stmt.Exec(projectID)
	if execErr != nil {
		return nil, errors.NewInternalServerError(fmt.Sprintf("error when trying to delete project: %s", err.Error()))
	}

	return &projectID, nil
}

// Update updates the value of a project in the database
func (project *Project) Update() *errors.RestErr {
	stmt, err := projectsdb.Client.Prepare(queryUpdateProject)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	_, err = stmt.Exec(project.Name, project.Blurb, project.RepoURL, project.ProjectURL, project.ImageURL, project.ID)
	if err != nil {
		return errors.NewInternalServerError(fmt.Sprintf("error when trying to update project: %s", err.Error()))
	}

	return nil
}
