package config

import (
	"fmt"

	"github.com/graphql-go/graphql"
	"github.com/restcontmx/camaleon-reservations-api/app/data"
	"github.com/restcontmx/camaleon-reservations-api/app/models"
)

//
// UserConfiguration will be the object for graphql user configuration
//
type UserConfiguration struct {
	Repository *data.UserRepository
}

// UserConfig is the global user config that will be used
var UserConfig UserConfiguration

//
// UserObject is the graphql object of the user model
//
var UserObject = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "User",
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
			"UserName": &graphql.Field{
				Type: graphql.String,
			},
			"Email": &graphql.Field{
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
// CreateUpdateUser graphql configuration
//
var CreateUpdateUser = &graphql.Field{
	Type: UserObject,
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
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
		"username": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"password": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		isOk, _ := ValidateAuthentication(p.Context.Value(authKey).(string))
		if isOk {
			id, isOk := p.Args["id"].(int)
			UserConfig.Repository.Model = models.UserModel{}
			if isOk {
				userEntity, _ := UserConfig.Repository.GetByID(id)
				UserConfig.Repository.Model = userEntity.(models.UserModel)

				if firstname, isOk := p.Args["firstname"].(string); isOk {
					UserConfig.Repository.Model.FirstName = firstname
				}
				if lastname, isOk := p.Args["lastname"].(string); isOk {
					UserConfig.Repository.Model.LastName = lastname
				}
				if email, isOk := p.Args["email"].(string); isOk {
					UserConfig.Repository.Model.Email = email
				}
				if username, isOk := p.Args["username"].(string); isOk {
					UserConfig.Repository.Model.UserName = username
				}
				return UserConfig.Repository.Update()
			}
			if firstname, isOk := p.Args["firstname"].(string); isOk {
				UserConfig.Repository.Model.FirstName = firstname
			}
			if lastname, isOk := p.Args["lastname"].(string); isOk {
				UserConfig.Repository.Model.LastName = lastname
			}
			if email, isOk := p.Args["email"].(string); isOk {
				UserConfig.Repository.Model.Email = email
				if username, isOk := p.Args["username"].(string); isOk {
					UserConfig.Repository.Model.UserName = username
					if password, isOk := p.Args["password"].(string); isOk {
						UserConfig.Repository.Model.Password = password
						return UserConfig.Repository.Create()
					}
					return nil, fmt.Errorf("You must provide a Password")
				}
				return nil, fmt.Errorf("You must provide a UserName ")
			}
			return nil, fmt.Errorf("You must provide an Email ")

		}
		return nil, fmt.Errorf("Invalid Credentials")
	},
}

//
// GetAllUsers graphql configuration
//
var GetAllUsers = &graphql.Field{
	Type: graphql.NewList(UserObject),
	Args: graphql.FieldConfigArgument{},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		isOk, _ := ValidateAuthentication(p.Context.Value(authKey).(string))
		if isOk {
			return UserConfig.Repository.GetAll()
		}
		return nil, fmt.Errorf("Invalid Credentials")
	},
}

//
// RetrieveUser graphql configuration
//
var RetrieveUser = &graphql.Field{
	Type: UserObject,
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
				return UserConfig.Repository.GetByID(id)
			}
			return nil, fmt.Errorf("There is no id field")
		}
		return nil, fmt.Errorf("Invalid Credentials")
	},
}

//
// DeleteUser graphql configuration
//
var DeleteUser = &graphql.Field{
	Type: graphql.NewList(UserObject),
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
				UserConfig.Repository.Model = models.UserModel{ID: id}
				if isOk, _ := UserConfig.Repository.Delete(); isOk {
					return nil, nil
				}
				return nil, fmt.Errorf("There object was not deleted")
			}
			return nil, fmt.Errorf("No id found")
		}
		return nil, fmt.Errorf("Invalid Credentials")
	},
}
