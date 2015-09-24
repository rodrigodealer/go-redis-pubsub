package main

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

func Subscribe() {
	c, err := redis.Dial("tcp", "192.168.99.100:6379")
	if err != nil {
		panic(err)
	}
	defer c.Close()

	psc := redis.PubSubConn{c}
	psc.Subscribe("example")
	for {
		switch v := psc.Receive().(type) {
		case redis.Message:
			fmt.Printf("%s: message: %s\n", v.Channel, v.Data)
		case redis.Subscription:
			fmt.Printf("%s: %s %d\n", v.Channel, v.Kind, v.Count)
		case error:
			fmt.Printf("%s", v)
		}
	}
}