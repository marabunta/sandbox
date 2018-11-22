package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509/pkix"
)

func main() {
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		panic(err)
	}

	subj := pkix.Name{
		CommonName:         "?????",
		Country:            []string{"-"},
		Province:           []string{"-"},
		Locality:           []string{"-"},
		Organization:       []string{"marabunta"},
		OrganizationalUnit: []string{"ant"},
	}
}
