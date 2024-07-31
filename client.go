package goviciclient

import (
	"github.com/strongswan/govici/vici"
)

// ViciClient 包装了与 VICI 协议通信的客户端
type ViciClient struct {
	session *vici.Session
}

// NewViciClient 创建一个新的 VICI 客户端
func NewViciClient(addr string, network string) (*ViciClient, error) {
	if network == "" {
		network = "unix"
	}
	if addr == "" {
		addr = "/var/run/charon.vici"
	}
	session, err := vici.NewSession(vici.WithAddr(network, addr))
	if err != nil {
		return nil, err
	}
	return &ViciClient{session: session}, nil
}

func (c *ViciClient) ListConns(ike string) ([]ConnectionsMap, error) {
	msg := vici.NewMessage()
	if ike != "" {
		msg.Set("ike", ike)
	} else {
		msg.Unset("ike")
	}
	resp, err := c.streamedCommandRequest("list-conns", "list-conn", msg)
	if err != nil {
		return nil, err
	}
	conns := make([]ConnectionsMap, 0)
	for _, m := range resp {
		var conn ConnectionsMap = make(map[string]IKEConnection)
		if err := vici.UnmarshalMessage(m, conn); err != nil {
			return nil, err
		}
		conns = append(conns, conn)
	}
	return conns, nil
}

func (c *ViciClient) ListConnsNames() (*ConnectionsNames, error) {
	resp, err := c.commandRequest("get-conns", vici.NewMessage())
	if err != nil {
		return nil, err
	}
	var names *ConnectionsNames = new(ConnectionsNames)
	if err := vici.UnmarshalMessage(resp, names); err != nil {
		return nil, err
	}
	return names, nil
}

func (c *ViciClient) LoadConns(conns map[string]IKEConfig) error {
	var msg *vici.Message
	var err error

	if msg, err = vici.MarshalMessage(conns); err != nil {
		return err
	}

	_, err = c.commandRequest("load-conn", msg)
	if err != nil {
		return err
	}
	return err
}

func (c *ViciClient) UnloadConns(conns string) error {
	msg := vici.NewMessage()
	msg.Set("name", conns)
	_, err := c.commandRequest("unload-conn", msg)
	return err
}

func (c *ViciClient) ListCerts() ([]string, error) {
	// m := vici.NewMessage()

	return nil, nil
}

// Close 关闭 VICI 客户端的会话
func (c *ViciClient) Close() error {
	return c.session.Close()
}

// CommandRequest 发送一个命令请求到服务器，并返回响应
func (c *ViciClient) commandRequest(cmd string, msg *vici.Message) (*vici.Message, error) {
	return c.session.CommandRequest(cmd, msg)
}

// StreamedCommandRequest 发送一个流式命令请求
func (c *ViciClient) streamedCommandRequest(cmd, event string, msg *vici.Message) ([]*vici.Message, error) {
	return c.session.StreamedCommandRequest(cmd, event, msg)
}

// Subscribe 订阅事件
func (c *ViciClient) subscribe(events ...string) error {
	return c.session.Subscribe(events...)
}

// Unsubscribe 取消订阅事件
func (c *ViciClient) unsubscribe(events ...string) error {
	return c.session.Unsubscribe(events...)
}
