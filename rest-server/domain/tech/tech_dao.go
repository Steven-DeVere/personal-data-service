package tech

import (
	"fmt"
	"log"

	"github.com/devere-here/personal-data-service/rest-server/datasources/mysql/techdb"
	"github.com/devere-here/personal-data-service/rest-server/domain/errors"
)

const (
	queryInsertTech = "INSERT INTO tech(name, type, proficiency) VALUES(?, ?, ?);"
	queryGetAllTech = "SELECT * FROM tech;"
	queryGetTech    = "SELECT * FROM tech WHERE id=?;"
	queryRemoveTech = "DELETE FROM tech WHERE id=?;"
	queryUpdateTech = "UPDATE tech SET name=?, type=?, proficiency=? WHERE id=?;"
)

// Only point in the application where you interact with the database

var (
	techDB = make(map[int64]*Tech)
)

// Get retrieves a tech item from the database
func Get(techID int64) (*Tech, *errors.RestErr) {
	stmt, err := techdb.Client.Prepare(queryGetTech)
	if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	tech := Tech{}
	result := stmt.QueryRow(techID)
	if err := result.Scan(&tech.ID, &tech.Name, &tech.Type, &tech.Proficiency); err != nil {
		return nil, errors.NewInternalServerError(fmt.Sprintf("error when trying to get tech %d", techID))
	}

	return &tech, nil
}

// GetAll retrieves all tech from the database
func GetAll() (*[]Tech, *errors.RestErr) {
	stmt, err := techdb.Client.Prepare(queryGetAllTech)
	if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	tech := []Tech{}
	rows, err := stmt.Query()
	if err != nil {
		return nil, errors.NewInternalServerError("error when trying to get tech")
	}
	defer rows.Close()

	for rows.Next() {
		singleTech := Tech{}
		if err := rows.Scan(&singleTech.ID, &singleTech.Name, &singleTech.Type, &singleTech.Proficiency); err != nil {
			log.Fatal(err)
		}

		tech = append(tech, singleTech)
	}

	return &tech, nil
}

// Save saves a tech to the database
func (tech *Tech) Save() *errors.RestErr {
	stmt, err := techdb.Client.Prepare(queryInsertTech)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	insertResult, err := stmt.Exec(tech.Name, tech.Type, tech.Proficiency)
	if err != nil {
		return errors.NewInternalServerError(fmt.Sprintf("error when trying to save tech: %s", err.Error()))
	}

	techID, err := insertResult.LastInsertId()
	if err != nil {
		return errors.NewInternalServerError(fmt.Sprintf("error when trying to save tech: %s", err.Error()))
	}
	tech.ID = techID

	return nil
}

// Delete removes a tech from the database
func Delete(techID int64) (*int64, *errors.RestErr) {
	stmt, err := techdb.Client.Prepare(queryRemoveTech)
	if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	_, execErr := stmt.Exec(techID)
	if execErr != nil {
		return nil, errors.NewInternalServerError(fmt.Sprintf("error when trying to delete tech: %s", err.Error()))
	}

	return &techID, nil
}

// Update updates the value of a tech in the database
func (tech *Tech) Update() *errors.RestErr {
	stmt, err := techdb.Client.Prepare(queryUpdateTech)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	_, err = stmt.Exec(tech.Name, tech.Type, tech.Proficiency, tech.ID)
	if err != nil {
		return errors.NewInternalServerError(fmt.Sprintf("error when trying to update tech: %s", err.Error()))
	}

	return nil
}
