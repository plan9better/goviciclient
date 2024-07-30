package goviciclient

// Observable is a type that can be observed.
type LocalAuth struct {
	Auth    string   `vici:"auth" json:"auth"`
	Class   string   `vici:"class" json:"class"`
	ID      string   `vici:"id"`
	Groups  []string `vici:"groups"`
	Certs   []string `vici:"certs"`
	Cacerts []string `vici:"cacerts"`
}

type RemoteAuth struct {
	Auth    string   `vici:"auth" json:"auth"`
	Class   string   `vici:"class" json:"class"`
	ID      string   `vici:"id"`
	Groups  []string `vici:"groups"`
	Certs   []string `vici:"certs"`
	Cacerts []string `vici:"cacerts"`
}

type ChildSA struct {
	Mode     string   `vici:"mode" json:"mode"`
	LocalTS  []string `vici:"local-ts" json:"local_ts"`
	RemoteTS []string `vici:"remote-ts" json:"remote_ts"`
}

type IKEConnection struct {
	LocalAddrs  []string           `vici:"local_addrs" json:"local_addrs"`
	RemoteAddrs []string           `vici:"remote_addrs" json:"remote_addrs"`
	Version     string             `vici:"version" json:"version"`
	ReauthTime  int                `vici:"reauth_time" json:"reauth_time"`
	RekeyTime   int                `vici:"rekey_time" json:"rekey_time"`
	LocalAuths  LocalAuth          `vici:"local-1" json:"local"`
	RemoteAuths RemoteAuth         `vici:"remote-1" json:"remote"`
	Children    map[string]ChildSA `vici:"children" json:"children"`
}

// ==============config strict================
type ChildSAConfig struct {
	Mode         string   `vici:"mode" json:"mode"`
	StartAction  string   `vici:"start_action" json:"start_action"`
	EspProposals string   `vici:"esp_proposals" json:"esp_proposals"`
	LocalTS      []string `vici:"local_ts" json:"local_ts"`
	RemoteTS     []string `vici:"remote_ts" json:"remote_ts"`
}

type LocalAuthConfig struct {
	ID      string   `vici:"id" json:"id"`
	Auth    string   `vici:"auth" json:"auth"`
	Certs   []string `vici:"certs" json:"certs"`
	Cacerts []string `vici:"cacerts" json:"cacerts"`
}

type RemoteAuthConfig struct {
	ID      string   `vici:"id" json:"id"`
	Auth    string   `vici:"auth" json:"auth"`
	Certs   []string `vici:"certs" json:"certs"`
	Cacerts []string `vici:"cacerts" json:"cacerts"`
}

type IKEConfig struct {
	LocalAddrs  []string                 `vici:"local_addrs" json:"local_addrs"`
	RemoteAddrs []string                 `vici:"remote_addrs" json:"remote_addrs"`
	Proposals   []string                 `vici:"proposals" json:"proposals"`
	Version     string                   `vici:"version" json:"version"`
	ReauthTime  int                      `vici:"reauth_time" json:"reauth_time"`
	RekeyTime   int                      `vici:"rekey_time" json:"rekey_time"`
	LocalAuths  *LocalAuthConfig         `vici:"local" json:"local"`
	RemoteAuths RemoteAuthConfig         `vici:"remote" json:"remote"`
	Children    map[string]ChildSAConfig `vici:"children" json:"children"`
}

type ConnectionsMap map[string]IKEConnection
type ConnectionsNames struct {
	Conns []string `vici:"conns" json:"conns"`
}
