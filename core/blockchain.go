package core

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/PabloG6/blockchain-tutorial/config"
	"github.com/PabloG6/blockchain-tutorial/db"
	"github.com/PabloG6/blockchain-tutorial/db/leveldb"
	"github.com/PabloG6/blockchain-tutorial/gentypes"
)

//file for the blockchain protocol
type BlockChain struct {
	db *leveldb.Database
	genesisBlock *gentypes.GenesisBlock
}




//start ur own chain by seeding a block.
func (bc *BlockChain) Seed() (*gentypes.MasterNetwork) {

	// masterSeed := &gentypes.MasterNetwork{}
	masterNode := gentypes.NewMaster();
	marshall, err := json.Marshal(masterNode);

	if err != nil {
		log.Fatalf("format string %v", err);
	}

	bc.db.Put([]byte(masterNode.Address), marshall);
	return masterNode
	
}


//


func NewBlockChain() (*BlockChain) {
	fmt.Println(config.GlobalConfig.DirName)
	db, _ := db.New(config.GlobalConfig);
	return &BlockChain{
		db: db,
		
	}
}

func (bc *BlockChain) Connect() {

	//load and connect to the existing chain by connecting and seeding info from a master node.
}


