package main

import (
	"crypto/tls"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	out, err := os.Create("/tmp/output.txt")
	defer out.Close()

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	// create a new request
	req, _ := http.NewRequest("GET", "http://localhost:8000/ca", nil)
	req.Header.Set("User-Agent", "ant")

	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	// Read Response Body
	n, err := io.Copy(out, res.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("n = %+v\n", n)
}
