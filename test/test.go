package main

import (
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
	go func() {
		redisClient := redis.NewClient(&redis.Options{
			Addr: redisAddr,
		})
		client1 := ripc.NewClient(redisClient)
		listener := client1.NewListener("c1")
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
	time.Sleep(3 * time.Second)
}
