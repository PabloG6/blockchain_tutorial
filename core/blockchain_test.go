package core

import (
	"testing"

	
)

func TestBlockChainSeed(test *testing.T) {
	test.Run("NewDatabase", func(t *testing.T) {

		
		bc := NewBlockChain()
		masterNode := bc.Seed();
		masterNodeData, err := bc.db.Get([]byte(masterNode.Address));
		t.Logf("masterNodeData: %s", string(masterNodeData))
		if err != nil {
			t.Fatalf("Failed to write values to db %v", err)
		}

		t.Cleanup(func() { bc.db.Close() })
	})

}
