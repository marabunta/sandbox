package main

import (
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
	c, err := x509.ParseCertificates(blocks)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	fmt.Printf("%s", c)
}
