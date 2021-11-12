package leveldb

import (
	"fmt"
	"log"
	"os"

	"github.com/syndtr/goleveldb/leveldb"
)

//wrapper for the leveldb database.

type Database struct {
	db *leveldb.DB
}

//generates a new database option
func New(filename string) (*Database, error) {
	fmt.Println(filename);
	if _, err := os.Stat(filename); !os.IsNotExist(err) {
		err = os.MkdirAll(filename, os.ModePerm)
		if err != nil {
			log.Panic(err)
		}

	}
	db, err := leveldb.OpenFile(filename, nil)

	if err != nil {

		log.Panicf("unable to create leveldb database %v %s", err, filename)
	}

	return &Database{db: db}, nil
}

func (ticketdb *Database) Put(key []byte, value []byte) (err error) {
	return ticketdb.db.Put(key, value, nil)

}

func (ticketdb *Database) Get(key []byte) ([]byte, error) {
	return ticketdb.db.Get(key, nil)
}

func (ticketdb *Database) Close() {
	ticketdb.db.Close()
}
