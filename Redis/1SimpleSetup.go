package main

import (
	"fmt"

	"github.com/go-redis/redis/v8"
)

func main() {
	fmt.Print("Hellow World")
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "eYVX7EwVmmxKPCDmwMtyKVge8oLd2t81",
		DB:       0,
	})

	pong, err := client.Ping(client.Context()).Result()
	fmt.Println(pong, err)
}
