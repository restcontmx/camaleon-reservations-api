package models

import "time"

//
// ReservationModel is the main reservation model for reservating tables
// @param ID : string - the unique id of reservation
// @param Location : LocationModel - the location the reservation has been made to
// @param Table : TaleModel - the table reserved
// @param Date : Date Time - the date time the table is reserved
// @param Timestamp : Date Time - the date time the reservation has been created
// @param Updated : Date Time - the date time the reservation has been updated
//
type ReservationModel struct {
	ID         int
	UID        string
	Table      TableModel
	Location   LocationModel
	Status     ReservationStatusModel
	ClientInfo ClientInfoModel
	TimeLimit  int
	Guests     int
	Date       time.Time
	Timestamp  time.Time
	Updated    time.Time
}
