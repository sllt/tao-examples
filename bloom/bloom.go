package main

import (
	"fmt"

	"github.com/sllt/tao/core/bloom"
	"github.com/sllt/tao/core/stores/redis"
)

func main() {
	store := redis.New("localhost:6379")
	filter := bloom.New(store, "testbloom", 64)
	filter.Add([]byte("kevin"))
	filter.Add([]byte("wan"))
	fmt.Println(filter.Exists([]byte("kevin")))
	fmt.Println(filter.Exists([]byte("wan")))
	fmt.Println(filter.Exists([]byte("nothing")))
}
