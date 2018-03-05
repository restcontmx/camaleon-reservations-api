package models

import (
	"time"
)

//
// UserReservationModel will be the main reservation panel model
//
type UserReservationModel struct {
	ID               int
	Business         BusinessModel
	User             UserModel
	AllowedLocations []LocationModel
	Rol              RolModel
	Timestamp        time.Time
	Updated          time.Time
}
