package main

import (
	"fmt"

	"github.com/go-redis/redis/v8"
	ripc "github.com/stormi-li/Ripc"
)

var redisAddr = "118.25.196.166:6379"

func main() {
	redisClient := redis.NewClient(&redis.Options{
		Addr: redisAddr,
	})
	client := ripc.NewClient(redisClient)
	listener := client.NewListener("c1") //参数为频道名称
	listener.Listen(func(msg string) {   //启动监听器
		fmt.Println(msg)
		if msg == "shutdown" {
			listener.Close() //关闭监听器
		}
	})
}
