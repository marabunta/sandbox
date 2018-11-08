package main

import (
	"crypto/rand"
	"fmt"
	"log"
	"time"

	"github.com/gomodule/redigo/redis"
)

func uuid() string {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		log.Fatal(err)
	}
	return fmt.Sprintf("%X-%X-%X-%X-%X", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
}

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

	for i := 0; i < 10; i++ {
		conn.Send("ZADD", "marabunta:todo", time.Now().Unix(), uuid())
		if _, err := conn.Do(""); err != nil {
			log.Fatal(err)
		}
		time.Sleep(time.Second)
	}
}
