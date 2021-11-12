package gentypes

import (
	"encoding/hex"

	"github.com/PabloG6/blockchain-tutorial/config"
	"github.com/PabloG6/blockchain-tutorial/crypto"
)

/*
	the master seed is responsible for seeding the network,
	liquidating coins, issuing new coins and managing access to the txpool
*/

type MasterNetwork struct {
	Address string;
	Tokens []Token;
	 
}




func NewMaster()(*MasterNetwork) {
	
	
	privKey, _ := crypto.GenerateKey();
	return &MasterNetwork{Address: hex.EncodeToString(privKey), Tokens: make([]Token, config.GlobalConfig.TokenSupply)}
	
}



