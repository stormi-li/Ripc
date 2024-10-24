package ripc

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

// Ripc结构体
type Client struct {
	redisClient *redis.Client
	Namespace   string
	Context     context.Context
}

// 创建Ripc客户端
func NewClient(redisClient *redis.Client) *Client {
	client := Client{redisClient: redisClient, Namespace: "", Context: context.Background()}
	return &client
}

// 设置命名空间，避免相同的键或管道的干扰
func (c *Client) SetNamespace(namespace string) {
	c.Namespace = namespace + ":"
}

// 向所有监听管道的进程发送通知
func (c *Client) Notify(channel, msg string) {
	//使用redis的Publish功能发送通知--------------------------redis代码
	c.redisClient.Publish(c.Context, c.Namespace+channel, msg)
}

// 监听一个消息，返回收到的信息，如果超时返回""
func (c *Client) Wait(channel string, timeout time.Duration) string {
	//使用redis的Subscribe功能订阅管道--------------------------redis代码
	sub := c.redisClient.Subscribe(c.Context, c.Namespace+channel)

	msgChan := sub.Channel()
	timer := time.NewTicker(timeout)
	defer timer.Stop()
	defer sub.Close()
	select {
	case <-timer.C:
		return ""
	case msg := <-msgChan:
		return msg.Payload
	}
}

// 监听器结构体
type Listener struct {
	shutdown chan struct{}
	sub      *redis.PubSub
}

// 关闭监听器
func (listener Listener) Close() {
	listener.shutdown <- struct{}{}
}

// 创建一个监听器
func (c *Client) NewListener(channel string) *Listener {
	listener := Listener{
		//使用redis的Subscribe功能订阅管道--------------------------redis代码
		sub:      c.redisClient.Subscribe(c.Context, c.Namespace+channel),
		shutdown: make(chan struct{}, 1),
	}
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
