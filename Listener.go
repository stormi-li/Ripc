package ripc

import "github.com/go-redis/redis/v8"

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
		sub:      c.redisClient.Subscribe(c.ctx, c.namespace+channel),
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
