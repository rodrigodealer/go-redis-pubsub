package main

import (
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

	c.Do("PUBLISH", "example", QueueMessageToPublish(m))
	defer c.Close()
}
