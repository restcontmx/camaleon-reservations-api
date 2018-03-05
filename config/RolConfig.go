package config

import (
	"fmt"

	"github.com/graphql-go/graphql"
	"github.com/restcontmx/camaleon-reservations-api/app/data"
	"github.com/restcontmx/camaleon-reservations-api/app/models"
)

//
// RolConfiguration is the main configuration graphql object that contains the repository
// @param Repository : RolRepository
//
type RolConfiguration struct {
	Repository *data.RolRepository
}

// RolConfig main object
var RolConfig RolConfiguration

//
// RolObject is the graphql object of the rol model
//
var RolObject = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Rol",
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
// CreateUpdateRol graphql configuration
//
var CreateUpdateRol = &graphql.Field{
	Type: RolObject,
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
			RolConfig.Repository.Model = models.RolModel{}
			if isOk {
				entity, _ := RolConfig.Repository.GetByID(id)
				RolConfig.Repository.Model = entity.(models.RolModel)

				if value, isOk := p.Args["value"].(int); isOk {
					RolConfig.Repository.Model.Value = value
				}
				if description, isOk := p.Args["description"].(string); isOk {
					RolConfig.Repository.Model.Description = description
				}
				return RolConfig.Repository.Update()
			}
			if value, isOk := p.Args["value"].(int); isOk {
				RolConfig.Repository.Model.Value = value
			}
			if description, isOk := p.Args["description"].(string); isOk {
				RolConfig.Repository.Model.Description = description
				return RolConfig.Repository.Create()
			}
			return nil, fmt.Errorf("You must provide a Description ")

		}
		return nil, fmt.Errorf("Invalid Credentials")
	},
}

//
// GetAllRol graphql configuration
//
var GetAllRol = &graphql.Field{
	Type: graphql.NewList(RolObject),
	Args: graphql.FieldConfigArgument{},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		isOk, _ := ValidateAuthentication(p.Context.Value(authKey).(string))
		if isOk {
			return RolConfig.Repository.GetAll()
		}
		return nil, fmt.Errorf("Invalid Credentials")
	},
}

//
// RetrieveRol graphql configuration
//
var RetrieveRol = &graphql.Field{
	Type: RolObject,
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		isOk, _ := ValidateAuthentication(p.Context.Value(authKey).(string))
		if isOk {
			if id, isOk := p.Args["id"].(int); isOk {
				return RolConfig.Repository.GetByID(id)
			}
			return nil, fmt.Errorf("There is no id field")
		}
		return nil, fmt.Errorf("Invalid Credentials")
	},
}
