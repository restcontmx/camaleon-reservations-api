package data

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/restcontmx/camaleon-reservations-api/app/models"
)

//
// ClientInfoRepository repository object
type ClientInfoRepository struct {
	Model models.ClientInfoModel
	DB    *sql.DB
}

//
// GetAll will return all the objects
// @params none
// @return interface array
// @return error - raise an error if so
//
func (c ClientInfoRepository) GetAll() ([]interface{}, error) {
	var sqlStm = `
				SELECT 	a.id, 
						a.first_name,
						a.last_name,
						a.email,
						a.phone,
						( SELECT CASE WHEN ( a.client_ref <> null ) THEN a.client_ref ELSE 0 END),
						a.location_id,
						a.timestamp, 
						a.updated 
					FROM reservations_client_info a
				`
	var objects []models.ClientInfoModel

	rows, err := c.DB.Query(sqlStm)

	if err != nil {
		return nil, fmt.Errorf("%s", err)
	}

	defer rows.Close()

	for rows.Next() {
		var id int
		var firstname string
		var lastname string
		var email string
		var phone string
		var clientref int
		var locationID int
		var timestamp time.Time
		var updated time.Time

		if err = rows.Scan(&id,
			&firstname,
			&lastname,
			&email,
			&phone,
			&clientref,
			&locationID,
			&timestamp,
			&updated); err != nil {
			return nil, fmt.Errorf("%s", err)
		}

		objects = append(objects, models.ClientInfoModel{
			ID:         id,
			FirstName:  firstname,
			LastName:   lastname,
			Email:      email,
			Phone:      phone,
			LocationID: locationID,
			Timestamp:  timestamp,
			Updated:    updated,
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
func (c ClientInfoRepository) GetByID(id int) (interface{}, error) {
	var sqlStm = `
		SELECT	a.id, 
				a.first_name,
				a.last_name,
				a.email,
				a.phone,
				a.timestamp, 
				a.updated,
				( SELECT CASE WHEN ( a.client_ref <> null ) THEN a.client_ref ELSE 0 END),
				a.location_id
		FROM 	reservations_client_info a
		WHERE a.id = $1`

	var object models.ClientInfoModel

	if err := c.DB.QueryRow(
		sqlStm,
		id,
	).Scan(
		&object.ID,
		&object.FirstName,
		&object.LastName,
		&object.Email,
		&object.Phone,
		&object.Timestamp,
		&object.Updated,
		&object.ClientRef.ID,
		&object.LocationID,
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
func (c ClientInfoRepository) Create() (interface{}, error) {
	var sqlStm = `SELECT create_client_info( $1, $2, $3, $4, $5, $6 )`

	if err := c.DB.QueryRow(
		sqlStm,
		c.Model.FirstName,
		c.Model.LastName,
		c.Model.Email,
		c.Model.Phone,
		c.Model.ClientRef.ID,
		c.Model.LocationID,
	).Scan(
		&c.Model.ID,
	); err != nil {
		return nil, fmt.Errorf("%s", err)
	}

	return c.Model, nil
}

//
// Update will update the object
// @params none
// @return interface array
// @return error - raise an error if so
//
func (c ClientInfoRepository) Update() (interface{}, error) {
	var sqlStm = `SELECT update_client_info( $1, $2, $3, $4, $5, $6, $7 )`

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
		c.Model.FirstName,
		c.Model.LastName,
		c.Model.Email,
		c.Model.Phone,
		c.Model.ClientRef.ID,
		c.Model.LocationID,
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
func (c ClientInfoRepository) Delete() (bool, error) {
	var sqlStm = `DELETE FROM reservations_client_info a WHERE a.id = $1`

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
