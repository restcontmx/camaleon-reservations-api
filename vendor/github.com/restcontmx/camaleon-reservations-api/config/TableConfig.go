package config

import (
	"fmt"

	"github.com/graphql-go/graphql"
	"github.com/restcontmx/camaleon-reservations-api/app/data"
	"github.com/restcontmx/camaleon-reservations-api/app/models"
)

// TableConfiguration table configuration object
type TableConfiguration struct {
	Repository *data.TableRepository
}

// TableConfig object
var TableConfig TableConfiguration

//
// TableObject graphql configuration
//
var TableObject = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Table",
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
			"ImgURL": &graphql.Field{
				Type: graphql.String,
			},
			"MaxGuests": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

//
// GetAllTables returns all the tables
//
var GetAllTables = &graphql.Field{
	Type: graphql.NewList(TableObject),
	Args: graphql.FieldConfigArgument{
		"area_id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		isOk, _ := ValidateAuthentication(p.Context.Value(authKey).(string))
		if isOk {
			areaID, isOk := p.Args["area_id"].(int)
			if isOk {
				return TableConfig.Repository.GetAllByAreaID(areaID)
			}
			return TableConfig.Repository.GetAll()
		}
		return nil, fmt.Errorf("Invalid Credentials")
	},
}

//
// CreateUpdateTable graphql configuration
//
var CreateUpdateTable = &graphql.Field{
	Type: TableObject,
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
		"max_guests": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"img_url": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"area_id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		isOk, _ := ValidateAuthentication(p.Context.Value(authKey).(string))
		if isOk {
			id, isOk := p.Args["id"].(int)
			TableConfig.Repository.Model = models.TableModel{}
			if isOk {
				entity, _ := TableConfig.Repository.GetByID(id)
				TableConfig.Repository.Model = entity.(models.TableModel)

				if name, isOk := p.Args["name"].(string); isOk {
					TableConfig.Repository.Model.Name = name
				}
				if description, isOk := p.Args["description"].(string); isOk {
					TableConfig.Repository.Model.Description = description
				}
				if imgurl, isOk := p.Args["img_url"].(string); isOk {
					TableConfig.Repository.Model.ImgURL = imgurl
				}
				if maxGuests, isOk := p.Args["max_guests"].(int); isOk {
					TableConfig.Repository.Model.MaxGuests = maxGuests
				}
				return TableConfig.Repository.Update()
			}

			if description, isOk := p.Args["description"].(string); isOk {
				TableConfig.Repository.Model.Description = description
			}
			if imgurl, isOk := p.Args["img_url"].(string); isOk {
				TableConfig.Repository.Model.ImgURL = imgurl
			}
			if maxGuests, isOk := p.Args["max_guests"].(int); isOk {
				TableConfig.Repository.Model.MaxGuests = maxGuests
			}

			if name, isOk := p.Args["name"].(string); isOk {
				TableConfig.Repository.Model.Name = name
				if areaID, isOk := p.Args["area_id"].(int); isOk {
					TableConfig.Repository.Model.Area.ID = areaID
					return TableConfig.Repository.Create()
				}
				return nil, fmt.Errorf("You must provide an Area ")
			}
			return nil, fmt.Errorf("You must provide a Name ")
		}
		return nil, fmt.Errorf("Invalid Credentials")
	},
}
