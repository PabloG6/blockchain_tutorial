package leveldb

import (
	"log"

	"github.com/syndtr/goleveldb/leveldb"
)

//wrapper for the leveldb database.

type Database struct {
	db *leveldb.DB
}


//generates a new database option
func New(filename string) (*Database, error) {
	db, err :=  leveldb.OpenFile(filename, nil)
	if err !=nil {
		log.Fatalf("unable to create leveldb database %v", err)
	}

	return &Database{db: db}, nil
}


func (ticketdb *Database) Put(key[] byte, value[] byte) (err error) {
	return ticketdb.db.Put(key, value, nil); 

}

func (ticketdb *Database) Get(key[] byte)([]byte, error) {
	return ticketdb.db.Get(key, nil);
}

func (ticketdb *Database) Close() {
	ticketdb.db.Close();
}