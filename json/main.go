package main

import (
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
)

// Payload json test
type Payload struct {
	Description string `json:"description"`
	Enabled     int    `json:"enabled"`
	Name        string `json:"name"`
	Retries     int    `json:"retries"`
	Target      string `json:"target"`
	When        string `json:"when"`
}

func main() {
	files, err := ioutil.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}
	hash := ""
	for _, f := range files {
		if filepath.Ext(f.Name()) == ".json" {
			h := getHash(f.Name())
			if hash != "" {
				if hash != h {
					fmt.Printf("file: %s hash: %s != previous %s\n", f.Name(), hash, h)
					return
				}
			}
			fmt.Printf("file: %s hash: %s\n", f.Name(), h)
			hash = h
		}
	}
}

func getHash(file string) string {
	byteValue, _ := ioutil.ReadFile(file)

	var payload Payload
	json.Unmarshal([]byte(byteValue), &payload)

	e, err := json.Marshal(payload)
	if err != nil {
		log.Fatal(err)
	}

	h := sha1.Sum(e) // len 20 (store as binary(20)
	return fmt.Sprintf("%x", h)
}
