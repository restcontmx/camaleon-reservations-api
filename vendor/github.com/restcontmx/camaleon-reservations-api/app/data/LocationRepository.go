package data

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
	"github.com/restcontmx/camaleon-reservations-api/app/models"
)

//
// LocationRepository object
// @param Model : LocationModel - the model of the repository
// @param DB : sql db pointer - the db connection of the project
//
type LocationRepository struct {
	Model models.LocationModel
	DB    *sql.DB
}

//
// GetAll will return all the Locations
// @params none
// @return interface array
// @return error - raise an error if so
//
func (l LocationRepository) GetAll() ([]interface{}, error) {
	var sqlStm = `SELECT	id, 
							business_id, 
							crcenter_id, 
							name, 
							description, 
							address1,
							address2,
							phone,
							city,
							state,
							country,
							timestamp, 
							updated 
					FROM reservations_location`
	var locations []models.LocationModel

	rows, err := l.DB.Query(sqlStm)

	if err != nil {
		return nil, fmt.Errorf("%s", err)
	}

	defer rows.Close()

	for rows.Next() {
		var id int
		var businessid int
		var crcenterid int
		var name string
		var description string
		var address1 string
		var address2 string
		var phone string
		var city string
		var state string
		var country string
		var timestamp time.Time
		var updated time.Time

		if err = rows.Scan(
			&id,
			&businessid,
			&crcenterid,
			&name,
			&description,
			&address1,
			&address2,
			&phone,
			&city,
			&state,
			&country,
			&timestamp,
			&updated); err != nil {
			return nil, fmt.Errorf("%s", err)
		}

		locations = append(locations, models.LocationModel{
			ID:          id,
			BusinessID:  businessid,
			CrCenterID:  crcenterid,
			Name:        name,
			Description: description,
			Address1:    address1,
			Address2:    address2,
			Timestamp:   timestamp,
			Updated:     updated,
		})
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("%s", err)
	}

	intfLocations := make([]interface{}, len(locations))

	for i, obj := range locations {
		intfLocations[i] = obj
	}

	return intfLocations, nil
}

//
// GetByID Returns a location by id field
// @param id - int
// @return interface
//
func (l LocationRepository) GetByID(id int) (interface{}, error) {
	var sqlStm = `
		SELECT	a.id, 
				a.business_id, 
				a.crcenter_id, 
				a.name, 
				a.description, 
				a.address1,
				a.address2,
				a.phone,
				a.city,
				a.state,
				a.country,
				a.timestamp, 
				a.updated 
		FROM reservations_location a
		WHERE a.id = $1`

	var location models.LocationModel

	if err := l.DB.QueryRow(
		sqlStm,
		id,
	).Scan(
		&location.ID,
		&location.BusinessID,
		&location.CrCenterID,
		&location.Name,
		&location.Description,
		&location.Address1,
		&location.Address2,
		&location.Phone,
		&location.City,
		&location.State,
		&location.Country,
		&location.Timestamp,
		&location.Updated,
	); err != nil {
		return nil, fmt.Errorf("%s", err)
	}

	return location, nil
}

//
// Create this will create a Location in the database
// @params none
// @returns interface - a location
//
func (l LocationRepository) Create() (interface{}, error) {
	var sqlStm = `SELECT create_location( $1, $2, $3, $4, $5, $6, $7, $8, $9, $10 )`

	tx, err := l.DB.Begin()

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
		l.Model.BusinessID,
		l.Model.CrCenterID,
		l.Model.Name,
		l.Model.Description,
		l.Model.Address1,
		l.Model.Address2,
		l.Model.Phone,
		l.Model.City,
		l.Model.State,
		l.Model.Country,
	); err != nil {
		return nil, fmt.Errorf("%s", err)
	}

	if err = tx.Commit(); err != nil {
		return nil, fmt.Errorf("%s", err)
	}

	return nil, nil
}

