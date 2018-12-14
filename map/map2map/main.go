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
		conns, ok := clients.Load(clientID)
		if ok {
			conns.(*sync.Map).Store(connID, i)
		} else {
			conns := &sync.Map{}
			conns.Store(connID, i)
			clients.Store(clientID, conns)
		}
	}

	list := func(k, v interface{}) bool {
		fmt.Printf("k: %v\n", k)
		listValue := func(j, l interface{}) bool {
			fmt.Printf("  j = %+v\n", j)
			return true
		}
		v.(*sync.Map).Range(listValue)
		return true
	}

	clients.Range(list)

	// Get all connections
	broadcast := func(k, v interface{}) bool {
		fmt.Printf("To all connections from= %q\n", k)
		listValue := func(j, l interface{}) bool {
			fmt.Printf("  msg to conn = %s -- %v\n", j, l)
			return true
		}
		v.(*sync.Map).Range(listValue)
		return true
	}
	clients.Range(broadcast)
}
