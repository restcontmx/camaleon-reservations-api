package models

import (
	"time"
)

//
// ClientModel is the client reference from the auth0
// @param Id : integer - the unique id of the client
// @param Auth0Id : string - is the uniuqe id from the auth0 framework
//
type ClientModel struct {
	ID        int
	Auth0ID   string
	Timestamp time.Time
	Updated   time.Time
}
