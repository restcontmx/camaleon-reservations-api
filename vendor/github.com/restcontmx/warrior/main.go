package main

import (
	"log"
	"os"
)

//
// migration object
//
type migrationType struct {
	migrationName string
	instruction   string
	date          string
}

//
// create table
// this will create the create migration for tables
// @param tableName string
// @returns a migrationType
//
func createTable(tableName string) interface{} {
	return nil
}

//
// destroy table
// this will create the drop migration for tables
// @param tableName string
// @returns a migrationType
//
func destroyTable(tableName string) interface{} {
	return nil
}

//
// update table
// this will create the update migration
// @param tableName string
// @params params [] string the update parameters
// @returns a migrationType
//
func updateTable(tableName string, params []string) interface{} {
	return nil
}

//
// create migration
// This will create a migration according
// @param migration migrationType
// @returns none
//
func createMigration(migration migrationType) {

}

// 
// main function
//
func main() {
	log.Println("This is warrior")
	// Get the working directory not the binary path
	dir, err := os.Getwd() 
	// in case of any error getting the file
	if err != nil {
		log.Fatal(err)
	}
	// Do something with directory path
    log.Println(dir)
}