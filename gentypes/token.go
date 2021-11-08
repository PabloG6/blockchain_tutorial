package gentypes



//stores important information about a minted token
type Token struct {
	tokenName string `json:"string"`;
	tokenId   string `json:"string"`;
	
	minted bool `json:"string"`; //an easy way to figure out whether a token has been minted

}