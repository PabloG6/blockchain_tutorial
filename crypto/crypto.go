package crypto

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"encoding/hex"
	"log"
)


//right now simply generate an address for the user. no need to get into all that crypto bullshit for now.
func GenerateAddress() (string) {

	key, err := ecdsa.GenerateKey(elliptic.P224(), rand.Reader)
	if err != nil {
		log.Fatal(err);
	}


	return hex.EncodeToString(key.D.Bytes());
	


}