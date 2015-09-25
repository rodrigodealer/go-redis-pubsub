package main

import (
	"fmt"
	"net/http"

	"github.com/garyburd/redigo/redis"
)

func PublishHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Gorilla!\n"))

	c, err := redis.Dial("tcp", "192.168.99.100:6379")
	if err != nil {
		panic(err)
	}

	m := QueueMessageFromJson(r.Body)

	fmt.Print("%s", m)

	c.Do("PUBLISH", "example", "Hello World")
	defer c.Close()
}
