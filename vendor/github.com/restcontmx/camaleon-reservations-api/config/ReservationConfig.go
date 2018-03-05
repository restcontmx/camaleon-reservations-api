package config

import (
	"fmt"
	"time"

	"github.com/restcontmx/camaleon-reservations-api/app/helpers"

	"github.com/restcontmx/camaleon-reservations-api/app/data"
	"github.com/restcontmx/camaleon-reservations-api/app/models"

	"github.com/graphql-go/graphql"
)

// ReservationConfiguration reservation configuration
type ReservationConfiguration struct {
	Repository *data.ReservationRepository
}

// ReservationConfig object
var ReservationConfig ReservationConfiguration

//
// ReservationObject reservation graphql configuration
//
var ReservationObject = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Reservation",
		Fields: graphql.Fields{
			"ID": &graphql.Field{
				Type: graphql.Int,
			},
			"UID": &graphql.Field{
				Type: graphql.String,
			},
			"Table": &graphql.Field{
				Type: TableObject,
			},
			"Status": &graphql.Field{
				Type: ReservationStatusObject,
			},
			"ClientInfo": &graphql.Field{
				Type: ClientInfoObject,
			},
			"Location": &graphql.Field{
				Type: LocationObject,
			},
			"Date": &graphql.Field{
				Type: graphql.DateTime,
			},
			"TimeLimit": &graphql.Field{
				Type: graphql.Int,
			},
			"Guests": &graphql.Field{
				Type: graphql.Int,
			},
			"Timestamp": &graphql.Field{
				Type: graphql.DateTime,
			},
			"Updated": &graphql.Field{
				Type: graphql.DateTime,
			},
		},
	},
)

//
// GetAllReservations graphql configuration
//
var GetAllReservations = &graphql.Field{
	Type: graphql.NewList(ReservationObject),
	Args: graphql.FieldConfigArgument{
		"location": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
		"start_date": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.DateTime),
		},
		"end_date": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.DateTime),
		},
		"status_id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		isOk, _ := ValidateAuthentication(p.Context.Value(authKey).(string))
		if isOk {
			layout := "01/02/2006 15:04:05"
			if location, isOk := p.Args["location"].(int); isOk {
				if startDate, isOk := p.Args["start_date"].(string); isOk {
					if t1, err := time.Parse(layout, startDate); err == nil {
						if endDate, isOk := p.Args["end_date"].(string); isOk {
							if t2, err := time.Parse(layout, endDate); err == nil {
								if statusID, isOk := p.Args["status_id"].(int); isOk {
									return ReservationConfig.Repository.GetByDates(location, statusID, t1, t2)
								}
								return ReservationConfig.Repository.GetByDates(location, 0, t1, t2)
							}
							return nil, fmt.Errorf("Format for 'end_date' is 'mm/dd/yyyy hh:mm:ss'")
						}
						return nil, fmt.Errorf("You must provide a 'end_date' variable")
					}
					return nil, fmt.Errorf("Format for 'start date' is 'mm/dd/yyyy hh:mm:ss'")
				}
				return nil, fmt.Errorf("You must provide a 'start_date' variable")
			}
			return nil, fmt.Errorf("You must provide a 'location' variable")
		}
		return nil, fmt.Errorf("Invalid Credentials")
	},
}

//
// CreateUpdateReservation will create a reservation
//
var CreateUpdateReservation = &graphql.Field{
	Type: ReservationObject,
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"table_id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"location": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"status_id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"client_info_id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"time_limit": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"guests": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"date": &graphql.ArgumentConfig{
			Type: graphql.DateTime,
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		isOk, _ := ValidateAuthentication(p.Context.Value(authKey).(string))
		if isOk {
			id, isOk := p.Args["id"].(int)
			ReservationConfig.Repository.Model = models.ReservationModel{}
			if isOk {
				entity, _ := ReservationConfig.Repository.GetByID(id)
				ReservationConfig.Repository.Model = entity.(models.ReservationModel)

				if statusID, isOk := p.Args["status_id"].(int); isOk {
					ReservationConfig.Repository.Model.Status.ID = statusID
				}
				if tableID, isOk := p.Args["table_id"].(int); isOk {
					ReservationConfig.Repository.Model.Table.ID = tableID
				}
				if timeLimit, isOk := p.Args["time_limit"].(int); isOk {
					ReservationConfig.Repository.Model.TimeLimit = timeLimit
				}
				if guests, isOk := p.Args["guests"].(int); isOk {
					ReservationConfig.Repository.Model.Guests = guests
				}
				if clientInfoID, isOk := p.Args["client_info_id"].(int); isOk {
					ReservationConfig.Repository.Model.ClientInfo.ID = clientInfoID
				}
				if date, isOk := p.Args["date"].(string); isOk {
					layout := "01/02/2006 15:04:05"
					if t, err := time.Parse(layout, date); err == nil {
						ReservationConfig.Repository.Model.Date = t
					} else {
						return nil, fmt.Errorf("Format for 'Date' is 'mm/dd/yyyy hh:mm:ss'")
					}
				}
				return ReservationConfig.Repository.Update()
			}

			if statusID, isOk := p.Args["status_id"].(int); isOk {
				ReservationConfig.Repository.Model.Status.ID = statusID
				if tableID, isOk := p.Args["table_id"].(int); isOk {
					ReservationConfig.Repository.Model.Table.ID = tableID
					if timeLimit, isOk := p.Args["time_limit"].(int); isOk {
						ReservationConfig.Repository.Model.TimeLimit = timeLimit
						if location, isOk := p.Args["location"].(int); isOk {
							ReservationConfig.Repository.Model.Location.ID = location
							if guests, isOk := p.Args["guests"].(int); isOk {
								ReservationConfig.Repository.Model.Guests = guests
								if date, isOk := p.Args["date"].(string); isOk {
									layout := "01/02/2006 15:04:05"
									if t, err := time.Parse(layout, date); err == nil {
										ReservationConfig.Repository.Model.Date = t
										if clientInfoID, isOk := p.Args["client_info_id"].(int); isOk {

											ReservationConfig.Repository.Model.ClientInfo.ID = clientInfoID
											reservation, err := ReservationConfig.Repository.Create()

											if err != nil {
												return nil, err
											}

											_, _ = helpers.SendConfirmationEmail(reservation.(models.ReservationModel))

											return reservation, nil
										}
										return nil, fmt.Errorf("You must provide a Client Info ID ")
									}
									return nil, fmt.Errorf("Format for 'Date' is 'mm/dd/yyyy hh:mm:ss'")
								}
								return nil, fmt.Errorf("You must provide a Date ")
							}
							return nil, fmt.Errorf("You must provide Guests ")
						}
						return nil, fmt.Errorf("You must provide a Location ")
					}
					return nil, fmt.Errorf("You must provide a Time Limit ")
				}
				return nil, fmt.Errorf("You must provide a Table ")
			}
			return nil, fmt.Errorf("You must provide a Status ")
		}
		return nil, fmt.Errorf("Invalid Credentials")
	},
}

//
// RetrieveReservation graphql configuration
//
var RetrieveReservation = &graphql.Field{
	Type: ReservationObject,
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"uid": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		isOk, _ := ValidateAuthentication(p.Context.Value(authKey).(string))
		if isOk {
			if id, isOk := p.Args["id"].(int); isOk {
				return ReservationConfig.Repository.GetByID(id)
			}
			return nil, fmt.Errorf("There is no id field")
		}
		return nil, fmt.Errorf("Invalid Credentials")
	},
}
