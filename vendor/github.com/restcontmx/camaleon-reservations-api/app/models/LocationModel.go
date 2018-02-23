package models

import "time"

//
// LocationModel Will be the locations model for the business
// @param Id : integer - the unique id of the model
// @param Name : string - the name of the location
// @param Description : string the description of the location
// @param Users : Array of UserModel - this will be the list of users; each location has its users
// @param Timestamp : Date Time - the time location was created
// @param Updated : Date Time - the time location was updated
//
type LocationModel struct {
	ID          int
	BusinessID  int
	CrCenterID  int
	Name        string
	Description string
	Address1    string
	Address2    string
	Phone       string
	City        string
	State       string
	Country     string
	Users       []UserModel
	Timestamp   time.Time
	Updated     time.Time
}
