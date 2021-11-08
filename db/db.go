package db

import (
	"log"

	"github.com/PabloG6/blockchain-tutorial/config"
	ticketdb "github.com/PabloG6/blockchain-tutorial/db/leveldb"
)

//wrapper for the leveldb database.

//generates a new database option
func New(config config.Config) (*ticketdb.Database, error) {
	db, err :=  ticketdb.New(config.FileName)
	if err != nil {
		log.Fatalf("failed to create new db with filename: %s error: %v", config.FileName, err)
	}
	return db, nil;
}



