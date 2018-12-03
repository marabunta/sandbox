package main

import (
	"bytes"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	data, err := ioutil.ReadFile("bundle.crt")
	if err != nil {
		log.Fatal(err)
	}
	var blocks []byte
	for {
		var block *pem.Block
		block, data = pem.Decode(data)
		if block == nil {
			fmt.Printf("Error: PEM not parsed\n")
			break
		}
		blocks = append(blocks, block.Bytes...)
		if len(data) == 0 {
			break
		}
	}
	certs, err := x509.ParseCertificates(blocks)
	if err != nil {
		log.Fatal(err)
	}

	var certCA, certAnt *x509.Certificate

	for _, cert := range certs {
		if bytes.Compare(cert.RawIssuer, cert.RawSubject) == 0 && cert.IsCA {
			certCA = cert
		} else {
			certAnt = cert
		}
	}

	fmt.Printf("certCA = %+v\n", certCA.IsCA)
	fmt.Printf("cert = %+v\n", certAnt.IsCA)
}
