package main

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"
)

type Message struct {
	Type    string
	Channel string
	Data    string
}

type QueueMessage struct {
	Queue   string `json:"queue"`
	Message string `json:"message"`
	Date    string `json:"date"`
}

func QueueMessageFromJson(jsonBody io.Reader) QueueMessage {
	var m QueueMessage
	decoder := json.NewDecoder(jsonBody)
	err := decoder.Decode(&m)
	if err != nil {
		panic(err)
	}
	return m
}

func QueueMessageFromSubscribe(value string) QueueMessage {
	values := strings.Split(value, ":")
	return QueueMessage{values[0], values[2], values[1]}
}

func QueueMessageToPublish(queueMessage QueueMessage) string {
	return fmt.Sprintf("%s:%s:%s", queueMessage.Queue, queueMessage.Date, queueMessage.Message)
}
