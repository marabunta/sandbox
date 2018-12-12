package main

import (
	"fmt"
	"sync"
)

type foo struct {
	bar string
}

func main() {
	m := &sync.Map{}

	test := foo{}

	streams, ok := m.Load("test")
	if !ok {
		streams = []foo{test}
	}
	m.Store("test", streams)

	streams2, ok := m.Load("test")
	if !ok {
		fmt.Println("have it")
	} else {
		fmt.Printf("streams2 = %v\n", streams2)
		streams2 = append(streams2.([]foo), test)
	}

	m.Store("test", streams2)

	x, ok := m.Load("test")

	fmt.Printf("x = %+v\n", x)

}
