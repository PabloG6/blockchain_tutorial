package core

import (
	// "os"

	"os"
	"testing"

	"github.com/PabloG6/blockchain-tutorial/config"
	// "github.com/PabloG6/blockchain-tutorial/config"
)

func TestBlockChainSeed(test *testing.T) {
	test.Run("NewDatabase", func(t *testing.T) {

		// st
		
		bc := NewBlockChain()
		masterNode:= bc.Seed();
		_, err := bc.db.Get([]byte(masterNode.Address));

		if int64(len(masterNode.Tokens)) != config.GlobalConfig.TokenSupply {
			t.Fatal("token supply not saved");
		}
		if err != nil {
			t.Fatalf("Failed to write values to db %v", err)
		}

		t.Cleanup(func() {
			 bc.db.Close();
			os.RemoveAll(config.GlobalConfig.DirName)
		})
	})

}
