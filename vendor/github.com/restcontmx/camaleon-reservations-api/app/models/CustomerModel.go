package models

import (
	"time"
)

//
// CustomerModel is for client information
// @param FirstName : string - the first name of the client
// @param LastName ; string - the last name of the client
// @param Address1 : string - the first address field of the client
// @param Address2 ; string - the second address field of the client
// @param Phone : string - clients phone number for contact this is the field we are going to find the user with
// @param Email : sring - clients email address
// @param ZipCode : string - zip code
// @param Timestamp : Date time - date created
// @param Updated : Date time - date updated
//
type CustomerModel struct {
	ID        int
	LastName  string
	FirstName string
	Address1  string
	Address2  string
	Phone     string
	Email     string
	ZipCode   string
	Timestamp time.Time
	Updated   time.Time
	Location  LocationModel
}
