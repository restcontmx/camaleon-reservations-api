package config

import (
	"fmt"

	"github.com/graphql-go/graphql"
	"github.com/restcontmx/camaleon-reservations-api/app/data"
	"github.com/restcontmx/camaleon-reservations-api/app/models"
)

// ReservationStatusConfiguration reservation status configuration
type ReservationStatusConfiguration struct {
	Repository *data.ReservationStatusRepository
}

// ReservationStatusConfig object
var ReservationStatusConfig ReservationStatusConfiguration

//
// ReservationStatusObject …
//
var ReservationStatusObject = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "ReservationStatus",
		Fields: graphql.Fields{
			"ID": &graphql.Field{
				Type: graphql.Int,
			},
			"Value": &graphql.Field{
				Type: graphql.Int,
			},
			"Description": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

//
// GetAllReservationStatus returns all the objects
//
var GetAllReservationStatus = &graphql.Field{
	Type: graphql.NewList(ReservationStatusObject),
	Args: graphql.FieldConfigArgument{},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		isOk, _ := ValidateAuthentication(p.Context.Value(authKey).(string))
		if isOk {
			return ReservationStatusConfig.Repository.GetAll()
		}
		return nil, fmt.Errorf("Invalid Credentials")
	},
}

//
// CreateUpdateReservationStatus graphql configuration
//
var CreateUpdateReservationStatus = &graphql.Field{
	Type: ReservationStatusObject,
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"value": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"description": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		isOk, _ := ValidateAuthentication(p.Context.Value(authKey).(string))
		if isOk {
			id, isOk := p.Args["id"].(int)
			ReservationStatusConfig.Repository.Model = models.ReservationStatusModel{}
			if isOk {
				entity, _ := ReservationStatusConfig.Repository.GetByID(id)
				ReservationStatusConfig.Repository.Model = entity.(models.ReservationStatusModel)

				if value, isOk := p.Args["value"].(int); isOk {
					ReservationStatusConfig.Repository.Model.Value = value
				}
				if description, isOk := p.Args["description"].(string); isOk {
					ReservationStatusConfig.Repository.Model.Description = description
				}
				return ReservationStatusConfig.Repository.Update()
			}
			if description, isOk := p.Args["description"].(string); isOk {
				ReservationStatusConfig.Repository.Model.Description = description
				if value, isOk := p.Args["value"].(int); isOk {
					ReservationStatusConfig.Repository.Model.Value = value
					return ReservationStatusConfig.Repository.Create()
				}
				return nil, fmt.Errorf("You must provide a Value ")
			}
			return nil, fmt.Errorf("You must provide a Description ")
		}
		return nil, fmt.Errorf("Invalid Credentials")
	},
}
