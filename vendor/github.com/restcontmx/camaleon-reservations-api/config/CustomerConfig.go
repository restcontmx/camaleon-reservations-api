package config

import (
	"fmt"

	"github.com/restcontmx/camaleon-reservations-api/app/models"

	"github.com/graphql-go/graphql"
	"github.com/restcontmx/camaleon-reservations-api/app/data"
)

//
// CustomerObject customer configuration
//
var CustomerObject = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Customer",
		Fields: graphql.Fields{
			"ID": &graphql.Field{
				Type: graphql.Int,
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
			"ZipCode": &graphql.Field{
				Type: graphql.String,
			},
			"Address1": &graphql.Field{
				Type: graphql.String,
			},
			"Address2": &graphql.Field{
				Type: graphql.String,
			},
			"Timestamp": &graphql.Field{
				Type: graphql.DateTime,
			},
			"Updated": &graphql.Field{
				Type: graphql.DateTime,
			},
			"Location": &graphql.Field{
				Type: LocationObject,
			},
		},
	},
)

//
// GetAllCustomers will return all the customers
//
var GetAllCustomers = &graphql.Field{
	Type: graphql.NewList(CustomerObject),
	Args: graphql.FieldConfigArgument{},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		isOk, _ := ValidateAuthentication(p.Context.Value(authKey).(string))
		if isOk {
			return data.CustomerRepository{}.GetAll()
		}
		return nil, fmt.Errorf("Invalid Credentials")
	},
}

//
// GetCustomer returns a sigle customer
//
var GetCustomer = &graphql.Field{
	Type: CustomerObject,
	Args: graphql.FieldConfigArgument{
		"location": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"phone": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		isOk, _ := ValidateAuthentication(p.Context.Value(authKey).(string))
		if isOk {
			location, _ := p.Args["location"].(int)
			phoneQuery, _ := p.Args["phone"].(string)

			return data.CustomerRepository{
				Model: models.CustomerModel{
					Phone:    phoneQuery,
					Location: models.LocationModel{ID: location},
				},
			}.GetByParams()
		}
		return nil, fmt.Errorf("Invalid Credentials")
	},
}
