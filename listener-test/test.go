package main

import (
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	ripc "github.com/stormi-li/Ripc"
)

var redisAddr = "118.25.196.166:3934"
var password = "12982397StrongPassw0rd"

func main() {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: password,
	})
	client1 := ripc.NewClient(redisClient, "ripc-namespace")
	listener := client1.NewListener("ripc-channel")
	listener.Listen(func(msg string) {
		fmt.Println(msg)
		if msg == "shutdown" {
			listener.Close()
		}
		time.Sleep(500 * time.Millisecond)
	})
	fmt.Println("listener stopped")
}
