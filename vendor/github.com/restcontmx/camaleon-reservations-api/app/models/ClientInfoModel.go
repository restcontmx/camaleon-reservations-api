package models

import (
	"time"
)

//
// ClientInfoModel is the main information of the client
//
type ClientInfoModel struct {
	ID        int
	FirstName string
	LastName  string
	Email     string
	Phone     string
	ClientRef ClientModel
	Timestamp time.Time
	Updated   time.Time
}
