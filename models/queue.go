package models

import "github.com/garyburd/redigo/redis"

type Queue struct {
	Name       string
	Url        string
	Parameters string
	Method     string
	Data       string
}

func GetQueueByName(name string) Queue {
	c, _ := redis.Dial("tcp", "192.168.99.100:6379")

	values, _ := redis.Values(c.Do("HGETALL", name))

	if len(values) > 0 {
		return Queue{string(values[1].([]byte)),
			string(values[3].([]byte)),
			string(values[5].([]byte)),
			string(values[7].([]byte)), ""}
	}
	return Queue{}
}
