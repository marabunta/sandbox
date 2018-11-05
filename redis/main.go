package main

import (
	"fmt"
	"log"

	"github.com/gomodule/redigo/redis"
)

func main() {
	pool := &redis.Pool{
		MaxIdle: 10,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", "localhost:6379")
			if err != nil {
				log.Fatal(err)
			}
			return c, err
		},
	}
	conn := pool.Get()
	defer conn.Close()

	pong, err := conn.Do("PING")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(pong)
}
