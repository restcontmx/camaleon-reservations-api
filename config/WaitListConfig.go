package config

import "github.com/restcontmx/camaleon-reservations-api/app/data"

// WaitListConfiguration is the configuration struct
type WaitListConfiguration struct {
	Repository *data.WaitListRepository
}

// WaitListConfig configuration object
var WaitListConfig WaitListConfiguration
