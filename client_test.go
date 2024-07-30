package goviciclient

import (
	"fmt"
	"testing"

	"github.com/strongswan/govici/vici"
)

func TestExampleUsage(t *testing.T) {
	type LocalAuth struct {
		Class      string   `vici:"class"`
		EAPType    string   `vici:"eap-type"`
		EAPVendor  string   `vici:"eap-vendor"`
		XAuth      string   `vici:"xauth"`
		Revocation string   `vici:"revocation"`
		ID         string   `vici:"id"`
		AAAID      string   `vici:"aaa_id"`
		EapID      string   `vici:"eap_id"`
		XauthID    string   `vici:"xauth_id"`
		Groups     []string `vici:"groups"`
		Certs      []string `vici:"certs"`
	}

	type RemoteAuth struct {
		// 与 LocalAuth 类似，根据实际返回的数据结构进行调整
	}

	type ChildSAConfig struct {
		Mode         string   `vici:"mode"`
		Label        string   `vici:"label"`
		RekeyTime    int      `vici:"rekey_time"`
		RekeyBytes   int      `vici:"rekey_bytes"`
		RekeyPackets int      `vici:"rekey_packets"`
		LocalTS      []string `vici:"local-ts"`
		RemoteTS     []string `vici:"remote-ts"`
	}
	type IKEConnection struct {
		LocalAddrs  []string                 `vici:"local_addrs"`
		RemoteAddrs []string                 `vici:"remote_addrs"`
		Version     string                   `vici:"version"`
		ReauthTime  int                      `vici:"reauth_time"`
		RekeyTime   int                      `vici:"rekey_time"`
		LocalAuths  []LocalAuth              `vici:"local"`
		RemoteAuths []RemoteAuth             `vici:"remote"`
		Children    map[string]ChildSAConfig `vici:"children"`
	}

	// 用于存储所有连接的映射
	type ConnectionsMap map[string]IKEConnection
	client, err := NewViciClient("/var/run/charon.vici")
	if err != nil {
		panic(err)
	}
	defer client.Close()

	// 示例命令请求
	m := vici.NewMessage()

	// if err := m.Set("ike", ""); err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	ms, err := client.session.StreamedCommandRequest("list-conns", "list-conn", m)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(len(ms))
	fmt.Println(ms)
	// for _, f := range ms {
	// 	if f.Err() != nil {
	// 		fmt.Println(err)
	// 		return
	// 	}
	// 	if len(f.Keys()) == 0 {
	// 		continue
	// 	}
	// 	var ssc ConnectionsMap = make(ConnectionsMap)
	// 	err := vici.UnmarshalMessage(f, ssc)
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}
	// 	fmt.Println(ssc["ces"].RemoteAddrs)
	// }
}
