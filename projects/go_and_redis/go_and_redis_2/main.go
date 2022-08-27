package main

import (
	"fmt"
	"github.com/go-redis/redis/v8"
)

func main()  {
	fmt.Println("Testing Golang - Redis")

	client := redis.NewClient(&redis.Options{
		Addr: "192.168.1.7:6379",
		Password: "",
		DB: 0,
	})

	pong, err := client.Ping(client.Context()).Result()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(pong)
}