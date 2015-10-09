package handlers

import (
	"net/http"

	"github.com/garyburd/redigo/redis"
	"github.com/rodrigodealer/queue-event-dispatcher/models"
)

func PublishHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Gorilla!\n"))

	c, err := redis.Dial("tcp", "192.168.99.100:6379")
	if err != nil {
		panic("Error connecting to redis")
	}

	c.Do("PUBLISH", "example", models.QueueMessageToPublish(models.QueueMessageFromJson(r.Body)))
	defer c.Close()
}
