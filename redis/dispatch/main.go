package main

import (
	"fmt"
	"log"
	"time"

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

	for {
		if _, err := conn.Do("WATCH", "marabunta:todo"); err != nil {
			log.Fatal(err)
		}

		task, err := redis.Strings(conn.Do("ZRANGEBYSCORE", "marabunta:todo", 0, time.Now().Unix(), "limit", 0, 1))
		if err != nil {
			// TODO
			log.Fatal(err)
		}
		if len(task) != 1 {
			conn.Do("UNWATCH")
			fmt.Println("sleep...")
			time.Sleep(time.Second)
			continue
		}
		conn.Send("MULTI")
		conn.Send("LPUSH", "marabunta:queued", task[0])
		conn.Send("ZREM", "marabunta:todo", task[0])
		queued, err := conn.Do("EXEC")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("queued = %+v\n", queued)
	}
}
