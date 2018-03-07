package data

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/restcontmx/camaleon-reservations-api/app/models"
)

// WaitListRepository is the main repository struct
type WaitListRepository struct {
	Model models.WaitListModel
	DB    *sql.DB
}

//
// GetAll will return all the users
// @params none
// @return interface array
// @return error - rise an error if so
//
func (w WaitListRepository) GetAll() ([]interface{}, error) {
	var sqlStm = `
				SELECT 	a.id,
						a.alert_time,
						a.time_limit,
						a.guests,
						a.timestamp, 
						a.updated,
						a.client_info_id,
						b.first_name,
						b.last_name,
						b.email,
						b.phone,
						a.status,
						c.description,
						c.value,
						a.location_id,
						e.name,
						e.business_id
				FROM 			reservations_waitlist a
					INNER JOIN 	reservations_client_info b ON a.client_info_id = b.id
					INNER JOIN 	reservations_waitlist_status c ON a.status = c.id
					INNER JOIN 	reservations_location e ON a.location_id = e.id`

	var objects []models.WaitListModel

	rows, err := w.DB.Query(sqlStm)

	if err != nil {
		return nil, fmt.Errorf("%s", err)
	}

	defer rows.Close()

	for rows.Next() {
		object := models.WaitListModel{
			ClientInfo: models.ClientInfoModel{},
			Location:   models.LocationModel{},
			Status:     models.WaitListStatusModel{},
		}

		if err = rows.Scan(
			&object.ID,
			&object.AlertTime,
			&object.TimeLimit,
			&object.Guests,
			&object.Timestamp,
			&object.Updated,
			&object.ClientInfo.ID,
			&object.ClientInfo.FirstName,
			&object.ClientInfo.LastName,
			&object.ClientInfo.Email,
			&object.ClientInfo.Phone,
			&object.Status.ID,
			&object.Status.Description,
			&object.Status.Value,
			&object.Location.ID,
			&object.Location.Name,
			&object.Location.BusinessID); err != nil {
			return nil, fmt.Errorf("%s", err)
		}

		objects = append(objects, object)
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
// GetByID Returns a user by id field
// @param id - int
// @return interface
//
func (w WaitListRepository) GetByID(id int) (interface{}, error) {
	var sqlStm = `
				SELECT 	a.id,
						a.alert_time,
						a.time_limit,
						a.guests,
						a.timestamp, 
						a.updated,
						a.client_info_id,
						b.first_name,
						b.last_name,
						b.email,
						b.phone,
						a.status,
						c.description,
						c.value,
						a.location_id,
						e.name,
						e.business_id
					FROM 		reservations_waitlist a
					INNER JOIN 	reservations_client_info b ON a.client_info_id = b.id
					INNER JOIN 	reservations_waitlist_status c ON a.status = c.id
					INNER JOIN 	reservations_location e ON a.location_id = e.id
				WHERE a.id = $1`

	object := models.WaitListModel{
		ClientInfo: models.ClientInfoModel{},
		Location:   models.LocationModel{},
		Status:     models.WaitListStatusModel{},
	}

	if err := w.DB.QueryRow(
		sqlStm,
		id,
	).Scan(
		&object.ID,
		&object.AlertTime,
		&object.TimeLimit,
		&object.Guests,
		&object.Timestamp,
		&object.Updated,
		&object.ClientInfo.ID,
		&object.ClientInfo.FirstName,
		&object.ClientInfo.LastName,
		&object.ClientInfo.Email,
		&object.ClientInfo.Phone,
		&object.Status.ID,
		&object.Status.Description,
		&object.Status.Value,
		&object.Location.ID,
		&object.Location.Name,
		&object.Location.BusinessID,
	); err != nil {
		return nil, fmt.Errorf("%s", err)
	}

	return object, nil
}

//
// Create this will create a user in the database
// @params none
// @returns interface - a user
//
func (w WaitListRepository) Create() (interface{}, error) {

	var sqlStm = `SELECT create_waitlist( $1, $2, $3, $4, $5, $6 )`

	if err := w.DB.QueryRow(
		sqlStm,
		w.Model.ClientInfo.ID,
		w.Model.Status.ID,
		w.Model.AlertTime,
		w.Model.Location.ID,
		w.Model.TimeLimit,
		w.Model.Guests,
	).Scan(
		&w.Model.ID,
	); err != nil {
		return nil, fmt.Errorf("%s", err)
	}

	return w.GetByID(w.Model.ID)
}

//
// Update function will update a user in the database
// @params none
// @returns interface - a user
//
func (w WaitListRepository) Update() (interface{}, error) {
	var sqlStm = `SELECT update_waitlist( $1, $2, $3, $4, $5, $6, $7 )`

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
		w.Model.ClientInfo.ID,
		w.Model.Status.ID,
		w.Model.AlertTime,
		w.Model.Location.ID,
		w.Model.TimeLimit,
		w.Model.Guests,
	); err != nil {
		return nil, fmt.Errorf("%s", err)
	}

	if err = tx.Commit(); err != nil {
		return nil, fmt.Errorf("%s", err)
	}

	return w.GetByID(w.Model.ID)
}

//
// Delete function will delete a user in the database
// @params none
// @returns Boolean
//
func (w WaitListRepository) Delete() (bool, error) {

	var sqlStm = `DELETE FROM reservations_waitlist a WHERE a.id = $1`

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

//
// GetByDates will return the waitlists location between two dates
// @param location : integer - the location id
// @param status_id : integer - the status id
// @param startDate ; time - the start date
// @param endDate : time - the end date
//
func (w WaitListRepository) GetByDates(location int, statusID int, startDate time.Time, endDate time.Time) ([]interface{}, error) {
	var sqlStm = `
			SELECT 	a.id,
					a.alert_time,
					a.time_limit,
					a.guests,
					a.timestamp, 
					a.updated,
					a.client_info_id,
					b.first_name,
					b.last_name,
					b.email,
					b.phone,
					a.status,
					c.description,
					c.value,
					a.location_id,
					e.name,
					e.business_id
				FROM 		reservations_waitlist a
				INNER JOIN 	reservations_client_info b ON a.client_info_id = b.id
				INNER JOIN 	reservations_waitlist_status c ON a.status = c.id
				INNER JOIN 	reservations_location e ON a.location_id = e.id
			WHERE 	a.location_id = $1
				AND (SELECT CASE WHEN ( $2 <> 0 )
						THEN ( a.status = $2 )
						ELSE ( a.status > 0 ) END )
				AND a.timestamp BETWEEN $3 AND $4`

	var objects []models.WaitListModel

	rows, err := w.DB.Query(sqlStm, location, statusID, startDate, endDate)

	if err != nil {
		return nil, fmt.Errorf("%s", err)
	}

	defer rows.Close()

	for rows.Next() {
		object := models.WaitListModel{
			ClientInfo: models.ClientInfoModel{},
			Location:   models.LocationModel{},
			Status:     models.WaitListStatusModel{},
		}

		if err = rows.Scan(
			&object.ID,
			&object.AlertTime,
			&object.TimeLimit,
			&object.Guests,
			&object.Timestamp,
			&object.Updated,
			&object.ClientInfo.ID,
			&object.ClientInfo.FirstName,
			&object.ClientInfo.LastName,
			&object.ClientInfo.Email,
			&object.ClientInfo.Phone,
			&object.Status.ID,
			&object.Status.Description,
			&object.Status.Value,
			&object.Location.ID,
			&object.Location.Name,
			&object.Location.BusinessID); err != nil {
			return nil, fmt.Errorf("%s", err)
		}

		objects = append(objects, object)
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
