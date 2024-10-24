package ripc

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

// Ripc结构体
type Client struct {
	RedisClient *redis.Client
}

// 创建Ripc客户端
func NewClient(addr string) (*Client, error) {
	client := Client{}
	rdsC := redis.NewClient(&redis.Options{
		Addr: addr,
	})

	// 测试连接
	if err := rdsC.Ping(context.Background()).Err(); err != nil {
		return nil, fmt.Errorf("failed to connect to Redis at %s", addr) // 返回错误
	}

	client.RedisClient = rdsC
	return &client, nil
}

var namespace = ""

func SetNameSpace(str string) {
	namespace = str + ":"
}

// 向所有监听管道的进程发送通知
func (client Client) Notify(ctx context.Context, channel, msg string) {
	client.RedisClient.Publish(ctx, namespace+channel, msg)
}

// 监听一个消息，返回收到的信息，如果超时返回""
func (client Client) Wait(ctx context.Context, channel string, timeout time.Duration) string {
	sub := client.RedisClient.Subscribe(ctx, namespace+channel)
	msg := sub.Channel()
	timer := time.NewTicker(timeout)
	defer timer.Stop()
	defer sub.Close()
	select {
	case <-timer.C:
		return ""
	case msg := <-msg:
		return msg.Payload
	}
}

type Listener struct {
	shutdown chan struct{}
	sub      *redis.PubSub
}

func (listener Listener) Close() {
	listener.shutdown <- struct{}{}
}

func (client *Client) NewListener(ctx context.Context, channel string) *Listener {
	listener := Listener{}
	listener.sub = client.RedisClient.Subscribe(ctx, namespace+channel)
	listener.shutdown = make(chan struct{}, 1)
	return &listener
}

// 接受所有发送过来的消息，并执行handler
func (listener *Listener) Listen(handler func(msg string)) {
	c := listener.sub.Channel()
	defer listener.sub.Close()
	for {
		select {
		case msg := <-c:
			handler(msg.Payload)
		case <-listener.shutdown:
			return
		}
	}
}
