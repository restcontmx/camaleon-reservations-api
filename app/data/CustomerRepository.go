package data

import (
	"fmt"

	"github.com/restcontmx/camaleon-reservations-api/app/models"
)

var customers []models.CustomerModel

//
// CustomerRepository this is the customer repository
// @param model : Customer Model - the repository model reference
//
type CustomerRepository struct {
	Model models.CustomerModel
}

//
// GetAll will return all the customers
// @params none
// @return interface array
// @return error - raise an error if so
//
func (c CustomerRepository) GetAll() ([]interface{}, error) {
	retObjs := make([]interface{}, len(customers))
	for i, obj := range customers {
		retObjs[i] = obj
	}
	return retObjs, nil
}

//
// GetByID returns a customer by id
// @param id : int - the customer id
// @return interface - any object
// @return error - raise error if any
//
func (c CustomerRepository) GetByID(id int) (interface{}, error) {
	return nil, nil
}

//
// Create this will create a customer in the database
// @params none
// @return interface - any object
// @return error - raise if any
//
func (c CustomerRepository) Create() (interface{}, error) {
	return nil, nil
}

//
// Update this will update a customer in the database
// @params none
// @return interface - any object
// @return error : raise an error if any
//
func (c CustomerRepository) Update() (interface{}, error) {
	return nil, nil
}

//
// Delete this will delete a customer on the db
// @params none
// @return bool : if deleted
// @return error : raise if deleted
//
func (c CustomerRepository) Delete() (bool, error) {
	return false, nil
}

//
// GetByParams returns the object by the repository model parameters in it
// @params none
// @return interface - any object
// @return error - raise error if any
//
func (c CustomerRepository) GetByParams() (interface{}, error) {
	for _, obj := range customers {
		if obj.Location.ID == c.Model.Location.ID && obj.Phone == c.Model.Phone {
			return obj, nil
		}
	}
	return nil, fmt.Errorf("Object not found")
}

//
// InitCustomers will create dummy customers
// @params none
// @returns none
//
func InitCustomers() {

	loc1 := models.LocationModel{ID: 2, Name: "Loc 1"}
	loc2 := models.LocationModel{ID: 3, Name: "Loc 2"}

	c1 := models.CustomerModel{
		ID:        34,
		FirstName: "a",
		LastName:  "b",
		Location:  loc1,
		Address1:  "Address ab",
		Address2:  "Address ab 2",
		Phone:     "4931143334",
		Email:     "gunt.raro@gmail.com",
	}

	c2 := models.CustomerModel{
		ID:        35,
		FirstName: "a2",
		LastName:  "b2",
		Location:  loc1,
		Address1:  "Address ab2",
		Address2:  "Address ab2 2",
		Phone:     "4931143335",
		Email:     "gunt.raro@gmail.com",
	}

	c3 := models.CustomerModel{
		ID:        36,
		FirstName: "a3",
		LastName:  "b3",
		Location:  loc2,
		Address1:  "Address ab3",
		Address2:  "Address ab3 2",
		Phone:     "4931143336",
		Email:     "gunt.raro@gmail.com",
	}

	c4 := models.CustomerModel{
		ID:        37,
		FirstName: "a4",
		LastName:  "b4",
		Location:  loc2,
		Address1:  "Address ab4",
		Address2:  "Address ab4 2",
		Phone:     "4931143337",
		Email:     "gunt.raro@gmail.com",
	}

	customers = append(customers, c1)
	customers = append(customers, c2)
	customers = append(customers, c3)
	customers = append(customers, c4)
}
