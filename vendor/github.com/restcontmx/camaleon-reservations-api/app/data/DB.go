package data

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq" // this is for posgtgresql driver
)

// ConnectionObject is the main connection object
type ConnectionObject struct {
	DB    *sql.DB
	Error error
}

// ConnObj will work as a global object
var ConnObj = ConnectionObject{}

// ConnStr will contain the connection string
var ConnStr = os.Getenv("DATABASE_URL")

//
// Repository is the general contract for all the repositories
// GetByID - for getting a model by id
// GetAll - for getting all the objects
// Create - for creating a new model on the database
// Update - for updating model on the database
// Delete - for deleting model on the database
//
type Repository interface {
	GetAll() ([]interface{}, error)
	GetByID(int) (interface{}, error)
	Create() (interface{}, error)
	Update() (interface{}, error)
	Delete() (bool, error)
	GetByParams() (interface{}, error)
}

//
// OpenConnection this will open a postgresql connection
// @params none
// @returns connection
//
func OpenConnection() (*sql.DB, error) {
	db, err := sql.Open("postgres", ConnStr)
	return db, err
}

//
// CloseConnection will close the connection
// @params none
// @returns none
//
func CloseConnection() {
	log.Println("Connection Closed")
}
