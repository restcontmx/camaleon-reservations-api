package data

import (
	"database/sql"
	"fmt"

	"github.com/restcontmx/camaleon-reservations-api/app/models"
)

// WaitListStatusRepository is the main repository struct
type WaitListStatusRepository struct {
	Model models.WaitListStatusModel
	DB    *sql.DB
}

//
// GetAll will return all the Objects
// @params none
// @return interface array
// @return error - rise an error if so
//
func (w WaitListStatusRepository) GetAll() ([]interface{}, error) {
	var sqlStm = `	SELECT 	a.id, 
							a.description,
							a.value
					FROM reservations_waitlist_status a`

	var objects []models.WaitListStatusModel

	rows, err := w.DB.Query(sqlStm)

	if err != nil {
		return nil, fmt.Errorf("%s", err)
	}

	defer rows.Close()

	for rows.Next() {

		var id int
		var description string
		var value int

		if err = rows.Scan(
			&id,
			&description,
			&value); err != nil {
			return nil, fmt.Errorf("%s", err)
		}

		objects = append(objects, models.WaitListStatusModel{
			ID:          id,
			Description: description,
			Value:       value,
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
// Create will create an object on the db
// @params none
// @return interface array
// @return error - raise an error if so
//
func (w WaitListStatusRepository) Create() (interface{}, error) {
	var sqlStm = `SELECT create_waitlist_status( $1, $2 )`

	tx, err := w.DB.Begin()

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
		w.Model.Description,
		w.Model.Value,
	); err != nil {
		return nil, fmt.Errorf("%s", err)
	}

	if err = tx.Commit(); err != nil {
		return nil, fmt.Errorf("%s", err)
	}

	return w.Model, nil
}

//
// GetByID get object by id
//
func (w WaitListStatusRepository) GetByID(id int) (interface{}, error) {
	var sqlStm = `
		SELECT	a.id,
				a.value,
				a.description
		FROM create_waitlist_status a
		WHERE a.id = $1`

	var object models.WaitListStatusModel

	if err := w.DB.QueryRow(
		sqlStm,
		id,
	).Scan(
		&object.ID,
		&object.Value,
		&object.Description,
	); err != nil {
		return nil, fmt.Errorf("%s", err)
	}

	return object, nil
}

//
// Update function will update a reservation status in the database
// @params none
// @returns interface - an object
//
func (w WaitListStatusRepository) Update() (interface{}, error) {
	var sqlStm = `SELECT update_waitlist_status( $1, $2, $3 )`

	tx, err := w.DB.Begin()

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
		w.Model.ID,
		w.Model.Description,
		w.Model.Value,
	); err != nil {
		return nil, fmt.Errorf("%s", err)
	}

	if err = tx.Commit(); err != nil {
		return nil, fmt.Errorf("%s", err)
	}

	return w.Model, nil
}

//
// Delete function will delete an object in the database
// @params none
// @returns Boolean
//
func (w WaitListStatusRepository) Delete() (bool, error) {
	var sqlStm = `DELETE FROM reservations_waitlist_status a WHERE a.id = $1`

	tx, err := w.DB.Begin()

	if err != nil {
		return false, fmt.Errorf("%s", err)
	}

	defer tx.Rollback()

	stmt, err := tx.Prepare(sqlStm)

	if err != nil {
		return false, fmt.Errorf("%s", err)
	}

	defer stmt.Close()

	if _, err = stmt.Exec(w.Model.ID); err != nil {
		return false, fmt.Errorf("%s", err)
	}

	if err = tx.Commit(); err != nil {
		return false, fmt.Errorf("%s", err)
	}

	return true, nil
}
