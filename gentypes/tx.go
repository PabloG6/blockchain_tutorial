package gentypes;
//stores information on transactions on the blockchain

type Tx struct {
	tokens []Token
	accountFrom string;
	accountTo string;
}