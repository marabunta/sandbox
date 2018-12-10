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
	m := &sync.Map{}

	for i := 0; i < 10; i++ {
		d, ok := m.Load("id")
		if !ok {
			d = [][]string{[]string{"---------:5000", fmt.Sprintf("%d", i)}}
		} else {
			d = append(d.([][]string), []string{fmt.Sprintf("127.0.0.1:500%d", i), fmt.Sprintf("%d", i)})
		}
		fmt.Printf("len(d) = %+v\n", len(d.([][]string)))
		m.Store("id", d)
	}

	mymap := func(k, v interface{}) bool {
		fmt.Printf("k = %+v\n", k)
		fmt.Printf("v = %+v\n", v)
		return true
	}
	m.Range(mymap)

	for i := 9; i > 0; i-- {
		if items, ok := m.Load("id"); ok {
			d := items.([][]string)
			d = append(d[:i], d[i+1:]...)
			if len(d) <= 1 {
				fmt.Println("delete...")
				m.Delete("id")
			} else {
				m.Store("id", d)
			}
			fmt.Printf("d = %+v\n", d)
		}
	}

	_, ok := m.Load("id")
	fmt.Printf("ok = %+v\n", ok)
}
