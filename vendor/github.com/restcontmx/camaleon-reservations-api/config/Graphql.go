package config

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/restcontmx/camaleon-reservations-api/app/data"

	"github.com/graphql-go/graphql"

	handler "github.com/graphql-go/graphql-go-handler"

	_ "github.com/heroku/x/hmetrics/onload"
)

// For some reason the with value function of the context obj
// doesn't want to set the variable with the
type key int

const authKey key = 0

//
// Schema this will add the query type to the main query of the graphql config
//
var Schema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query:    QueryType,
	Mutation: MutationQueryType,
})

//
// MutationQueryType this will set all the mutation queries
//
var MutationQueryType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Mutation",
	Fields: graphql.Fields{
		"user":              CreateUpdateUser,
		"deleteUser":        DeleteUser,
		"business":          CreateUpdateBusiness,
		"deleteBusiness":    DeleteBusiness,
		"location":          CreateUpdateLocation,
		"deleteLocation":    DeleteLocation,
		"area":              CreateUpdateArea,
		"reservationStatus": CreateUpdateReservationStatus,
		"table":             CreateUpdateTable,
		"clientInfo":        CreateUpdateClientInfo,
		"reservation":       CreateUpdateReservation,
	},
})

//
// QueryType the query type for graphql
// This is the main query to add an resolve the objects
// the api is going to use
//
var QueryType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"user":  RetrieveUser,
			"users": GetAllUsers,

			"business":   RetrieveBusiness,
			"businesses": GetAllBusinesses,

			"location":  RetrieveLocation,
			"locations": GetAllLocations,

			"area":  RetrieveArea,
			"areas": GetAllAreas,

			"reservationStatuses": GetAllReservationStatus,
			"tables":              GetAllTables,

			"clientInfo":  RetrieveClientInfo,
			"clientInfos": GetAllClientInfo,

			"reservation":  RetrieveReservation,
			"reservations": GetAllReservations,
		},
	},
)

//
// InitRepositories inits all the repositories on the server
// @params none
// @returns none
//
func InitRepositories() {
	db, err := data.OpenConnection()
	if err != nil {
		log.Fatalf("%s", err)
	}

	userRepo := &data.UserRepository{DB: db}
	businessRepo := &data.BusinessRepository{DB: db}
	locationRepo := &data.LocationRepository{DB: db}
	areaRepo := &data.AreaRepository{DB: db}
	reservationStatusRepo := &data.ReservationStatusRepository{DB: db}
	tableRepo := &data.TableRepository{DB: db}
	clientInfoRepo := &data.ClientInfoRepository{DB: db}
	reservationRepo := &data.ReservationRepository{DB: db}

	UserConfig = UserConfiguration{Repository: userRepo}
	BusinessConfig = BusinessConfiguration{Repository: businessRepo}
	LocationConfig = LocationConfiguration{Repository: locationRepo}
	AreaConfig = AreaConfiguration{Repository: areaRepo}
	ReservationStatusConfig = ReservationStatusConfiguration{Repository: reservationStatusRepo}
	TableConfig = TableConfiguration{Repository: tableRepo}
	ClientInfoConfig = ClientInfoConfiguration{Repository: clientInfoRepo}
	ReservationConfig = ReservationConfiguration{Repository: reservationRepo}
}

//
// GraphqlHandlerFunc will create all the handler coniguration
// This way I can get the params context
// @param http response writer
// @param http request
// @returns none
//
func GraphqlHandlerFunc(w http.ResponseWriter, r *http.Request) {
	// get query
	opts := handler.NewRequestOptions(r)
	basicToken := r.Header.Get("Authorization")

	ctx := context.WithValue(r.Context(), authKey, basicToken)

	// execute graphql query
	params := graphql.Params{
		Schema:         Schema,
		RequestString:  opts.Query,
		VariableValues: opts.Variables,
		OperationName:  opts.OperationName,
		Context:        ctx,
	}

	result := graphql.Do(params)

	var buff []byte

	w.WriteHeader(http.StatusOK)

	buff, _ = json.MarshalIndent(result, "", "\t")

	w.Write(buff)
}
