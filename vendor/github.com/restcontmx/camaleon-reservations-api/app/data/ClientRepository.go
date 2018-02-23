package data

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/restcontmx/camaleon-reservations-api/app/models"
)

//
// ClientRepository object
//
type ClientRepository struct {
	Model models.ClientModel
	DB    *sql.DB
}

//
// GetAll will return all the objects
// @params none
// @return interface array
// @return error - raise an error if so
//
func (c ClientRepository) GetAll() ([]interface{}, error) {
	var sqlStm = `
				SELECT 	id, 
						auth0_id,
						timestamp, 
						updated 
						FROM reservations_client
				`
	var objects []models.ClientModel

	rows, err := c.DB.Query(sqlStm)

	if err != nil {
		return nil, fmt.Errorf("%s", err)
	}

	defer rows.Close()

	for rows.Next() {
		var id int
		var auth0id string
		var timestamp time.Time
		var updated time.Time

		if err = rows.Scan(&id, &auth0id, &timestamp, &updated); err != nil {
			return nil, fmt.Errorf("%s", err)
		}

		objects = append(objects, models.ClientModel{
			ID:        id,
			Auth0ID:   auth0id,
			Timestamp: timestamp,
			Updated:   updated,
		})
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("%s", err)
	}

	intfObjects := make([]interface{}, len(objects))

	for i, obj := range objects {
		intfObjects[i] = obj
	}

	return intfObjects, nil
}

//
// GetByID returns an object by id
// @params none
// @return interface array
// @return error - raise an error if so
//
func (c ClientRepository) GetByID(id int) (interface{}, error) {
	var sqlStm = `
		SELECT	a.id, 
				a.auth0_id,
				a.timestamp, 
				a.updated 
		FROM reservations_client a
		WHERE a.id = $1`

	var object models.ClientModel

	if err := c.DB.QueryRow(
		sqlStm,
		id,
	).Scan(
		&object.ID,
		&object.Auth0ID,
		&object.Timestamp,
		&object.Updated,
	); err != nil {
		return nil, fmt.Errorf("%s", err)
	}

	return object, nil
}

//
// Create will create an object on the db
// @params none
// @return interface array
// @return error - raise an error if so
//
func (c ClientRepository) Create() (interface{}, error) {
	var sqlStm = `SELECT create_client( $1 )`

	tx, err := c.DB.Begin()

	if err != nil {
		return nil, fmt.Errorf("%s", err)
	}

	defer tx.Rollback()

	stmt, err := tx.Prepare(sqlStm)

	if err != nil {
		return nil, fmt.Errorf("%s", err)
	}

	defer stmt.Close()

	if err != nil {
		return nil, fmt.Errorf("%s", err)
	}

	if _, err = stmt.Exec(
		c.Model.Auth0ID,
	); err != nil {
		return nil, fmt.Errorf("%s", err)
	}

	if err = tx.Commit(); err != nil {
		return nil, fmt.Errorf("%s", err)
	}

	return nil, nil
}

//
// Update will update the object
// @params none
// @return interface array
// @return error - raise an error if so
//
func (c ClientRepository) Update() (interface{}, error) {
	var sqlStm = `SELECT update_client( $1, $2 )`

	tx, err := c.DB.Begin()

	if err != nil {
		return nil, fmt.Errorf("%s", err)
	}

	defer tx.Rollback()

	stmt, err := tx.Prepare(sqlStm)

	if err != nil {
		return nil, fmt.Errorf("%s", err)
	}

	defer stmt.Close()

	if _, err = stmt.Exec(
		c.Model.ID,
		c.Model.Auth0ID,
	); err != nil {
		return nil, fmt.Errorf("%s", err)
	}

	if err = tx.Commit(); err != nil {
		return nil, fmt.Errorf("%s", err)
	}

	return c.Model, nil
}

//
// Delete function will delete an object in the database
// @params none
// @returns Boolean
//
func (c ClientRepository) Delete() (bool, error) {
	var sqlStm = `DELETE FROM reservations_client a WHERE a.id = $1`

	tx, err := c.DB.Begin()

	if err != nil {
		return false, fmt.Errorf("%s", err)
	}

	defer tx.Rollback()

	stmt, err := tx.Prepare(sqlStm)

	if err != nil {
		return false, fmt.Errorf("%s", err)
	}

	defer stmt.Close()

	if _, err = stmt.Exec(c.Model.ID); err != nil {
		return false, fmt.Errorf("%s", err)
	}

	if err = tx.Commit(); err != nil {
		return false, fmt.Errorf("%s", err)
	}

	return true, nil
}
