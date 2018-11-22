package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/asn1"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/satori/go.uuid"
)

func main() {
	key, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		panic(err)
	}

	uuid1, err := uuid.NewV1()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Something went wrong: %s", err)
		os.Exit(1)
	}

	sub := pkix.Name{
		CommonName:         fmt.Sprintf("%s", uuid1),
		Country:            []string{"-"},
		Province:           []string{"-"},
		Locality:           []string{"-"},
		Organization:       []string{"marabunta"},
		OrganizationalUnit: []string{"ant"},
	}

	asn1Sub, err := asn1.Marshal(sub.ToRDNSequence())
	if err != nil {
		fmt.Printf("unable to marshal asn1: %v", err)
		os.Exit(1)
	}

	template := x509.CertificateRequest{
		RawSubject:         asn1Sub,
		SignatureAlgorithm: x509.ECDSAWithSHA256,
	}

	csrCertificate, err := x509.CreateCertificateRequest(rand.Reader, &template, key)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	csr := pem.EncodeToMemory(&pem.Block{
		Type: "CERTIFICATE REQUEST", Bytes: csrCertificate,
	})

	err = ioutil.WriteFile("test.csr", csr, 0644)
	if err != nil {
		log.Fatal(err)
	}

	x509Encoded, _ := x509.MarshalECPrivateKey(key)
	pemEncoded := pem.EncodeToMemory(&pem.Block{Type: "PRIVATE KEY", Bytes: x509Encoded})

	err = ioutil.WriteFile("test.key", pemEncoded, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
