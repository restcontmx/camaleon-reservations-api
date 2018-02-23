package config

import (
	"fmt"

	"github.com/graphql-go/graphql"
	"github.com/restcontmx/camaleon-reservations-api/app/data"
	"github.com/restcontmx/camaleon-reservations-api/app/models"
)

// ClientInfoConfiguration configuration client info
type ClientInfoConfiguration struct {
	Repository *data.ClientInfoRepository
}

// ClientInfoConfig is the global client info configuration
var ClientInfoConfig ClientInfoConfiguration

//
// ClientInfoObject is the graphql object of the client info
//
var ClientInfoObject = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "ClientInfo",
		Fields: graphql.Fields{
			"ID": &graphql.Field{
				Type: graphql.String,
			},
			"FirstName": &graphql.Field{
				Type: graphql.String,
			},
			"LastName": &graphql.Field{
				Type: graphql.String,
			},
			"Email": &graphql.Field{
				Type: graphql.String,
			},
			"Phone": &graphql.Field{
				Type: graphql.String,
			},
			"LocationID": &graphql.Field{
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
// CreateUpdateClientInfo graphql configuration
//
var CreateUpdateClientInfo = &graphql.Field{
	Type: ClientInfoObject,
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"location": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"firstname": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"lastname": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"email": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"phone": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		isOk, _ := ValidateAuthentication(p.Context.Value(authKey).(string))
		if isOk {
			id, isOk := p.Args["id"].(int)
			ClientInfoConfig.Repository.Model = models.ClientInfoModel{}
			if isOk {
				entity, _ := ClientInfoConfig.Repository.GetByID(id)
				ClientInfoConfig.Repository.Model = entity.(models.ClientInfoModel)

				if firstname, isOk := p.Args["firstname"].(string); isOk {
					ClientInfoConfig.Repository.Model.FirstName = firstname
				}
				if lastname, isOk := p.Args["lastname"].(string); isOk {
					ClientInfoConfig.Repository.Model.LastName = lastname
				}
				if email, isOk := p.Args["email"].(string); isOk {
					ClientInfoConfig.Repository.Model.Email = email
				}
				if phone, isOk := p.Args["phone"].(string); isOk {
					ClientInfoConfig.Repository.Model.Phone = phone
				}
				return ClientInfoConfig.Repository.Update()
			}
			if firstname, isOk := p.Args["firstname"].(string); isOk {
				ClientInfoConfig.Repository.Model.FirstName = firstname
			}
			if lastname, isOk := p.Args["lastname"].(string); isOk {
				ClientInfoConfig.Repository.Model.LastName = lastname
			}
			if email, isOk := p.Args["email"].(string); isOk {
				ClientInfoConfig.Repository.Model.Email = email
				if phone, isOk := p.Args["phone"].(string); isOk {
					ClientInfoConfig.Repository.Model.Phone = phone
					if location, isOk := p.Args["location"].(int); isOk {
						ClientInfoConfig.Repository.Model.LocationID = location
						return ClientInfoConfig.Repository.Create()
					}
					return nil, fmt.Errorf("You must provide a Location ")
				}
				return nil, fmt.Errorf("You must provide a Phone ")
			}
			return nil, fmt.Errorf("You must provide an Email ")
		}
		return nil, fmt.Errorf("Invalid Credentials")
	},
}

//
// GetAllClientInfo graphql configuration
//
var GetAllClientInfo = &graphql.Field{
	Type: graphql.NewList(ClientInfoObject),
	Args: graphql.FieldConfigArgument{},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		isOk, _ := ValidateAuthentication(p.Context.Value(authKey).(string))
		if isOk {
			return ClientInfoConfig.Repository.GetAll()
		}
		return nil, fmt.Errorf("Invalid Credentials")
	},
}

//
// RetrieveClientInfo graphql configuration
//
var RetrieveClientInfo = &graphql.Field{
	Type: ClientInfoObject,
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
				return ClientInfoConfig.Repository.GetByID(id)
			}
			return nil, fmt.Errorf("There is no id field")
		}
		return nil, fmt.Errorf("Invalid Credentials")
	},
}
