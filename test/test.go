package main

import (
	"context"
	"fmt"
	"time"

	ripc "github.com/stormi-li/Ripc"
)

func main() {
	client, err := ripc.NewClient("118.25.196.166:6379")
	ripc.SetNameSpace("lili")
	if err != nil {
		fmt.Println(err)
		return
	}
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
		fmt.Println(client.Wait(context.Background(), "c1", 200*time.Millisecond))
	}()
	time.Sleep(100 * time.Millisecond)
	go func() {
		client.Notify(context.Background(), "c1", "1")
		client.Notify(context.Background(), "c1", "2")
		client.Notify(context.Background(), "c1", "3")
		client.Notify(context.Background(), "c1", "4")
	}()

	select {}
}
