package utils

import (
	"context"
	"fmt"
)

const (
	PublishKey = "websocket"
)

// Publish 发布消息到redis
// channel是发布的目标信道
// payload是要发布的消息内容
func Publish(ctx context.Context, channel string, message string) error {
	var err error
	fmt.Println("Publish=====", message)
	err = rdb.Publish(ctx, channel, message).Err()
	return err
}

// Subscribe 订阅redis消息
// channel是订阅的目标信道
func Subscribe(ctx context.Context, channel string) (string, error) {
	sub := rdb.Subscribe(ctx, channel)
	fmt.Println("Subscribe=====", ctx)
	msg, err := sub.ReceiveMessage(ctx)
	fmt.Println("Subscribe=====", msg.Payload)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return msg.Payload, err
}
