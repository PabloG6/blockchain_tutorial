package db

import (
	"bytes"

	"testing"
)

func TestNewDatabase(test *testing.T) {
	test.Run("NewDatabase", func(t *testing.T) {
		
		
	
		db, err := New(test.TempDir())
		if err != nil {
			t.Fatalf(`error %v`, err)

		}

		db.Put([]byte("kv1"), []byte("kv2"))
		
		if err != nil {
			t.Fatalf("Failed to write values to db %v", err)
		}


		val, err := db.Get([]byte("kv1"))
		if err != nil {
			t.Fatalf("Failed to store data in leveldb db");
		}

		if !bytes.Equal(val, []byte("kv2")){
			t.Fatalf("Failed to store correct data in leveldb db %s", string(val));
		}

		
		t.Cleanup(func() {db.Close()})
	})
	
	//todo separate tests for cleaner reading.
	test.Run("Put information in database", func(t *testing.T) {})

}

