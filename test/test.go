package main

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	ripc "github.com/stormi-li/Ripc"
)

var redisAddr = "118.25.196.166:6379"

func main() {
	redisClient := redis.NewClient(&redis.Options{
		Addr: redisAddr,
	})
	client := ripc.NewClient(redisClient)
	client.SetNamespace("lili")
	listener := client.NewListener(context.Background(), "c1")
	go func() {
		listener.Listen(func(msg string) {
			fmt.Println(msg)
			if msg == "4" {
				listener.Close()
			}
			time.Sleep(500 * time.Millisecond)
		})
		fmt.Println("hhhh")
	}()
	go func() {
		fmt.Println(client.Wait("c1", 200*time.Millisecond))
	}()
	time.Sleep(100 * time.Millisecond)
	go func() {
		client.Notify("c1", "1")
		client.Notify("c1", "2")
		client.Notify("c1", "3")
		client.Notify("c1", "4")
	}()
	select {}
}
