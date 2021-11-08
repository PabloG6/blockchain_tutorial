package db

import (
	"log"

	ticketdb "github.com/PabloG6/blockchain-tutorial/db/leveldb"
)

//wrapper for the leveldb database.

//generates a new database option
func New(filename string) (*ticketdb.Database, error) {
	db, err :=  ticketdb.New(filename)
	if err != nil {
		log.Fatalf("failed to create new db with filename: %s error: %v", filename, err)
	}
	return db, nil;
}



