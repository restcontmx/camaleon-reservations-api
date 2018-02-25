package data

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/restcontmx/camaleon-reservations-api/app/models"
)

//
// ReservationRepository this will be the real user object
//
type ReservationRepository struct {
	Model models.ReservationModel
	DB    *sql.DB
}

//
// GetAll will return all the users
// @params none
// @return interface array
// @return error - rise an error if so
//
func (r ReservationRepository) GetAll() ([]interface{}, error) {
	var sqlStm = `
				SELECT 	a.id, 
						a.uid,
						a.date,
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
						a.table_id,
						d.name,
						d.description,
						d.img_url,
						d.max_guests,
						d.area_id,
						area.name,
						area.description,
						area.img_url,
						a.location_id,
						e.name
				FROM 			reservations_reservation a
					INNER JOIN 	reservations_client_info b ON a.client_info_id = b.id
					INNER JOIN 	reservations_reservation_status c ON a.status = c.id
					INNER JOIN 	reservations_table d ON a.table_id = d.id
					INNER JOIN 	reservations_area area ON d.area_id = area.id
					INNER JOIN 	reservations_location e ON a.location_id = e.id`

	var objects []models.ReservationModel

	rows, err := r.DB.Query(sqlStm)

	if err != nil {
		return nil, fmt.Errorf("%s", err)
	}

	defer rows.Close()

	for rows.Next() {
		reservation := models.ReservationModel{
			ClientInfo: models.ClientInfoModel{},
			Location:   models.LocationModel{},
			Table: models.TableModel{
				Area: models.AreaModel{},
			},
			Status: models.ReservationStatusModel{},
		}

		if err = rows.Scan(
			&reservation.ID,
			&reservation.UID,
			&reservation.Date,
			&reservation.TimeLimit,
			&reservation.Guests,
			&reservation.Timestamp,
			&reservation.Updated,
			&reservation.ClientInfo.ID,
			&reservation.ClientInfo.FirstName,
			&reservation.ClientInfo.LastName,
			&reservation.ClientInfo.Email,
			&reservation.ClientInfo.Phone,
			&reservation.Status.ID,
			&reservation.Status.Description,
			&reservation.Status.Value,
			&reservation.Table.ID,
			&reservation.Table.Name,
			&reservation.Table.Description,
			&reservation.Table.ImgURL,
			&reservation.Table.MaxGuests,
			&reservation.Table.Area.ID,
			&reservation.Table.Area.Name,
			&reservation.Table.Area.Description,
			&reservation.Table.Area.ImgURL,
			&reservation.Location.ID,
			&reservation.Location.Name); err != nil {
			return nil, fmt.Errorf("%s", err)
		}

		objects = append(objects, reservation)
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
func (r ReservationRepository) GetByID(id int) (interface{}, error) {
	var sqlStm = `
				SELECT 	a.id, 
						a.uid,
						a.date,
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
						a.table_id,
						d.name,
						d.description,
						d.img_url,
						d.max_guests,
						d.area_id,
						area.name,
						area.description,
						area.img_url,
						a.location_id,
						e.name
				FROM 			reservations_reservation a
					INNER JOIN 	reservations_client_info b ON a.client_info_id = b.id
					INNER JOIN 	reservations_reservation_status c ON a.status = c.id
					INNER JOIN 	reservations_table d ON a.table_id = d.id
					INNER JOIN 	reservations_area area ON d.area_id = area.id
					INNER JOIN 	reservations_location e ON a.location_id = e.id
				WHERE a.id = $1`

	reservation := models.ReservationModel{
		ClientInfo: models.ClientInfoModel{},
		Location:   models.LocationModel{},
		Table: models.TableModel{
			Area: models.AreaModel{},
		},
		Status: models.ReservationStatusModel{},
	}

	if err := r.DB.QueryRow(
		sqlStm,
		id,
	).Scan(
		&reservation.ID,
		&reservation.UID,
		&reservation.Date,
		&reservation.TimeLimit,
		&reservation.Guests,
		&reservation.Timestamp,
		&reservation.Updated,
		&reservation.ClientInfo.ID,
		&reservation.ClientInfo.FirstName,
		&reservation.ClientInfo.LastName,
		&reservation.ClientInfo.Email,
		&reservation.ClientInfo.Phone,
		&reservation.Status.ID,
		&reservation.Status.Description,
		&reservation.Status.Value,
		&reservation.Table.ID,
		&reservation.Table.Name,
		&reservation.Table.Description,
		&reservation.Table.ImgURL,
		&reservation.Table.MaxGuests,
		&reservation.Table.Area.ID,
		&reservation.Table.Area.Name,
		&reservation.Table.Area.Description,
		&reservation.Table.Area.ImgURL,
		&reservation.Location.ID,
		&reservation.Location.Name,
	); err != nil {
		return nil, fmt.Errorf("%s", err)
	}

	return reservation, nil
}

//
// Create this will create a user in the database
// @params none
// @returns interface - a user
//
func (r ReservationRepository) Create() (interface{}, error) {

	var sqlStm = `SELECT create_reservation( uuid_generate_v4(), $1, $2, $3, $4, $5, $6, $7 )`

	if err := r.DB.QueryRow(
		sqlStm,
		r.Model.ClientInfo.ID,
		r.Model.Status.ID,
		r.Model.Date,
		r.Model.Table.ID,
		r.Model.Location.ID,
		r.Model.TimeLimit,
		r.Model.Guests,
	).Scan(
		&r.Model.ID,
	); err != nil {
		return nil, fmt.Errorf("%s", err)
	}

	return r.GetByID(r.Model.ID)
}

//
// Update function will update a user in the database
// @params none
// @returns interface - a user
//
func (r ReservationRepository) Update() (interface{}, error) {
	var sqlStm = `SELECT update_reservation( $1, $2, $3, $4, $5, $6, $7, $8 )`

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
		r.Model.ClientInfo.ID,
		r.Model.Status.ID,
		r.Model.Date,
		r.Model.Table.ID,
		r.Model.Location.ID,
		r.Model.TimeLimit,
		r.Model.Guests,
	); err != nil {
		return nil, fmt.Errorf("%s", err)
	}

	if err = tx.Commit(); err != nil {
		return nil, fmt.Errorf("%s", err)
	}

	return r.GetByID(r.Model.ID)
}

//
// Delete function will delete a user in the database
// @params none
// @returns Boolean
//
func (r ReservationRepository) Delete() (bool, error) {

	var sqlStm = `DELETE FROM reservations_reservation a WHERE a.id = $1`

	tx, err := r.DB.Begin()

	if err != nil {
		return false, fmt.Errorf("%s", err)
	}

	defer tx.Rollback()

	stmt, err := tx.Prepare(sqlStm)

	if err != nil {
		return false, fmt.Errorf("%s", err)
	}

	defer stmt.Close()

	if _, err = stmt.Exec(r.Model.ID); err != nil {
		return false, fmt.Errorf("%s", err)
	}

	if err = tx.Commit(); err != nil {
		return false, fmt.Errorf("%s", err)
	}

	return true, nil
}

//
// GetByDates will return the reservations location between two dates
// @param location : integer - the location id
// @param status_id : integer - the status id
// @param startDate ; time - the start date
// @param endDate : time - the end date
//
func (r ReservationRepository) GetByDates(location int, statusID int, startDate time.Time, endDate time.Time) ([]interface{}, error) {
	var sqlStm = `
				SELECT 	a.id, 
						a.uid,
						a.date,
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
						a.table_id,
						d.name,
						d.description,
						d.img_url,
						d.max_guests,
						d.area_id,
						area.name,
						area.description,
						area.img_url,
						a.location_id,
						e.name
				FROM 			reservations_reservation a
					INNER JOIN 	reservations_client_info b ON a.client_info_id = b.id
					INNER JOIN 	reservations_reservation_status c ON a.status = c.id
					INNER JOIN 	reservations_table d ON a.table_id = d.id
					INNER JOIN 	reservations_area area ON d.area_id = area.id
					INNER JOIN 	reservations_location e ON a.location_id = e.id
				WHERE 	a.location_id = $1
					AND (SELECT CASE WHEN ( $2 <> 0 )
							THEN ( a.status = $2 )
							ELSE ( a.status > 0 ) END )
					AND a.date BETWEEN $3 AND $4`

	var objects []models.ReservationModel

	rows, err := r.DB.Query(sqlStm, location, statusID, startDate, endDate)

	if err != nil {
		return nil, fmt.Errorf("%s", err)
	}

	defer rows.Close()

	for rows.Next() {
		reservation := models.ReservationModel{
			ClientInfo: models.ClientInfoModel{},
			Location:   models.LocationModel{},
			Table: models.TableModel{
				Area: models.AreaModel{},
			},
			Status: models.ReservationStatusModel{},
		}

		if err = rows.Scan(
			&reservation.ID,
			&reservation.UID,
			&reservation.Date,
			&reservation.TimeLimit,
			&reservation.Guests,
			&reservation.Timestamp,
			&reservation.Updated,
			&reservation.ClientInfo.ID,
			&reservation.ClientInfo.FirstName,
			&reservation.ClientInfo.LastName,
			&reservation.ClientInfo.Email,
			&reservation.ClientInfo.Phone,
			&reservation.Status.ID,
			&reservation.Status.Description,
			&reservation.Status.Value,
			&reservation.Table.ID,
			&reservation.Table.Name,
			&reservation.Table.Description,
			&reservation.Table.ImgURL,
			&reservation.Table.MaxGuests,
			&reservation.Table.Area.ID,
			&reservation.Table.Area.Name,
			&reservation.Table.Area.Description,
			&reservation.Table.Area.ImgURL,
			&reservation.Location.ID,
			&reservation.Location.Name); err != nil {
			return nil, fmt.Errorf("%s", err)
		}

		objects = append(objects, reservation)
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