//
// Update function will update a location in the database
// @params none
// @returns interface - a location
//
func (l LocationRepository) Update() (interface{}, error) {
	var sqlStm = `SELECT update_location( $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11 )`

	tx, err := l.DB.Begin()

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
		l.Model.ID,
		l.Model.BusinessID,
		l.Model.CrCenterID,
		l.Model.Name,
		l.Model.Description,
		l.Model.Address1,
		l.Model.Address2,
		l.Model.Phone,
		l.Model.City,
		l.Model.State,
		l.Model.Country,
	); err != nil {
		return nil, fmt.Errorf("%s", err)
	}

	if err = tx.Commit(); err != nil {
		return nil, fmt.Errorf("%s", err)
	}

	return l.Model, nil
}

//
// Delete function will delete a location in the database
// @params none
// @returns Boolean
//
func (l LocationRepository) Delete() (bool, error) {
	var sqlStm = `DELETE FROM reservations_location a WHERE a.id = $1`

	tx, err := l.DB.Begin()

	if err != nil {
		return false, fmt.Errorf("%s", err)
	}

	defer tx.Rollback()

	stmt, err := tx.Prepare(sqlStm)

	if err != nil {
		return false, fmt.Errorf("%s", err)
	}

	defer stmt.Close()

	if _, err = stmt.Exec(l.Model.ID); err != nil {
		return false, fmt.Errorf("%s", err)
	}

	if err = tx.Commit(); err != nil {
		return false, fmt.Errorf("%s", err)
	}

	return true, nil
}

//
// GetAllByBusinessID will return all the Locations by business id
// @param business_id : integer - the business id
// @return interface array
// @return error - raise an error if so
//
func (l LocationRepository) GetAllByBusinessID(businessID int) ([]interface{}, error) {
	var sqlStm = `SELECT	id, 
							business_id, 
							crcenter_id,
							name,
							description, 
							address1,
							address2,
							phone,
							city,
							state,
							country,
							timestamp, 
							updated 
					FROM reservations_location
					WHERE business_id = $1`

	var locations []models.LocationModel

	rows, err := l.DB.Query(sqlStm, businessID)

	if err != nil {
		return nil, fmt.Errorf("%s", err)
	}

	defer rows.Close()

	for rows.Next() {
		var id int
		var businessid int
		var crcenterid int
		var name string
		var description string
		var address1 string
		var address2 string
		var phone string
		var city string
		var state string
		var country string
		var timestamp time.Time
		var updated time.Time

		if err = rows.Scan(
			&id,
			&businessid,
			&crcenterid,
			&name,
			&description,
			&address1,
			&address2,
			&phone,
			&city,
			&state,
			&country,
			&timestamp,
			&updated); err != nil {
			return nil, fmt.Errorf("%s", err)
		}

		locations = append(locations, models.LocationModel{
			ID:          id,
			CrCenterID:  crcenterid,
			Name:        name,
			Description: description,
			Address1:    address1,
			Address2:    address2,
			Timestamp:   timestamp,
			Updated:     updated,
		})
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("%s", err)
	}

	intfLocations := make([]interface{}, len(locations))

	for i, obj := range locations {
		intfLocations[i] = obj
	}

	return intfLocations, nil
}

//
// GetByCrCenterID Returns a business by control center business id field
// @param id - int
// @return interface
//
func (l LocationRepository) GetByCrCenterID(crcenterid int) (interface{}, error) {
	var sqlStm = `
		SELECT	a.id, 
				a.crcenter_id,
				a.name, 
				a.description,
				a.timestamp, 
				a.updated 
		FROM reservations_location a
		WHERE a.crcenter_id = $1`

	var location models.LocationModel

	if err := l.DB.QueryRow(
		sqlStm,
		crcenterid,
	).Scan(
		&location.ID,
		&location.CrCenterID,
		&location.Name,
		&location.Description,
		&location.Address1,
		&location.Address2,
		&location.Phone,
		&location.City,
		&location.State,
		&location.Country,
		&location.Timestamp,
		&location.Updated,
	); err != nil {
		return nil, fmt.Errorf("%s", err)
	}

	return location, nil
}
