package data

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/restcontmx/camaleon-reservations-api/app/models"
)

//
// BusinessRepository object
// @param Model : BusinessModel - the model of the repository
// @param DB : sql db pointer - the db connection of the project
//
type BusinessRepository struct {
	Model models.BusinessModel
	DB    *sql.DB
}

//
// GetAll will return all the Businesss
// @params none
// @return interface array
// @return error - raise an error if so
//
func (b BusinessRepository) GetAll() ([]interface{}, error) {
	var sqlStm = "SELECT id, crcenter_id, name, description, timestamp, updated FROM reservations_business"
	var businesses []models.BusinessModel

	rows, err := b.DB.Query(sqlStm)

	if err != nil {
		return nil, fmt.Errorf("%s", err)
	}

	defer rows.Close()

	for rows.Next() {
		var id int
		var crcenterid int
		var name string
		var description string
		var timestamp time.Time
		var updated time.Time

		if err = rows.Scan(&id, &crcenterid, &name, &description, &timestamp, &updated); err != nil {
			return nil, fmt.Errorf("%s", err)
		}

		businesses = append(businesses, models.BusinessModel{
			ID:          id,
			CrCenterID:  crcenterid,
			Name:        name,
			Description: description,
			Timestamp:   timestamp,
			Updated:     updated,
		})
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("%s", err)
	}

	intfBusinesses := make([]interface{}, len(businesses))

	for i, obj := range businesses {
		intfBusinesses[i] = obj
	}

	return intfBusinesses, nil
}

//
// GetByID Returns a business by id field
// @param id - int
// @return interface
//
func (b BusinessRepository) GetByID(id int) (interface{}, error) {
	var sqlStm = `
		SELECT	a.id, 
				a.crcenter_id,
				a.name, 
				a.description,
				a.permalink,
				a.timestamp, 
				a.updated 
		FROM reservations_business a
		WHERE a.id = $1`

	var business models.BusinessModel

	if err := b.DB.QueryRow(
		sqlStm,
		id,
	).Scan(
		&business.ID,
		&business.CrCenterID,
		&business.Name,
		&business.Description,
		&business.Permalink,
		&business.Timestamp,
		&business.Updated,
	); err != nil {
		return nil, fmt.Errorf("%s", err)
	}

	business.Locations, _ = b.GetLocations(business.ID)

	return business, nil
}

//
// GetByCrCenterID Returns a business by control center business id field
// @param id - int
// @return interface
//
func (b BusinessRepository) GetByCrCenterID(crcenterid int) (interface{}, error) {
	var sqlStm = `
		SELECT	a.id, 
				a.crcenter_id,
				a.name, 
				a.description,
				a.permalink,
				a.timestamp, 
				a.updated 
		FROM reservations_business a
		WHERE a.crcenter_id = $1`

	var business models.BusinessModel

	if err := b.DB.QueryRow(
		sqlStm,
		crcenterid,
	).Scan(
		&business.ID,
		&business.CrCenterID,
		&business.Name,
		&business.Description,
		&business.Permalink,
		&business.Timestamp,
		&business.Updated,
	); err != nil {
		return nil, fmt.Errorf("%s", err)
	}

	return business, nil
}

//
// Create this will create a Business in the database
// @params none
// @returns interface - a business
//
func (b BusinessRepository) Create() (interface{}, error) {
	var sqlStm = `SELECT create_business( $1, $2, $3, $4 )`

	tx, err := b.DB.Begin()

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
		b.Model.CrCenterID,
		b.Model.Name,
		b.Model.Description,
		b.Model.Permalink,
	); err != nil {
		return nil, fmt.Errorf("%s", err)
	}

	if err = tx.Commit(); err != nil {
		return nil, fmt.Errorf("%s", err)
	}

	return nil, nil
}

//
// Update function will update a business in the database
// @params none
// @returns interface - a business
//
func (b BusinessRepository) Update() (interface{}, error) {
	var sqlStm = `SELECT update_business( $1, $2, $3, $4, $5 )`

	tx, err := b.DB.Begin()

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
		b.Model.ID,
		b.Model.CrCenterID,
		b.Model.Name,
		b.Model.Description,
		b.Model.Permalink,
	); err != nil {
		return nil, fmt.Errorf("%s", err)
	}

	if err = tx.Commit(); err != nil {
		return nil, fmt.Errorf("%s", err)
	}

	return b.Model, nil
}

//
// Delete function will delete a business in the database
// @params none
// @returns Boolean
//
func (b BusinessRepository) Delete() (bool, error) {
	var sqlStm = `DELETE FROM reservations_business a WHERE a.id = $1`

	tx, err := b.DB.Begin()

	if err != nil {
		return false, fmt.Errorf("%s", err)
	}

	defer tx.Rollback()

	stmt, err := tx.Prepare(sqlStm)

	if err != nil {
		return false, fmt.Errorf("%s", err)
	}

	defer stmt.Close()

	if _, err = stmt.Exec(b.Model.ID); err != nil {
		return false, fmt.Errorf("%s", err)
	}

	if err = tx.Commit(); err != nil {
		return false, fmt.Errorf("%s", err)
	}

	return true, nil
}

//
// GetLocations will return all the locations by business id
// @param businessID : int - business id
// @returns none
//
func (b BusinessRepository) GetLocations(businessID int) ([]models.LocationModel, error) {
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

	rows, err := b.DB.Query(sqlStm, businessID)

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

	return locations, nil
}
