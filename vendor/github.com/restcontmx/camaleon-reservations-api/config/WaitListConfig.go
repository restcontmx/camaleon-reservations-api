package config

import (
	"fmt"
	"time"

	"github.com/restcontmx/camaleon-reservations-api/app/helpers"
	"github.com/restcontmx/camaleon-reservations-api/app/models"

	"github.com/graphql-go/graphql"
	"github.com/restcontmx/camaleon-reservations-api/app/data"
)

// WaitListConfiguration is the configuration struct
type WaitListConfiguration struct {
	Repository *data.WaitListRepository
}

// WaitListConfig configuration object
var WaitListConfig WaitListConfiguration

//
// WaitListObject reservation graphql configuration
//
var WaitListObject = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "WaitList",
		Fields: graphql.Fields{
			"ID": &graphql.Field{
				Type: graphql.Int,
			},
			"Status": &graphql.Field{
				Type: WaitListStatusObject,
			},
			"ClientInfo": &graphql.Field{
				Type: ClientInfoObject,
			},
			"Location": &graphql.Field{
				Type: LocationObject,
			},
			"AlertTime": &graphql.Field{
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
// GetAllWaitLists graphql configuration
//
var GetAllWaitLists = &graphql.Field{
	Type: graphql.NewList(WaitListObject),
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
									return WaitListConfig.Repository.GetByDates(location, statusID, t1, t2)
								}
								return WaitListConfig.Repository.GetByDates(location, 0, t1, t2)
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
// CreateUpdateWaitList will create a reservation
//
var CreateUpdateWaitList = &graphql.Field{
	Type: ReservationObject,
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
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
		"alert_time": &graphql.ArgumentConfig{
			Type: graphql.DateTime,
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		isOk, _ := ValidateAuthentication(p.Context.Value(authKey).(string))
		if isOk {
			id, isOk := p.Args["id"].(int)
			WaitListConfig.Repository.Model = models.WaitListModel{}
			if isOk {
				entity, _ := WaitListConfig.Repository.GetByID(id)
				WaitListConfig.Repository.Model = entity.(models.WaitListModel)

				if statusID, isOk := p.Args["status_id"].(int); isOk {
					WaitListConfig.Repository.Model.Status.ID = statusID
				}
				if timeLimit, isOk := p.Args["time_limit"].(int); isOk {
					WaitListConfig.Repository.Model.TimeLimit = timeLimit
				}
				if guests, isOk := p.Args["guests"].(int); isOk {
					WaitListConfig.Repository.Model.Guests = guests
				}
				if clientInfoID, isOk := p.Args["client_info_id"].(int); isOk {
					WaitListConfig.Repository.Model.ClientInfo.ID = clientInfoID
				}
				if alertTime, isOk := p.Args["alert_time"].(string); isOk {
					layout := "01/02/2006 15:04:05"
					if t, err := time.Parse(layout, alertTime); err == nil {
						WaitListConfig.Repository.Model.AlertTime = t
					} else {
						return nil, fmt.Errorf("Format for 'AlertTime' is 'mm/dd/yyyy hh:mm:ss'")
					}
				}
				return WaitListConfig.Repository.Update()
			}

			if statusID, isOk := p.Args["status_id"].(int); isOk {
				WaitListConfig.Repository.Model.Status.ID = statusID
				if timeLimit, isOk := p.Args["time_limit"].(int); isOk {
					WaitListConfig.Repository.Model.TimeLimit = timeLimit
					if location, isOk := p.Args["location"].(int); isOk {
						WaitListConfig.Repository.Model.Location.ID = location
						if guests, isOk := p.Args["guests"].(int); isOk {
							WaitListConfig.Repository.Model.Guests = guests
							if alertTime, isOk := p.Args["alert_time"].(string); isOk {
								layout := "01/02/2006 15:04:05"
								if t, err := time.Parse(layout, alertTime); err == nil {
									WaitListConfig.Repository.Model.AlertTime = t
									if clientInfoID, isOk := p.Args["client_info_id"].(int); isOk {

										WaitListConfig.Repository.Model.ClientInfo.ID = clientInfoID
										object, err := WaitListConfig.Repository.Create()

										if err != nil {
											return nil, err
										}

										_, _ = helpers.SendConfirmationEmailWaitList(object.(models.WaitListModel))

										return object, nil
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
			return nil, fmt.Errorf("You must provide a Status ")
		}
		return nil, fmt.Errorf("Invalid Credentials")
	},
}

//
// RetrieveWaitList graphql configuration
//
var RetrieveWaitList = &graphql.Field{
	Type: WaitListObject,
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		isOk, _ := ValidateAuthentication(p.Context.Value(authKey).(string))
		if isOk {
			if id, isOk := p.Args["id"].(int); isOk {
				return WaitListConfig.Repository.GetByID(id)
			}
			return nil, fmt.Errorf("There is no id field")
		}
		return nil, fmt.Errorf("Invalid Credentials")
	},
}
