package data

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/restcontmx/camaleon-reservations-api/app/models"
)

// UserReservationRepository repository object
type UserReservationRepository struct {
	Model models.UserReservationModel
	DB    *sql.DB
}

//
// GetAll will return all the objects
// @params none
// @return interface array
// @return error - raise an error if so
//
func (u UserReservationRepository) GetAll() ([]interface{}, error) {
	var sqlStm = `
			SELECT 	a.id, 

					a.business_id,
					b.crcenter_id,
					b.name,
					b.description,
					b.permalink,

					a.user_id,
					u.firstname,
					u.lastname,
					u.username,
					u.email,

					a.rol_id,
					r.description,
					r.value,

					a.timestamp, 
					a.updated
			FROM reservations_userreservation a
			INNER JOIN reservations_business b
			ON b.id = a.business_id
			INNER JOIN reservations_user u
			ON u.id = a.user_id
			INNER JOIN reservations_rol r 
			ON r.id = a.rol_id `

	var objects []models.UserReservationModel

	rows, err := u.DB.Query(sqlStm)

	if err != nil {
		return nil, fmt.Errorf("%s", err)
	}

	defer rows.Close()

	for rows.Next() {

		user := models.UserReservationModel{}

		if err = rows.Scan(
			&user.ID,

			&user.Business.ID,
			&user.Business.CrCenterID,
			&user.Business.Name,
			&user.Business.Description,
			&user.Business.Permalink,

			&user.User.ID,
			&user.User.FirstName,
			&user.User.LastName,
			&user.User.UserName,
			&user.User.Email,

			&user.Rol.ID,
			&user.Rol.Description,
			&user.Rol.Value,

			&user.Timestamp,
			&user.Updated,
		); err != nil {
			return nil, fmt.Errorf("%s", err)
		}

		objects = append(objects, user)
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
func (u UserReservationRepository) GetByID(id int) (interface{}, error) {
	var sqlStm = `
			SELECT 	a.id, 
					a.business_id,
					b.crcenter_id,
					b.name,
					b.description,
					b.permalink,

					a.user_id,
					u.firstname,
					u.lastname,
					u.username,
					u.email,

					a.rol_id,
					r.description,
					r.value,

					a.timestamp, 
					a.updated
			FROM reservations_userreservation a
				INNER JOIN reservations_business b
					ON b.id = a.business_id
				INNER JOIN reservations_user u
					ON u.id = a.user_id
				INNER JOIN reservations_rol r 
					ON r.id = a.rol_id
			WHERE a.id = $1`

	var user models.UserReservationModel

	if err := u.DB.QueryRow(
		sqlStm,
		id,
	).Scan(
		&user.ID,

		&user.Business.ID,
		&user.Business.CrCenterID,
		&user.Business.Name,
		&user.Business.Description,
		&user.Business.Permalink,

		&user.User.ID,
		&user.User.FirstName,
		&user.User.LastName,
		&user.User.UserName,
		&user.User.Email,

		&user.Rol.ID,
		&user.Rol.Description,
		&user.Rol.Value,

		&user.Timestamp,
		&user.Updated,
	); err != nil {
		return nil, fmt.Errorf("%s", err)
	}

	user.AllowedLocations, _ = u.GetAllowedLocations(user.ID)

	return user, nil
}

//
// GetAllowedLocations will return all the allowed locations by user reservation id
// @params none
// @return interface array
// @return error - rise an error if so
//
func (u UserReservationRepository) GetAllowedLocations(userReservationID int) ([]models.LocationModel, error) {
	var sqlStm = `
		SELECT 	a.id, 
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
		INNER JOIN reservations_userreservation_locations ul
			on ul.location_id = a.id
		WHERE ul.userreservation_id = $1`

	var objects []models.LocationModel

	rows, err := u.DB.Query(sqlStm, userReservationID)

	if err != nil {
		return nil, fmt.Errorf("%s", err)
	}

	defer rows.Close()

	for rows.Next() {

		location := models.LocationModel{}

		if err = rows.Scan(
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

		objects = append(objects, location)
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
func (u UserReservationRepository) Create() (interface{}, error) {
	var sqlStm = `SELECT create_userreservation( $1, $2, $3 )`

	if err := u.DB.QueryRow(
		sqlStm,
		u.Model.Business.ID,
		u.Model.User.ID,
		u.Model.Rol.ID,
	).Scan(
		&u.Model.ID,
	); err != nil {
		return nil, fmt.Errorf("%s", err)
	}

	return u.GetByID(u.Model.ID)
}

//
// AddAllowedLocation will add an allowed location to the allowed locations list
// @params none
// @return interface array
// @return error - raise an error if so
//
func (u UserReservationRepository) AddAllowedLocation(locationID int) (interface{}, error) {
	var sqlStm = `SELECT create_userreservation_locations( $1, $2 )`
	var relID int
	log.Println(u.Model.ID)
	log.Println(locationID)
	if err := u.DB.QueryRow(
		sqlStm,
		u.Model.ID,
		locationID,
	).Scan(
		&relID,
	); err != nil {
		return nil, fmt.Errorf("%s", err)
	}

	return nil, nil
}

//
// RemoveAllAllowedLocations will delete all the allowed locations from the user reservation
// @params none
// @returns boolean and error if any
//
func (u UserReservationRepository) RemoveAllAllowedLocations() (bool, error) {
	var sqlStm = `DELETE FROM reservations_userreservation_locations a WHERE a.userreservation_id = $1`

	tx, err := u.DB.Begin()

	if err != nil {
		return false, fmt.Errorf("%s", err)
	}

	defer tx.Rollback()

	stmt, err := tx.Prepare(sqlStm)

	if err != nil {
		return false, fmt.Errorf("%s", err)
	}

	defer stmt.Close()

	if _, err = stmt.Exec(u.Model.ID); err != nil {
		return false, fmt.Errorf("%s", err)
	}

	if err = tx.Commit(); err != nil {
		return false, fmt.Errorf("%s", err)
	}

	return true, nil
}

//
// Update will update an object on the db
// @params none
// @return interface array
// @return error - raise an error if so
//
func (u UserReservationRepository) Update() (interface{}, error) {
	var sqlStm = `SELECT update_userreservation( $1, $2, $3, $4 )`

	tx, err := u.DB.Begin()

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
		u.Model.ID,
		u.Model.Business.ID,
		u.Model.User.ID,
		u.Model.Rol,
	); err != nil {
		return nil, fmt.Errorf("%s", err)
	}

	if err = tx.Commit(); err != nil {
		return nil, fmt.Errorf("%s", err)
	}

	return u.GetByID(u.Model.ID)
}

//
// Delete function will delete an object in the database
// @params none
// @returns Boolean
//
func (u UserReservationRepository) Delete() (bool, error) {
	var sqlStm = `DELETE FROM update_userreservation a WHERE a.id = $1`

	tx, err := u.DB.Begin()

	if err != nil {
		return false, fmt.Errorf("%s", err)
	}

	defer tx.Rollback()

	stmt, err := tx.Prepare(sqlStm)

	if err != nil {
		return false, fmt.Errorf("%s", err)
	}

	defer stmt.Close()

	if _, err = stmt.Exec(u.Model.ID); err != nil {
		return false, fmt.Errorf("%s", err)
	}

	if err = tx.Commit(); err != nil {
		return false, fmt.Errorf("%s", err)
	}

	return true, nil
}
