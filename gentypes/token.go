package gentypes



//stores important information about a minted token
type Token struct {
	TokenName string `json:"string"`;
	TokenId   string `json:"string"`;
	
	Minted bool `json:"string"`; //an easy way to figure out whether a token has been minted

}