# RIPC Guides


Simple and efficient cross-process communication library

# Overview

- Just need Redis
- No need Ip/Port
- Every feature comes with tests
- Developer Friendly

# Install


```shell
go get -u github.com/stormi-li/Ripc
```

# Quick Start

```go
package main

import (
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	ripc "github.com/stormi-li/Ripc"
)

var redisAddr = “localhost:6379”
var password = “your password”

func main() {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: password,
	})
	client := ripc.NewClient(redisClient, "ripc-namespace")

	go func() {
		res := client.Wait("ripc-channel", 0)
		fmt.Println(res)
	}()

	time.Sleep(100 * time.Millisecond)

	client.Notify("ripc-channel", "welcome use ripc")
}
```

# Interface-ripc

## NewClient

### Create ripc client
```go

package main

import (
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	ripc "github.com/stormi-li/Ripc"
)

var redisAddr = “localhost:6379”
var password = “your password”

func main() {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: password,
	})
	client := ripc.NewClient(redisClient, "ripc-namespace")
}
```

The first parameter is a redis client of successful connection, the second parameter is a unique namespace.

# Interface-ripc.Client

## Notify

### Send a message
```go
client.Notify("ripc-channel", "welcome use ripc")
```
The first parameter is a unique channel name, the second parameter is the message to be sent.

## Wait 

### Wait to receive a message
```go
res := client.Wait("my-channel", 0)
```
The first parameter is a unique channel name,  the second parameter is maximum waiting time.
The return value is the received message， timeout returns “”.

## NewListener

### Create Listener
```go
listener := client.NewListener("ripc-channel")
```
The first parameter is a unique channel name.
The return value is a listener for a specific channel.

# Interface-ripc.Listener

## Listen

### Start to listen specific channel
```go
listener.Listen(func(msg string) {
	fmt.Println(msg)
})
```
The parameter is a handler for received messages.

## Close

### Close listener
```go
listener.Close()
```

# Community

## Ask

### How do I ask a good question?
- Email - 2785782829@qq.com
- Github Issues - https://github.com/stormi-li/Ripc/issues