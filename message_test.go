package main

import (
	"bytes"

	. "gopkg.in/check.v1"
)

// Hook up gocheck into the "go test" runner.
// func Test(t *testing.T) { TestingT(t) }

// type MySuite struct{}

// var _ = Suite(&MySuite{})

func (s *MySuite) TestMessageQueueFromJson(c *C) {
	b := []byte(`{"queue":"bla","message": "bla", "date":"2015-03-03"}`)

	var queueMessage = QueueMessageFromJson(bytes.NewReader(b))

	c.Assert(queueMessage.Queue, Equals, "bla")
	c.Assert(queueMessage.Message, Equals, "bla")
	c.Assert(queueMessage.Date, Equals, "2015-03-03")
}

func (s *MySuite) TestQueueMessageToPublish(c *C) {
	queueMessage := QueueMessage{"bla", "bla", "2016-03-03"}
	value := QueueMessageToPublish(queueMessage)
	c.Assert(value, Equals, "bla:2016-03-03:bla")
}
