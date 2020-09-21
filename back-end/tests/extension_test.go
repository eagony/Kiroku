package tests

import (
	"fmt"
	extension "rinterest/extensions"
	"testing"
)

func TestRedis(t *testing.T) {
	myredis := extension.Redis()
	defer myredis.Close()

	_, err := myredis.Do("MSet", "abc", 100, "efg", 300)
	if err != nil {
		fmt.Println(err)
		return
	}
}
