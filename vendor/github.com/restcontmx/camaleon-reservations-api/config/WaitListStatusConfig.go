package config

import (
	"fmt"

	"github.com/graphql-go/graphql"
	"github.com/restcontmx/camaleon-reservations-api/app/data"
	"github.com/restcontmx/camaleon-reservations-api/app/models"
)

// WaitListStatusConfiguration is the configuration struct
type WaitListStatusConfiguration struct {
	Repository *data.WaitListStatusRepository
}

// WaitListStatusConfig configuration object
var WaitListStatusConfig WaitListStatusConfiguration

var WaitListStatusObject = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "WaitListStatus",
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
// GetAllWaitListStatus returns all the objects
//
var GetAllWaitListStatus = &graphql.Field{
	Type: graphql.NewList(WaitListStatusObject),
	Args: graphql.FieldConfigArgument{},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		isOk, _ := ValidateAuthentication(p.Context.Value(authKey).(string))
		if isOk {
			return WaitListStatusConfig.Repository.GetAll()
		}
		return nil, fmt.Errorf("Invalid Credentials")
	},
}

//
// CreateUpdateWaitListStatus graphql configuration
//
var CreateUpdateWaitListStatus = &graphql.Field{
	Type: WaitListStatusObject,
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
			WaitListStatusConfig.Repository.Model = models.WaitListStatusModel{}
			if isOk {
				entity, _ := WaitListStatusConfig.Repository.GetByID(id)
				WaitListStatusConfig.Repository.Model = entity.(models.WaitListStatusModel)

				if value, isOk := p.Args["value"].(int); isOk {
					WaitListStatusConfig.Repository.Model.Value = value
				}
				if description, isOk := p.Args["description"].(string); isOk {
					WaitListStatusConfig.Repository.Model.Description = description
				}
				return WaitListStatusConfig.Repository.Update()
			}
			if description, isOk := p.Args["description"].(string); isOk {
				WaitListStatusConfig.Repository.Model.Description = description
				if value, isOk := p.Args["value"].(int); isOk {
					WaitListStatusConfig.Repository.Model.Value = value
					return WaitListStatusConfig.Repository.Create()
				}
				return nil, fmt.Errorf("You must provide a Value ")
			}
			return nil, fmt.Errorf("You must provide a Description ")
		}
		return nil, fmt.Errorf("Invalid Credentials")
	},
}
