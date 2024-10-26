package main

import (
	"fmt"

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
	res := client.Wait("my-channel", 0)
	fmt.Println(res)
}
