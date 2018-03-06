package config

import (
	"fmt"

	"github.com/graphql-go/graphql"
)

// Login function for authentication
var Login = &graphql.Field{
	Type: UserObject,
	Args: graphql.FieldConfigArgument{
		"username": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"password": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		if username, isOk := p.Args["username"].(string); isOk {
			if password, isOk := p.Args["password"].(string); isOk {
				UserConfig.Repository.Model.UserName = username
				UserConfig.Repository.Model.Password = password
				return UserConfig.Repository.AuthenticateUser()
			}
			return nil, fmt.Errorf("Missing 'password' field")
		}
		return nil, fmt.Errorf("Missing 'username' field")
	},
}
