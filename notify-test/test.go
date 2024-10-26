package main

import (
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
	client := ripc.NewClient(redisClient, "ripc-namespace")
	client.Notify("ripc-channel", "1")
	client.Notify("ripc-channel", "2")
	client.Notify("ripc-channel", "3")
	client.Notify("ripc-channel", "shutdown")
}
