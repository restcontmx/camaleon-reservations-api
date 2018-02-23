package data

import (
	"database/sql"
	"fmt"

	"github.com/restcontmx/camaleon-reservations-api/app/models"
)

// ReservationStatusRepository repository object
type ReservationStatusRepository struct {
	Model models.ReservationStatusModel
	DB    *sql.DB
}

//
// GetAll will return all the Objects
// @params none
// @return interface array
// @return error - rise an error if so
//
func (r ReservationStatusRepository) GetAll() ([]interface{}, error) {
	var sqlStm = `	SELECT 	a.id, 
							a.description,
							a.value
					FROM reservations_reservation_status a`

	var objects []models.ReservationStatusModel

	rows, err := r.DB.Query(sqlStm)

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

		objects = append(objects, models.ReservationStatusModel{
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
func (r ReservationStatusRepository) Create() (interface{}, error) {
	var sqlStm = `SELECT create_reservation_status( $1, $2 )`

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
// GetByID get object by id
//
func (r ReservationStatusRepository) GetByID(id int) (interface{}, error) {
	var sqlStm = `
		SELECT	a.id,
				a.value,
				a.description
		FROM reservations_reservation_status a
		WHERE a.id = $1`

	var reservation models.ReservationStatusModel

	if err := r.DB.QueryRow(
		sqlStm,
		id,
	).Scan(
		&reservation.ID,
		&reservation.Value,
		&reservation.Description,
	); err != nil {
		return nil, fmt.Errorf("%s", err)
	}

	return reservation, nil
}

//
// Update function will update a reservation status in the database
// @params none
// @returns interface - an object
//
func (r ReservationStatusRepository) Update() (interface{}, error) {
	var sqlStm = `SELECT update_reservation_status( $1, $2, $3 )`

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
