package config

import (
	b64 "encoding/base64"
	"strings"

	"github.com/restcontmx/camaleon-reservations-api/app/models"
)

//
// ValidateAuthentication will format the basic token to username and password
// then will use the authentication repository for validate it on the database
// @params authKey string - the authentication string "Basic token"
// @returns boolean - if valid authentication true else false
//
func ValidateAuthentication(authKey string) (bool, interface{}) {

	token := strings.Split(authKey, " ")[1]
	sDec, _ := b64.StdEncoding.DecodeString(token)
	userInfo := strings.Split(string(sDec), ":")

	UserConfig.Repository.Model = models.UserModel{
		UserName: userInfo[0],
		Password: userInfo[1],
	}

	user, err := UserConfig.Repository.AuthenticateUser()

	if err != nil {
		return false, nil
	}

	return true, user
}
