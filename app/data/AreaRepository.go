package data

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/restcontmx/camaleon-reservations-api/app/models"
)

// AreaRepository main repository
type AreaRepository struct {
	Model models.AreaModel
	DB    *sql.DB
}

//
// GetAll will return all the Areas
// @params none
// @return interface array
// @return error - rise an error if so
//
func (a AreaRepository) GetAll() ([]interface{}, error) {
	var sqlStm = `
				SELECT 	a.id, 
						a.name,
						a.description,
						a.img_url,
						a.timestamp, 
						a.updated,
						b.id,
						b.name,
						b.description,
						b.crcenter_id,
						b.business_id
				FROM reservations_area a
				INNER JOIN reservations_location b
					ON b.id = a.location_id
				`

	var objects []models.AreaModel

	rows, err := a.DB.Query(sqlStm)

	if err != nil {
		return nil, fmt.Errorf("%s", err)
	}

	defer rows.Close()

	for rows.Next() {
		var id int
		var name string
		var description string
		var imgurl string
		var timestamp time.Time
		var updated time.Time
		var locationid int
		var locationName string
		var locationDescription string
		var locationCrCenterID int
		var businessID int

		if err = rows.Scan(
			&id,
			&name,
			&description,
			&imgurl,
			&timestamp,
			&updated,
			&locationid,
			&locationName,
			&locationDescription,
			&locationCrCenterID,
			&businessID); err != nil {
			return nil, fmt.Errorf("%s", err)
		}

		tables, _ := a.GetTables(id)

		objects = append(objects, models.AreaModel{
			ID:        id,
			Name:      name,
			Timestamp: timestamp,
			Updated:   updated,
			ImgURL:    imgurl,
			Location: models.LocationModel{
				ID:          locationid,
				Name:        locationName,
				Description: locationDescription,
				CrCenterID:  locationCrCenterID,
				BusinessID:  businessID,
			},
			Tables: tables,
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
// GetAllByLocation will return all the Areas
// @param locationID : integer - the location id
// @return interface array
// @return error - rise an error if so
//
func (a AreaRepository) GetAllByLocation(locationID int) ([]interface{}, error) {
	var sqlStm = `
				SELECT 	a.id, 
						a.name,
						a.description,
						a.img_url,
						a.timestamp, 
						a.updated,
						b.id,
						b.name,
						b.description,
						b.crcenter_id,
						b.business_id
				FROM reservations_area a
				INNER JOIN reservations_location b
					ON b.id = a.location_id
				WHERE b.id = $1
				`

	var objects []models.AreaModel

	rows, err := a.DB.Query(sqlStm, locationID)

	if err != nil {
		return nil, fmt.Errorf("%s", err)
	}

	defer rows.Close()

	for rows.Next() {
		var id int
		var name string
		var description string
		var imgurl string
		var timestamp time.Time
		var updated time.Time
		var locationid int
		var locationName string
		var locationDescription string
		var locationCrCenterID int
		var businessID int

		if err = rows.Scan(
			&id,
			&name,
			&description,
			&imgurl,
			&timestamp,
			&updated,
			&locationid,
			&locationName,
			&locationDescription,
			&locationCrCenterID,
			&businessID); err != nil {
			return nil, fmt.Errorf("%s", err)
		}

		tables, _ := a.GetTables(id)

		objects = append(objects, models.AreaModel{
			ID:        id,
			Name:      name,
			Timestamp: timestamp,
			Updated:   updated,
			ImgURL:    imgurl,
			Location: models.LocationModel{
				ID:          locationid,
				Name:        locationName,
				Description: locationDescription,
				CrCenterID:  locationCrCenterID,
				BusinessID:  businessID,
			},
			Tables: tables,
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
// GetByID get object by id
//
func (a AreaRepository) GetByID(id int) (interface{}, error) {
	var sqlStm = `
			SELECT 	a.id, 
					a.name,
					a.description,
					a.img_url,
					a.timestamp, 
					a.updated,
					b.id,
					b.name,
					b.description,
					b.crcenter_id,
					b.business_id
			FROM reservations_area a INNER JOIN reservations_location b
				ON b.id = a.location_id
			WHERE a.id = $1`

	var object models.AreaModel
	object.Location = models.LocationModel{}

	if err := a.DB.QueryRow(
		sqlStm,
		id,
	).Scan(
		&object.ID,
		&object.Name,
		&object.Description,
		&object.ImgURL,
		&object.Timestamp,
		&object.Updated,
		&object.Location.ID,
		&object.Location.Name,
		&object.Location.Description,
		&object.Location.CrCenterID,
		&object.Location.BusinessID,
	); err != nil {
		return nil, fmt.Errorf("%s", err)
	}

	object.Tables, _ = a.GetTables(object.ID)

	return object, nil
}

//
// GetTables will return all the Objects by area id
// @params none
// @return interface array
// @return error - rise an error if so
//
func (a AreaRepository) GetTables(areaID int) ([]models.TableModel, error) {
	var sqlStm = `
				SELECT 	a.id, 
						a.name,
						a.description,
						a.img_url,
						a.max_guests
				FROM reservations_table a WHERE a.area_id = $1
				`

	var objects []models.TableModel

	rows, err := a.DB.Query(sqlStm, areaID)

	if err != nil {
		return nil, fmt.Errorf("%s", err)
	}

	defer rows.Close()

	for rows.Next() {
		var id int
		var name string
		var description string
		var imgurl string
		var maxGuests int

		if err = rows.Scan(
			&id,
			&name,
			&description,
			&imgurl,
			&maxGuests); err != nil {
			return nil, fmt.Errorf("%s", err)
		}

		objects = append(objects, models.TableModel{
			ID:          id,
			Name:        name,
			Description: description,
			ImgURL:      imgurl,
			MaxGuests:   maxGuests,
		})
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("%s", err)
	}

	return objects, nil
}

//
// Create will create an object on the db
// @params none
// @return interface array
// @return error - raise an error if so
//
func (a AreaRepository) Create() (interface{}, error) {
	var sqlStm = `SELECT create_area( $1, $2, $3, $4 )`

	tx, err := a.DB.Begin()

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
		a.Model.Name,
		a.Model.Description,
		a.Model.Location.ID,
		a.Model.ImgURL,
	); err != nil {
		return nil, fmt.Errorf("%s", err)
	}

	if err = tx.Commit(); err != nil {
		return nil, fmt.Errorf("%s", err)
	}

	return nil, nil
}

//
// Update function will update an object in the database
// @params none
// @returns interface - a user
//
func (a AreaRepository) Update() (interface{}, error) {

	var sqlStm = `SELECT update_area( $1, $2, $3, $4, $5 )`

	tx, err := a.DB.Begin()

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
		a.Model.ID,
		a.Model.Name,
		a.Model.Description,
		a.Model.Location.ID,
		a.Model.ImgURL,
	); err != nil {
		return nil, fmt.Errorf("%s", err)
	}

	if err = tx.Commit(); err != nil {
		return nil, fmt.Errorf("%s", err)
	}

	return a.Model, nil
}
