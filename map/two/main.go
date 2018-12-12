package main

import (
	"fmt"
	"sync"
)

func main() {

	clients := &sync.Map{}
	for i := 0; i < 100; i++ {
		clientID := fmt.Sprintf("client-%d", i%5)
		connID := fmt.Sprintf("conn-%d", i)
		client, ok := clients.LoadOrStore(clientID, connID)
		if ok {
			switch c := client.(type) {
			case string:
				client = []string{c, connID}
			case []string:
				client = append(client.([]string), connID)
			}
			clients.Store(clientID, client)
		}
	}

	list := func(k, v interface{}) bool {
		fmt.Printf("k: %+v v: %v\n", k, v)
		return true
	}

	clients.Range(list)

	// Get all connections
	broadcast := func(k, v interface{}) bool {
		fmt.Printf("To all connections from= %q\n", k)
		for _, v := range v.([]string) {
			fmt.Printf("msg to conn = %s\n", v)
		}
		return true
	}
	clients.Range(broadcast)
}
