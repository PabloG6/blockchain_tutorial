package core

import (
	"testing"

	
)

func TestBlockChainSeed(test *testing.T) {
	test.Run("NewDatabase", func(t *testing.T) {

		
		bc := NewBlockChain()
		config := Config{fn: t.TempDir()}
		masterNode := bc.Seed(config);
		masterNodeData, err := bc.db.Get([]byte(masterNode.Address));
		t.Logf("masterNodeData: %s", string(masterNodeData))
		if err != nil {
			t.Fatalf("Failed to write values to db %v", err)
		}

		t.Cleanup(func() { bc.db.Close() })
	})

}
