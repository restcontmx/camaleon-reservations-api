package config

import (
	"fmt"

	"github.com/graphql-go/graphql"
	"github.com/restcontmx/camaleon-reservations-api/app/data"
	"github.com/restcontmx/camaleon-reservations-api/app/models"
)

//
// LocationConfiguration is the main configuration object defintion
// @param Repository : Location Repository - the configuration repository
//
type LocationConfiguration struct {
	Repository *data.LocationRepository
}

// LocationConfig : location configuration Main Object
var LocationConfig LocationConfiguration

//
// LocationObject location graphql configuration
//
var LocationObject = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Location",
		Fields: graphql.Fields{
			"ID": &graphql.Field{
				Type: graphql.Int,
			},
			"BusinessID": &graphql.Field{
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
			"Address1": &graphql.Field{
				Type: graphql.String,
			},
			"Address2": &graphql.Field{
				Type: graphql.String,
			},
			"Phone": &graphql.Field{
				Type: graphql.String,
			},
			"City": &graphql.Field{
				Type: graphql.String,
			},
			"State": &graphql.Field{
				Type: graphql.String,
			},
			"Country": &graphql.Field{
				Type: graphql.String,
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
// CreateUpdateLocation graqphl configuration
//
var CreateUpdateLocation = &graphql.Field{
	Type: LocationObject,
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"business_id": &graphql.ArgumentConfig{
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
		"address1": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"address2": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"phone": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"city": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"state": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"country": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		isOk, _ := ValidateAuthentication(p.Context.Value(authKey).(string))
		if isOk {
			id, isOk := p.Args["id"].(int)
			LocationConfig.Repository.Model = models.LocationModel{}
			if isOk {
				locationEntity, _ := LocationConfig.Repository.GetByID(id)
				LocationConfig.Repository.Model = locationEntity.(models.LocationModel)

				if name, isOk := p.Args["name"].(string); isOk {
					LocationConfig.Repository.Model.Name = name
				}
				if crcenterid, isOk := p.Args["crcenter_id"].(int); isOk {
					LocationConfig.Repository.Model.CrCenterID = crcenterid
				}
				if description, isOk := p.Args["description"].(string); isOk {
					LocationConfig.Repository.Model.Description = description
				}
				if address1, isOk := p.Args["address1"].(string); isOk {
					LocationConfig.Repository.Model.Address1 = address1
				}
				if address2, isOk := p.Args["address2"].(string); isOk {
					LocationConfig.Repository.Model.Address2 = address2
				}
				if phone, isOk := p.Args["phone"].(string); isOk {
					LocationConfig.Repository.Model.Phone = phone
				}
				if city, isOk := p.Args["city"].(string); isOk {
					LocationConfig.Repository.Model.City = city
				}
				if state, isOk := p.Args["state"].(string); isOk {
					LocationConfig.Repository.Model.State = state
				}
				if country, isOk := p.Args["country"].(string); isOk {
					LocationConfig.Repository.Model.Country = country
				}
				return LocationConfig.Repository.Update()
			}

			if description, isOk := p.Args["description"].(string); isOk {
				LocationConfig.Repository.Model.Description = description
			}
			if address1, isOk := p.Args["address1"].(string); isOk {
				LocationConfig.Repository.Model.Address1 = address1
			}
			if address2, isOk := p.Args["address2"].(string); isOk {
				LocationConfig.Repository.Model.Address2 = address2
			}
			if phone, isOk := p.Args["phone"].(string); isOk {
				LocationConfig.Repository.Model.Phone = phone
			}
			if city, isOk := p.Args["city"].(string); isOk {
				LocationConfig.Repository.Model.City = city
			}
			if state, isOk := p.Args["state"].(string); isOk {
				LocationConfig.Repository.Model.State = state
			}
			if country, isOk := p.Args["country"].(string); isOk {
				LocationConfig.Repository.Model.Country = country
			}

			if name, isOk := p.Args["name"].(string); isOk {
				LocationConfig.Repository.Model.Name = name
				if businessid, isOk := p.Args["business_id"].(int); isOk {
					LocationConfig.Repository.Model.BusinessID = businessid
					if crcenterid, isOk := p.Args["crcenter_id"].(int); isOk {
						LocationConfig.Repository.Model.CrCenterID = crcenterid
						return LocationConfig.Repository.Create()
					}
					return nil, fmt.Errorf("You must provide Cr Center ID")
				}
				return nil, fmt.Errorf("You must provide Business ID")
			}
			return nil, fmt.Errorf("You must provide a Name ")
		}
		return nil, fmt.Errorf("Invalid Credentials")
	},
}

//
// GetAllLocations graphql configuration
//
var GetAllLocations = &graphql.Field{
	Type: graphql.NewList(LocationObject),
	Args: graphql.FieldConfigArgument{},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		isOk, _ := ValidateAuthentication(p.Context.Value(authKey).(string))
		if isOk {
			return LocationConfig.Repository.GetAll()
		}
		return nil, fmt.Errorf("Invalid Credentials")
	},
}

//
// RetrieveLocation graphql configuration
//
var RetrieveLocation = &graphql.Field{
	Type: LocationObject,
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
				return LocationConfig.Repository.GetByID(id)
			}
			if crcenterid, isOk := p.Args["crcenter_id"].(int); isOk {
				return LocationConfig.Repository.GetByCrCenterID(crcenterid)
			}
			return nil, fmt.Errorf("There is no id field")
		}
		return nil, fmt.Errorf("Invalid Credentials")
	},
}

//
// DeleteLocation graphql configuration
//
var DeleteLocation = &graphql.Field{
	Type: graphql.NewList(LocationObject),
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
				LocationConfig.Repository.Model = models.LocationModel{ID: id}
				if isOk, _ := LocationConfig.Repository.Delete(); isOk {
					return nil, nil
				}
				return nil, fmt.Errorf("There object was not deleted")
			}
			return nil, fmt.Errorf("No id found")
		}
		return nil, fmt.Errorf("Invalid Credentials")
	},
}
