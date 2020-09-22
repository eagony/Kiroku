package extensions

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

var myredis redis.Conn

func init() {
	conn, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		fmt.Println("Connect redis failed,", err)
		return
	}
	myredis = conn
}

// Redis ...
func Redis() redis.Conn {
	return myredis
}
