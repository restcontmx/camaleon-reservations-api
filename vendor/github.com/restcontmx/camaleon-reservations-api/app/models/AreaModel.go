package models

import (
	"time"
)

//
// AreaModel is the area model, in peru is called 'Salon'
// @param ID : integer - the unique id of the area
// @param Name : string - the name of the area
// @param Description : string - the description of the area
// @prams Tables : Table Model Array - the tables included in the area
//
type AreaModel struct {
	ID          int
	Name        string
	Description string
	Tables      []TableModel
	ImgURL      string
	Location    LocationModel
	Timestamp   time.Time
	Updated     time.Time
}
