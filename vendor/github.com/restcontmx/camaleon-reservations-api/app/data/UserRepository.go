package data

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"

	"github.com/restcontmx/camaleon-reservations-api/app/models"
	"golang.org/x/crypto/bcrypt"
)

//
// UserRepository this will be the real user object
// @param Model : UserModel - the model of the repository
// @param DB : sql db pointer - the db connection of the project
//
type UserRepository struct {
	Model models.UserModel
	DB    *sql.DB
}

//
// GetAll will return all the users
// @params none
// @return interface array
// @return error - raise an error if so
//
func (u UserRepository) GetAll() ([]interface{}, error) {
	
	var sqlStm = "SELECT id, firstname, lastname, username, email, password, timestamp, updated FROM reservations_user"
	var users []models.UserModel

	rows, err := u.DB.Query(sqlStm)

	if err != nil {
		return nil, fmt.Errorf("%s", err)
	}

	defer rows.Close()

	for rows.Next() {
		var id int
		var firstname string
		var lastname string
		var username string
		var password string
		var email string
		var timestamp time.Time
		var updated time.Time

		if err = rows.Scan(&id, &firstname, &lastname, &username, &email, &password, &timestamp, &updated); err != nil {
			return nil, fmt.Errorf("%s", err)
		}

		users = append(users, models.UserModel{
			ID:        id,
			FirstName: firstname,
			LastName:  lastname,
			Email:     email,
			UserName:  username,
			Password:  password,
			Timestamp: timestamp,
			Updated:   updated,
		})
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("%s", err)
	}

	intfUsers := make([]interface{}, len(users))

	for i, obj := range users {
		intfUsers[i] = obj
	}

	return intfUsers, nil
}

//
// GetByID Returns a user by id field
// @param id - int
// @return interface
//
func (u UserRepository) GetByID(id int) (interface{}, error) {

	var sqlStm = `
		SELECT	a.id, 
			a.firstname, 
			a.lastname, 
			a.username, 
			a.email, 
			a.password, 
			a.timestamp, 
			a.updated 
		FROM reservations_user a
		WHERE a.id = $1`

	var user models.UserModel

	if err := u.DB.QueryRow(
		sqlStm,
		id,
	).Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.UserName,
		&user.Email,
		&user.Password,
		&user.Timestamp,
		&user.Updated,
	); err != nil {
		return nil, fmt.Errorf("%s", err)
	}

	return user, nil
}

//
// Create this will create a user in the database
// @params none
// @returns interface - a user
//
func (u UserRepository) Create() (interface{}, error) {

	var sqlStm = `SELECT create_user( $1, $2, $3, $4, $5 )`

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

	hash, err := bcrypt.GenerateFromPassword([]byte(u.Model.Password), bcrypt.DefaultCost)

	if err != nil {
		return nil, fmt.Errorf("%s", err)
	}

	if _, err = stmt.Exec(
		u.Model.FirstName,
		u.Model.LastName,
		u.Model.UserName,
		u.Model.Email,
		string(hash),
	); err != nil {
		return nil, fmt.Errorf("%s", err)
	}

	if err = tx.Commit(); err != nil {
		return nil, fmt.Errorf("%s", err)
	}

	return u.AuthenticateUser()

}

//
// Update function will update a user in the database
// @params none
// @returns interface - a user
//
func (u UserRepository) Update() (interface{}, error) {

	var sqlStm = `SELECT update_user( $1, $2, $3, $4, $5, $6 )`

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
		u.Model.FirstName,
		u.Model.LastName,
		u.Model.UserName,
		u.Model.Email,
		u.Model.Password,
	); err != nil {
		return nil, fmt.Errorf("%s", err)
	}

	if err = tx.Commit(); err != nil {
		return nil, fmt.Errorf("%s", err)
	}

	return u.Model, nil
}

//
// Delete function will delete a user in the database
// @params none
// @returns Boolean
//
func (u UserRepository) Delete() (bool, error) {

	var sqlStm = `DELETE FROM reservations_user a WHERE a.id = $1`

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
// AuthenticateUser will authenticate the user on the DB
// @params none
// @return interface : any object - the user
// @return error : error - raise if any
//
func (u UserRepository) AuthenticateUser() (interface{}, error) {
	var sqlStm = `
		SELECT	a.id, 
				a.firstname, 
				a.lastname, 
				a.username,
				a.email, 
				a.password, 
				a.timestamp, 
				a.updated 
		FROM reservations_user a
		WHERE a.username = $1`

	var user models.UserModel

	if err := u.DB.QueryRow(
		sqlStm,
		u.Model.UserName,
	).Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.UserName,
		&user.Email,
		&user.Password,
		&user.Timestamp,
		&user.Updated,
	); err != nil {
		return nil, fmt.Errorf("%s", err)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(u.Model.Password)); err != nil {
		return nil, fmt.Errorf("Invalid Crednetials")
	}

	return user, nil

}
