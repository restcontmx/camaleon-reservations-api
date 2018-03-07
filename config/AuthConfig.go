package config

import (
	"fmt"

	"github.com/restcontmx/camaleon-reservations-api/app/models"

	"github.com/graphql-go/graphql"
)

// Login function for authentication
var Login = &graphql.Field{
	Type: UserReservationObject,
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
				user, err := UserConfig.Repository.AuthenticateUser()
				if err != nil {
					return nil, err
				}
				userP := user.(models.UserModel)
				return UserReservationConfig.Repository.GetByUserID(userP.ID)
			}
			return nil, fmt.Errorf("Missing 'password' field")
		}
		return nil, fmt.Errorf("Missing 'username' field")
	},
}
