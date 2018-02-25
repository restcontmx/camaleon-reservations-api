package data

import (
	"database/sql"
	"log"
)

type ConnectionObject struct {
	DB    *sql.DB
	Error error
}

// ConnObj will work as a global object
var ConnObj = ConnectionObject{}

// ConnStr will contain the connection string
// var ConnStr = "postgres://rrgyhmwrbuwlms:02bae98c8f452f35d1ea720b73cba64200cbf084f2f2cea5e3d97dd2ba3c9a4f@ec2-107-20-255-96.compute-1.amazonaws.com/d98gaialpftbn2?sslmode=require"
var ConnStr = "postgres://jdrpybcnbztezu:19bdffd60d278b00bb5e1cced37636de1ec50274415f63418cc1011aa34bd6c1@ec2-54-227-251-233.compute-1.amazonaws.com:5432/d54lb0tevtmv9k?sslmode=require"

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
