package data

import (
	"database/sql"
	"fmt"

	"github.com/restcontmx/camaleon-reservations-api/app/models"
)

// RolRepository main rol repo struct
type RolRepository struct {
	Model models.RolModel
	DB    *sql.DB
}

//
// GetAll will return all the rols
// @params none
// @return interface array
// @return error - rise an error if so
//
func (r RolRepository) GetAll() ([]interface{}, error) {
	var sqlStm = `
				SELECT 	a.id,
						a.description,
						a.value
				FROM reservations_rol a`

	var objects []models.RolModel

	rows, err := r.DB.Query(sqlStm)

	if err != nil {
		return nil, fmt.Errorf("%s", err)
	}

	defer rows.Close()

	for rows.Next() {
		var id int
		var value int
		var description string

		if err = rows.Scan(
			&id,
			&description,
			&value); err != nil {
			return nil, fmt.Errorf("%s", err)
		}

		objects = append(objects, models.RolModel{
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
func (r RolRepository) Create() (interface{}, error) {
	var sqlStm = `SELECT create_rol( $1, $2 )`

	tx, err := r.DB.Begin()

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
		r.Model.Description,
		r.Model.Value,
	); err != nil {
		return nil, fmt.Errorf("%s", err)
	}

	if err = tx.Commit(); err != nil {
		return nil, fmt.Errorf("%s", err)
	}

	return r.Model, nil
}

//
// Update function will update an object in the database
// @params none
// @returns interface - a rol
// @returns error - in case of one
//
func (r RolRepository) Update() (interface{}, error) {

	var sqlStm = `SELECT update_rol( $1, $2, $3 )`

	tx, err := r.DB.Begin()

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
		r.Model.ID,
		r.Model.Description,
		r.Model.Value,
	); err != nil {
		return nil, fmt.Errorf("%s", err)
	}

	if err = tx.Commit(); err != nil {
		return nil, fmt.Errorf("%s", err)
	}

	return r.Model, nil
}

//
// GetByID get object by id
//
func (r RolRepository) GetByID(id int) (interface{}, error) {
	var sqlStm = `
			SELECT 	a.id,
					a.description,
					a.value
			FROM reservations_rol a
			WHERE a.id = $1`

	var object models.RolModel

	if err := r.DB.QueryRow(
		sqlStm,
		id,
	).Scan(
		&object.ID,
		&object.Description,
		&object.Value,
	); err != nil {
		return nil, fmt.Errorf("%s", err)
	}

	return object, nil
}
