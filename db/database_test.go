package db

import (
	"bytes"
	"os"

	"testing"

	"github.com/PabloG6/blockchain-tutorial/config"
	"github.com/pelletier/go-toml"
)

func TestNewDatabase(test *testing.T) {
	test.Run("NewDatabase", func(t *testing.T) {
		
		
	
		config := config.Config{DirName: test.TempDir()}
		db, err := New(config)
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
	test.Run("New Database with test.toml file", func(t *testing.T) {
		
		config := config.Config{}
		data, err := os.ReadFile("../test.toml");

		if err != nil {
			t.Fatal(err);
		}


		err = toml.Unmarshal(data, &config)
		if err != nil {
			t.Fatal(err);
		}


		db, err := New(config)
		if err != nil {
			t.Fatalf(`error %v`, err)

		}

		t.Cleanup(func ()  {
			os.RemoveAll(config.DirName)
			db.Close();
		})
	})

}

