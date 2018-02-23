package models

import (
	"time"
)

//
// BusinessModel is the model for the business
// @param ID : integer - the unique id of the business
// @param Name : string - the unique name of the business
// @param Description : string - the description of the business
// @param Timestamp : Date Time - the date it was created
// @param Updated : Date Time - the date it was updated
//
type BusinessModel struct {
	ID          int
	CrCenterID  int
	Name        string
	Description string
	Locations   []LocationModel
	Timestamp   time.Time
	Updated     time.Time
}
