package goviciclient

import (
	"fmt"

	"github.com/strongswan/govici/vici"
)

// ViciClient 包装了与 VICI 协议通信的客户端
type ViciClient struct {
	session *vici.Session
}

// NewViciClient 创建一个新的 VICI 客户端
func NewViciClient(addr string) (*ViciClient, error) {
	session, err := vici.NewSession(vici.WithAddr("unix", addr))
	if err != nil {
		return nil, err
	}
	return &ViciClient{session: session}, nil
}

// Close 关闭 VICI 客户端的会话
func (c *ViciClient) Close() error {
	return c.session.Close()
}

// CommandRequest 发送一个命令请求到服务器，并返回响应
func (c *ViciClient) CommandRequest(cmd string, msg *vici.Message) (*vici.Message, error) {
	return c.session.CommandRequest(cmd, msg)
}

// StreamedCommandRequest 发送一个流式命令请求
func (c *ViciClient) StreamedCommandRequest(cmd, event string, msg *vici.Message) ([]*vici.Message, error) {
	return c.session.StreamedCommandRequest(cmd, event, msg)
}

// Subscribe 订阅事件
func (c *ViciClient) Subscribe(events ...string) error {
	return c.session.Subscribe(events...)
}

// Unsubscribe 取消订阅事件
func (c *ViciClient) Unsubscribe(events ...string) error {
	return c.session.Unsubscribe(events...)
}

// ExampleUsage 展示如何使用 ViciClient
func ExampleUsage() {
	client, err := NewViciClient("/var/run/charon.vici")
	if err != nil {
		panic(err)
	}
	defer client.Close()

	// 示例命令请求
	inParams := vici.NewMessage()
	inParams.Set("command", "status")
	outMessage, err := client.CommandRequest("list-conns", inParams)
	if err != nil {
		fmt.Println("命令请求失败:", err)
		return
	}
	fmt.Printf("命令响应: %+v\n", outMessage)

	// 取消订阅所有事件
	client.session.UnsubscribeAll()
}
