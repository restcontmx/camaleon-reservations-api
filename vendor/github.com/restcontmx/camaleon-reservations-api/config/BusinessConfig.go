package config

import (
	"fmt"

	"github.com/graphql-go/graphql"
	"github.com/restcontmx/camaleon-reservations-api/app/data"
	"github.com/restcontmx/camaleon-reservations-api/app/models"
)

//
// BusinessConfiguration is the main configuration graphql object that contains the repository
// @param Repository : BusinessRepository
//
type BusinessConfiguration struct {
	Repository *data.BusinessRepository
}

// BusinessConfig main object
var BusinessConfig BusinessConfiguration

//
// BusinessObject is the graphql object of the business model
//
var BusinessObject = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Business",
		Fields: graphql.Fields{
			"ID": &graphql.Field{
				Type: graphql.Int,
			},
			"CrCenterID": &graphql.Field{
				Type: graphql.Int,
			},
			"Name": &graphql.Field{
				Type: graphql.String,
			},
			"Description": &graphql.Field{
				Type: graphql.String,
			},
			"Locations": &graphql.Field{
				Type: graphql.NewList(LocationObject),
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
// CreateUpdateBusiness graphql configuration
//
var CreateUpdateBusiness = &graphql.Field{
	Type: BusinessObject,
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"crcenter_id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"name": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"description": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		isOk, _ := ValidateAuthentication(p.Context.Value(authKey).(string))
		if isOk {
			id, isOk := p.Args["id"].(int)
			BusinessConfig.Repository.Model = models.BusinessModel{}
			if isOk {
				businessEntity, _ := BusinessConfig.Repository.GetByID(id)
				BusinessConfig.Repository.Model = businessEntity.(models.BusinessModel)

				if name, isOk := p.Args["name"].(string); isOk {
					BusinessConfig.Repository.Model.Name = name
				}
				if crcenterid, isOk := p.Args["crcenter_id"].(int); isOk {
					BusinessConfig.Repository.Model.CrCenterID = crcenterid
				}
				if description, isOk := p.Args["description"].(string); isOk {
					BusinessConfig.Repository.Model.Description = description
				}
				return BusinessConfig.Repository.Update()
			}
			if description, isOk := p.Args["description"].(string); isOk {
				BusinessConfig.Repository.Model.Description = description
			}

			if crcenterid, isOk := p.Args["crcenter_id"].(int); isOk {
				BusinessConfig.Repository.Model.CrCenterID = crcenterid
			}

			if name, isOk := p.Args["name"].(string); isOk {
				BusinessConfig.Repository.Model.Name = name
				return BusinessConfig.Repository.Create()
			}
			return nil, fmt.Errorf("You must provide a Name ")

		}
		return nil, fmt.Errorf("Invalid Credentials")
	},
}

//
// GetAllBusinesses graphql configuration
//
var GetAllBusinesses = &graphql.Field{
	Type: graphql.NewList(BusinessObject),
	Args: graphql.FieldConfigArgument{},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		isOk, _ := ValidateAuthentication(p.Context.Value(authKey).(string))
		if isOk {
			return BusinessConfig.Repository.GetAll()
		}
		return nil, fmt.Errorf("Invalid Credentials")
	},
}

//
// RetrieveBusiness graphql configuration
//
var RetrieveBusiness = &graphql.Field{
	Type: BusinessObject,
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"crcenter_id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		isOk, _ := ValidateAuthentication(p.Context.Value(authKey).(string))
		if isOk {
			if id, isOk := p.Args["id"].(int); isOk {
				return BusinessConfig.Repository.GetByID(id)
			}
			if crcenterid, isOk := p.Args["crcenter_id"].(int); isOk {
				return BusinessConfig.Repository.GetByCrCenterID(crcenterid)
			}
			return nil, fmt.Errorf("There is no id field")
		}
		return nil, fmt.Errorf("Invalid Credentials")
	},
}

//
// DeleteBusiness graphql configuration
//
var DeleteBusiness = &graphql.Field{
	Type: graphql.NewList(BusinessObject),
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		isOk, _ := ValidateAuthentication(p.Context.Value(authKey).(string))
		if isOk {
			id, isOk := p.Args["id"].(int)
			if isOk {
				BusinessConfig.Repository.Model = models.BusinessModel{ID: id}
				if isOk, _ := BusinessConfig.Repository.Delete(); isOk {
					return nil, nil
				}
				return nil, fmt.Errorf("There object was not deleted")
			}
			return nil, fmt.Errorf("No id found")
		}
		return nil, fmt.Errorf("Invalid Credentials")
	},
}
