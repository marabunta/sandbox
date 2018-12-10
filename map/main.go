package main

import (
	"fmt"
	"sync"
)

type Stream struct {
	ip   string
	text string
}

func main() {
	var (
		m sync.Map
	)

	for i := 0; i < 10; i++ {
		d, ok := m.Load("id")
		if !ok {
			d = map[string]int{"127.0.0.1:5000": i}
		} else {
			d.(map[string]int)[fmt.Sprintf("127.0.0.1:500%d", i)] = i
		}
		m.Store("id", d)
	}

	mymap := func(k, v interface{}) bool {
		fmt.Printf("k = %+v\n", k)
		fmt.Printf("v = %+v\n", v)
		return true
	}
	m.Range(mymap)

}
