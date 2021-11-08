package crypto

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"

	"log"
)

//right now simply generate an address for the user. no need to get into all that crypto bullshit for now.



func GenerateKey()(privKey []byte, pubKey []byte) {
	key, err := ecdsa.GenerateKey(elliptic.P384(), rand.Reader);
	if err != nil {
		log.Fatal(err);
	}
	privKey, err = x509.MarshalECPrivateKey(key);
	if err != nil {
		log.Fatal(err);
	}
	pubKey = elliptic.Marshal(key.Curve, key.PublicKey.X, key.PublicKey.Y)
	return privKey, pubKey

}
	

	

	

