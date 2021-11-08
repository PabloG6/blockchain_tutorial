package gentypes

import "github.com/PabloG6/blockchain-tutorial/crypto"

/*
	the master seed is responsible for seeding the network,
	liquidating coins, issuing new coins and managing access to the txpool
*/

type MasterNetwork struct {
	Address string;
	Tokens []Token;
	 

}




func NewMaster()(*MasterNetwork) {
	
	return &MasterNetwork{Address: crypto.GenerateAddress(), Tokens: make([]Token, 0)}
	
}



