# Ripc 框架

## 简介

Ripc 框架是用 Go 语言开发的一个基于 Redis 的 Pub/Sub 机制的高效跨进程通信框架。它提供了简单易用的接口，能够帮助开发者实现不同进程之间的消息传递。Ripc框架的三个主要功能—**Notify**、**Wait**和**Listen**，为开发者提供了灵活的消息传递解决方案，适用于各种实时应用场景。

## 安装

```shell
go get github.com/stormi-li/Ripc
```

## 使用

### 1. Notify

**Notify** 功能允许开发者向所有监听特定频道的进程广播消息。这种机制特别适合需要实时消息传递和共享数据的场景

**示例代码**：

```go
package main

import (
	"github.com/go-redis/redis/v8"
	ripc "github.com/stormi-li/Ripc"
)

var redisAddr = "your redis addr"

func main() {
	redisClient := redis.NewClient(&redis.Options{
		Addr: redisAddr,
	})
	client := ripc.NewClient(redisClient)
	client.Notify("channel-ripc", "helloword") //第一个参数为频道名称，第二个参数为传递的消息
}
```

### 2. Wait

**Wait** 功能允许进程等待并接收来自频道的单条消息。这种机制适用于事件驱动的场景。

**示例代码**：

```go
package main

import (
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	ripc "github.com/stormi-li/Ripc"
)

var redisAddr = "your redis addr"

func main() {
	redisClient := redis.NewClient(&redis.Options{
		Addr: redisAddr,
	})
	client := ripc.NewClient(redisClient)
	res := client.Wait("channel-ripc", 10*time.Second) //第一个参数为频道名称，第二个参数为最长等待时间
	fmt.Println(res)
}
```

### 3. Listen

**Listen**功能提供了持续监听频道的消息流，无需中断。这种机制适合于实时监控和处理场景。

**示例代码**：

```go
package main

import (
	"fmt"

	"github.com/go-redis/redis/v8"
	ripc "github.com/stormi-li/Ripc"
)

var redisAddr = "your redis addr"

func main() {
	redisClient := redis.NewClient(&redis.Options{
		Addr: redisAddr,
	})
	client := ripc.NewClient(redisClient)
	listener := client.NewListener("channel-ripc") //参数为频道名称
	listener.Listen(func(msg string) {   //启动监听器
		fmt.Println(msg)
		if msg == "shutdown" {
			listener.Close() //关闭监听器
		}
	})
}

```