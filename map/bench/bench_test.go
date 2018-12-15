package main

import (
	"fmt"
	"sync"
	"testing"
)

func BenchmarkMap(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	clients := &sync.Map{}
	for i := 0; i < b.N; i++ {
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
		listValue := func(j, l interface{}) bool {
			return true
		}
		v.(*sync.Map).Range(listValue)
		return true
	}

	clients.Range(list)

	// Get all connections
	broadcast := func(k, v interface{}) bool {
		listValue := func(j, l interface{}) bool {
			return true
		}
		v.(*sync.Map).Range(listValue)
		return true
	}
	clients.Range(broadcast)

}

func BenchmarkArray(b *testing.B) {
	b.ReportAllocs()
	clients := &sync.Map{}
	for i := 0; i < b.N; i++ {
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
		return true
	}

	clients.Range(list)

	// Get all connections
	broadcast := func(k, v interface{}) bool {
		switch t := v.(type) {
		case []string:
			for range t {
			}
		}
		return true
	}
	clients.Range(broadcast)
}
