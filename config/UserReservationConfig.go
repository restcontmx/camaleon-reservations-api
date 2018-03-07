package config

import (
	"fmt"

	"github.com/graphql-go/graphql"
	"github.com/restcontmx/camaleon-reservations-api/app/data"
	"github.com/restcontmx/camaleon-reservations-api/app/models"
)

// UserReservationConfiguration user configuration object
type UserReservationConfiguration struct {
	Repository *data.UserReservationRepository
}

// UserReservationConfig main variable
var UserReservationConfig UserReservationConfiguration

//
// UserReservationObject is the graphql object of the client info
//
var UserReservationObject = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "UserReservation",
		Fields: graphql.Fields{
			"ID": &graphql.Field{
				Type: graphql.String,
			},
			"User": &graphql.Field{
				Type: UserObject,
			},
			"Business": &graphql.Field{
				Type: BusinessObject,
			},
			"AllowedLocations": &graphql.Field{
				Type: graphql.NewList(LocationObject),
			},
			"Rol": &graphql.Field{
				Type: RolObject,
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
// GetAllUserReservations returns all the areas
//
var GetAllUserReservations = &graphql.Field{
	Type: graphql.NewList(UserReservationObject),
	Args: graphql.FieldConfigArgument{},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		isOk, _ := ValidateAuthentication(p.Context.Value(authKey).(string))
		if isOk {
			return UserReservationConfig.Repository.GetAll()
		}
		return nil, fmt.Errorf("Invalid Credentials")
	},
}

//
// CreateUpdateUserReservation graphql configuration
//
var CreateUpdateUserReservation = &graphql.Field{
	Type: UserReservationObject,
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"user_id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"business_id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"allowed_locations": &graphql.ArgumentConfig{
			Type: graphql.NewList(graphql.Int),
		},
		"rol_id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		isOk, _ := ValidateAuthentication(p.Context.Value(authKey).(string))
		if isOk {
			id, isOk := p.Args["id"].(int)
			UserReservationConfig.Repository.Model = models.UserReservationModel{}
			if isOk {
				entity, _ := UserReservationConfig.Repository.GetByID(id)
				UserReservationConfig.Repository.Model = entity.(models.UserReservationModel)

				if rolID, isOk := p.Args["rol_id"].(int); isOk {
					UserReservationConfig.Repository.Model.Rol.ID = rolID
				}
				if businessID, isOk := p.Args["business_id"].(int); isOk {
					UserReservationConfig.Repository.Model.Business.ID = businessID
				}
				if userID, isOk := p.Args["user_id"].(int); isOk {
					UserReservationConfig.Repository.Model.User.ID = userID
				}
				return UserReservationConfig.Repository.Update()
			}
			if businessID, isOk := p.Args["business_id"].(int); isOk {
				UserReservationConfig.Repository.Model.Business.ID = businessID
				if userID, isOk := p.Args["user_id"].(int); isOk {
					UserReservationConfig.Repository.Model.User.ID = userID
					if rolID, isOk := p.Args["rol_id"].(int); isOk {
						UserReservationConfig.Repository.Model.Rol.ID = rolID
						userReservation, err := UserReservationConfig.Repository.Create()
						if err == nil {
							UserReservationConfig.Repository.Model = userReservation.(models.UserReservationModel)
							if allowedLocationsID, isOk := p.Args["allowed_locations"].([]interface{}); isOk {
								isOk, err := UserReservationConfig.Repository.RemoveAllAllowedLocations()
								if isOk {
									for _, obj := range allowedLocationsID {
										_, err := UserReservationConfig.Repository.AddAllowedLocation(obj.(int))
										if err != nil {
											return nil, err
										}
									}
								} else {
									return nil, err
								}
							}
							return UserReservationConfig.Repository.GetByID(UserReservationConfig.Repository.Model.ID)
						}
						return nil, err
					}
					return nil, fmt.Errorf("You must provide a Rol ID")
				}
				return nil, fmt.Errorf("You must provide a User ID")
			}
			return nil, fmt.Errorf("You must provide a Business ID")
		}
		return nil, fmt.Errorf("Invalid Credentials")
	},
}

//
// RetrieveUserReservation graphql configuration
//
var RetrieveUserReservation = &graphql.Field{
	Type: UserReservationObject,
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"user_id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		isOk, _ := ValidateAuthentication(p.Context.Value(authKey).(string))
		if isOk {
			if id, isOk := p.Args["id"].(int); isOk {
				return UserReservationConfig.Repository.GetByID(id)
			}
			if userID, isOk := p.Args["user_id"].(int); isOk {
				return UserReservationConfig.Repository.GetByUserID(userID)
			}
			return nil, fmt.Errorf("There is no id or user id field")
		}
		return nil, fmt.Errorf("Invalid Credentials")
	},
}
