package data

import (
	"database/sql"
	"fmt"

	"github.com/restcontmx/camaleon-reservations-api/app/models"
)

//
// TableRepository repository object
//
type TableRepository struct {
	Model models.TableModel
	DB    *sql.DB
}

//
// Create will create an object on the db
// @params none
// @return interface array
// @return error - raise an error if so
//
func (t TableRepository) Create() (interface{}, error) {
	var sqlStm = `SELECT create_table( $1, $2, $3, $4, $5 )`

	tx, err := t.DB.Begin()

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
		t.Model.Name,
		t.Model.Description,
		t.Model.Area.ID,
		t.Model.ImgURL,
		t.Model.MaxGuests,
	); err != nil {
		return nil, fmt.Errorf("%s", err)
	}

	if err = tx.Commit(); err != nil {
		return nil, fmt.Errorf("%s", err)
	}

	return t.Model, nil
}

//
// GetAll will return all the Objects
// @params none
// @return interface array
// @return error - rise an error if so
//
func (t TableRepository) GetAll() ([]interface{}, error) {
	var sqlStm = `
				SELECT 	a.id, 
						a.name,
						a.description,
						a.img_url,
						a.max_guests,
						a.area_id,
						b.name,
						b.description,
						b.img_url,
						c.id,
						c.name,
						c.description
				FROM reservations_table a
				INNER JOIN reservations_area b
					ON b.id = a.area_id
				INNER JOIN reservations_location c
					ON b.location_id = c.id
				`

	var objects []models.TableModel

	rows, err := t.DB.Query(sqlStm)

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
		var areaid int
		var areaName string
		var areaDescription string
		var areaImgURL string
		var locationID int
		var locationName string
		var locationDescription string

		if err = rows.Scan(
			&id,
			&name,
			&description,
			&imgurl,
			&maxGuests,
			&areaid,
			&areaName,
			&areaDescription,
			&areaImgURL,
			&locationID,
			&locationName,
			&locationDescription); err != nil {
			return nil, fmt.Errorf("%s", err)
		}

		objects = append(objects, models.TableModel{
			ID:          id,
			Name:        name,
			Description: description,
			ImgURL:      imgurl,
			MaxGuests:   maxGuests,
			Area: models.AreaModel{
				ID:          areaid,
				Name:        areaName,
				Description: areaDescription,
				ImgURL:      areaImgURL,
				Location: models.LocationModel{
					ID:          locationID,
					Name:        locationName,
					Description: locationDescription,
				},
			},
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
// GetAllByAreaID will return all the Objects by area id
// @params none
// @return interface array
// @return error - rise an error if so
//
func (t TableRepository) GetAllByAreaID(areaID int) ([]interface{}, error) {
	var sqlStm = `
				SELECT 	a.id, 
						a.name,
						a.description,
						a.img_url,
						a.max_guests
				FROM reservations_table a WHERE a.area_id = $1
				`

	var objects []models.TableModel

	rows, err := t.DB.Query(sqlStm, areaID)

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

	intfObjects := make([]interface{}, len(objects))

	for i, obj := range objects {
		intfObjects[i] = obj
	}

	return intfObjects, nil
}

//
// GetByID get object by id
//
func (t TableRepository) GetByID(id int) (interface{}, error) {
	var sqlStm = `
				SELECT 	a.id, 
						a.name,
						a.description,
						a.img_url,
						a.max_guests,
						a.area_id,
						b.name,
						b.description,
						b.img_url,
						c.id,
						c.name,
						c.description
					FROM reservations_table a
					INNER JOIN reservations_area b
						ON b.id = a.area_id
					INNER JOIN reservations_location c
						ON b.location_id = c.id
					WHERE a.id = $1`

	object := models.TableModel{
		Area: models.AreaModel{
			Location: models.LocationModel{},
		},
	}

	if err := t.DB.QueryRow(
		sqlStm,
		id,
	).Scan(
		&object.ID,
		&object.Name,
		&object.Description,
		&object.ImgURL,
		&object.MaxGuests,
		&object.Area.ID,
		&object.Area.Name,
		&object.Area.Description,
		&object.Area.ImgURL,
		&object.Area.Location.ID,
		&object.Area.Location.Name,
		&object.Area.Location.Description,
	); err != nil {
		return nil, fmt.Errorf("%s", err)
	}

	return object, nil
}

//
// Update function will update an object in the database
// @params none
// @returns interface - a user
//
func (t TableRepository) Update() (interface{}, error) {

	var sqlStm = `SELECT update_table( $1, $2, $3, $4, $5, $6 )`

	tx, err := t.DB.Begin()

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
		t.Model.ID,
		t.Model.Name,
		t.Model.Description,
		t.Model.Area.ID,
		t.Model.ImgURL,
		t.Model.MaxGuests,
	); err != nil {
		return nil, fmt.Errorf("%s", err)
	}

	if err = tx.Commit(); err != nil {
		return nil, fmt.Errorf("%s", err)
	}

	return t.Model, nil
}
