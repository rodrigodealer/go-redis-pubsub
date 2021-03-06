package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/garyburd/redigo/redis"
	"github.com/rodrigodealer/queue-event-dispatcher/models"
)

type RedisClient struct {
	conn redis.Conn
	redis.PubSubConn
	sync.Mutex
}

func NewRedisClient(host string, port string) *RedisClient {
	host = fmt.Sprintf("%s:%s", host, port)
	conn, _ := redis.Dial("tcp", host)
	pubsub, _ := redis.Dial("tcp", host)
	client := RedisClient{conn, redis.PubSubConn{pubsub}, sync.Mutex{}}
	go func() {
		for {
			time.Sleep(200 * time.Millisecond)
			client.Lock()
			client.conn.Flush()
			client.Unlock()
		}
	}()
	return &client
}

func (client *RedisClient) Publish(channel, message string) {
	client.Lock()
	client.conn.Send("PUBLISH", channel, message)
	client.Unlock()
}

func (client *RedisClient) Receive() models.Message {
	switch message := client.PubSubConn.Receive().(type) {
	case redis.Message:
		return models.Message{"message", message.Channel, string(message.Data)}
	case redis.Subscription:
		return models.Message{message.Kind, message.Channel, string(message.Count)}
	}
	return models.Message{}
}
