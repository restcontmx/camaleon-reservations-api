package config

import (
	"fmt"

	"github.com/graphql-go/graphql"
	"github.com/restcontmx/camaleon-reservations-api/app/data"
	"github.com/restcontmx/camaleon-reservations-api/app/models"
)

// AreaConfiguration area configuration
type AreaConfiguration struct {
	Repository *data.AreaRepository
}

// AreaConfig object
var AreaConfig AreaConfiguration

//
// AreaObject graphql configuration
//
var AreaObject = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Area",
		Fields: graphql.Fields{
			"ID": &graphql.Field{
				Type: graphql.Int,
			},
			"Name": &graphql.Field{
				Type: graphql.String,
			},
			"Description": &graphql.Field{
				Type: graphql.String,
			},
			"Tables": &graphql.Field{
				Type: graphql.NewList(TableObject),
			},
			"Location": &graphql.Field{
				Type: LocationObject,
			},
		},
	},
)

//
// GetAllAreas returns all the areas
//
var GetAllAreas = &graphql.Field{
	Type: graphql.NewList(AreaObject),
	Args: graphql.FieldConfigArgument{
		"location": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		isOk, _ := ValidateAuthentication(p.Context.Value(authKey).(string))
		if isOk {
			locationID, isOk := p.Args["location"].(int)
			if isOk {
				return AreaConfig.Repository.GetAllByLocation(locationID)
			}
			return AreaConfig.Repository.GetAll()
		}
		return nil, fmt.Errorf("Invalid Credentials")
	},
}

//
// CreateUpdateArea graphql configuration
//
var CreateUpdateArea = &graphql.Field{
	Type: AreaObject,
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"name": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"description": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"location": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"img_url": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		isOk, _ := ValidateAuthentication(p.Context.Value(authKey).(string))
		if isOk {
			id, isOk := p.Args["id"].(int)
			AreaConfig.Repository.Model = models.AreaModel{}
			if isOk {
				entity, _ := AreaConfig.Repository.GetByID(id)
				AreaConfig.Repository.Model = entity.(models.AreaModel)

				if name, isOk := p.Args["name"].(string); isOk {
					AreaConfig.Repository.Model.Name = name
				}
				if description, isOk := p.Args["description"].(string); isOk {
					AreaConfig.Repository.Model.Description = description
				}
				if imgurl, isOk := p.Args["img_url"].(string); isOk {
					AreaConfig.Repository.Model.ImgURL = imgurl
				}
				return AreaConfig.Repository.Update()
			}

			if description, isOk := p.Args["description"].(string); isOk {
				AreaConfig.Repository.Model.Description = description
			}
			if imgurl, isOk := p.Args["img_url"].(string); isOk {
				AreaConfig.Repository.Model.ImgURL = imgurl
			}

			if name, isOk := p.Args["name"].(string); isOk {
				AreaConfig.Repository.Model.Name = name
				if location, isOk := p.Args["location"].(int); isOk {
					AreaConfig.Repository.Model.Location.ID = location
					return AreaConfig.Repository.Create()
				}
				return nil, fmt.Errorf("You must provide a Location ")
			}
			return nil, fmt.Errorf("You must provide a Name ")

		}
		return nil, fmt.Errorf("Invalid Credentials")
	},
}

//
// RetrieveArea graphql configuration
//
var RetrieveArea = &graphql.Field{
	Type: AreaObject,
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		isOk, _ := ValidateAuthentication(p.Context.Value(authKey).(string))
		if isOk {
			if id, isOk := p.Args["id"].(int); isOk {
				return AreaConfig.Repository.GetByID(id)
			}
			return nil, fmt.Errorf("There is no id field")
		}
		return nil, fmt.Errorf("Invalid Credentials")
	},
}
