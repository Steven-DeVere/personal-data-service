package sideprojects

import (
	"fmt"
	"log"

	"github.com/devere-here/personal-data-service/rest-server/datasources/mysql/sideprojectsdb"
	"github.com/devere-here/personal-data-service/rest-server/domain/errors"
)

const (
	queryInsertSideProjects = "INSERT INTO sideProjects(name, blurb, technologies, repoUrl, projectUrl, imageUrl) VALUES(?, ?, ?, ?, ?, ?);"
	queryGetAllSideProjects = "SELECT * FROM sideProjects;"
	queryGetSideProject     = "SELECT * FROM sideProjects WHERE id=?;"
	queryRemoveSideProject  = "DELETE FROM sideProjects WHERE id=?;"
	queryUpdateSideProject  = "UPDATE sideProjects SET name=?, blurb=?, technologies=?, repoUrl=?, projectUrl=?, imageUrl=? WHERE id=?;"
)

// Only point in the application where you interact with the database

var (
	sideProjectsDB = make(map[int64]*SideProject)
)

// Get retrieves a sideProject from the database
func Get(sideProjectID int64) (*SideProject, *errors.RestErr) {
	stmt, err := sideprojectsdb.Client.Prepare(queryGetSideProject)
	if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	sideProject := SideProject{}
	result := stmt.QueryRow(sideProjectID)
	if err := result.Scan(&sideProject.ID, &sideProject.Name, &sideProject.Blurb, &sideProject.Technologies, &sideProject.RepoUrl, &sideProject.ProjectUrl, &sideProject.ImageUrl); err != nil {
		return nil, errors.NewInternalServerError(fmt.Sprintf("error when trying to get side project %d", sideProjectID))
	}

	return &sideProject, nil
}

// GetAll retrieves all sideProjects from the database
func GetAll() (*[]SideProject, *errors.RestErr) {
	stmt, err := sideprojectsdb.Client.Prepare(queryGetAllSideProjects)
	if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	sideProjects := []SideProject{}
	rows, err := stmt.Query()
	if err != nil {
		return nil, errors.NewInternalServerError("error when trying to get side projects")
	}
	defer rows.Close()

	for rows.Next() {
		sideProject := SideProject{}
		if err := rows.Scan(&sideProject.ID, &sideProject.Name, &sideProject.Blurb, &sideProject.Technologies, &sideProject.RepoUrl, &sideProject.ProjectUrl, &sideProject.ImageUrl); err != nil {
			log.Fatal(err)
		}

		sideProjects = append(sideProjects, sideProject)
	}

	return &sideProjects, nil
}

// Save saves a sideProject to the database
func (sideProject *SideProject) Save() *errors.RestErr {
	stmt, err := sideprojectsdb.Client.Prepare(queryInsertSideProjects)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	insertResult, err := stmt.Exec(sideProject.ID, sideProject.Name, sideProject.Blurb, sideProject.Technologies, sideProject.RepoUrl, sideProject.ProjectUrl, sideProject.ImageUrl)
	if err != nil {
		return errors.NewInternalServerError(fmt.Sprintf("error when trying to save side project: %s", err.Error()))
	}

	sideProjectID, err := insertResult.LastInsertId()
	if err != nil {
		return errors.NewInternalServerError(fmt.Sprintf("error when trying to save user: %s", err.Error()))
	}
	sideProject.ID = sideProjectID

	return nil
}

// Delete removes a sideProject from the database
func Delete(sideProjectID int64) (*int64, *errors.RestErr) {
	stmt, err := sideprojectsdb.Client.Prepare(queryRemoveSideProject)
	if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	_, execErr := stmt.Exec(sideProjectID)
	if execErr != nil {
		return nil, errors.NewInternalServerError(fmt.Sprintf("error when trying to delete side project: %s", err.Error()))
	}

	return &sideProjectID, nil
}

// Update updates the value of a side project in the database
func (sideProject *SideProject) Update() *errors.RestErr {
	stmt, err := sideprojectsdb.Client.Prepare(queryUpdateSideProject)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	_, err = stmt.Exec(sideProject.ID, sideProject.Name, sideProject.Blurb, sideProject.Technologies, sideProject.RepoUrl, sideProject.ProjectUrl, sideProject.ImageUrl)
	if err != nil {
		return errors.NewInternalServerError(fmt.Sprintf("error when trying to update side project: %s", err.Error()))
	}

	return nil
}
