package models

import "time"

//
// UserModel is the main model for user authentication in the application
// @param ID : integer - the unique id of the user
// @param FirstName : string - the first name of the user
// @param LastName : string - the last name of the user
// @param UserName : string - the username of the user
// @param Email : string - the email of the user
// @param Password : string - the user password for authentication
// @param Timestamp : Date Time - the date the user has been createed
// @param Updated : Date Time - the date the user has been updated
//
type UserModel struct {
	ID        int       `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	UserName  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Timestamp time.Time `json:"timestamp"`
	Updated   time.Time `json:"updated"`
}
