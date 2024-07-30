package goviciclient

// ViciConfig represents the configuration structure with vici tags.
type ViciConfig struct {
	CAConfig    `vici:",inline"`
	Connections []ConnectionConfig `vici:",inline"`
}

// CAConfig defines the CA certificate configuration.
type CAConfig struct {
	CACertificate string `vici:"cacert"`
	File          string `vici:"file"`
	Handle        string `vici:"handle"`
	Slot          string `vici:"slot"`
	Module        string `vici:"module"`
	CertURIBase   string `vici:"cert_uri_base"`
	CRLURIs       string `vici:"crl_uris"`
	OCSPURIs      string `vici:"ocsp_uris"`
}

type Connections struct {
	Name string `vici:"inline"`
}

type Customconnections struct {
	Version    int    `vici:"version" json:"version"`
	LocalAddrs string `vici:"local_addrs"`
}

// ConnectionConfig defines the IKE connection configurations.
type ConnectionConfig struct {
	Version       int    `vici:"version"`
	LocalAddrs    string `vici:"local_addrs"`
	RemoteAddrs   string `vici:"remote_addrs"`
	LocalPort     int    `vici:"local_port"`
	RemotePort    int    `vici:"remote_port"`
	Proposals     string `vici:"proposals"`
	VIPs          string `vici:"vips"`
	Aggressive    bool   `vici:"aggressive"`
	Pull          bool   `vici:"pull"`
	DSCP          string `vici:"dscp"`
	Encap         bool   `vici:"encap"`
	Mobike        bool   `vici:"mobike"`
	DPDDelay      string `vici:"dpd_delay"`
	DPDTimeout    string `vici:"dpd_timeout"`
	Fragmentation string `vici:"fragmentation"`
	Childless     string `vici:"childless"`
	SendCertReq   bool   `vici:"send_certreq"`
	SendCert      string `vici:"send_cert"`
	PPKID         string `vici:"ppk_id"`
	PPKRequired   bool   `vici:"ppk_required"`
	Keyingtries   int    `vici:"keyingtries"`
	Unique        string `vici:"unique"`
	ReauthTime    string `vici:"reauth_time"`
	RekeyTime     string `vici:"rekey_time"`
	OverTime      string `vici:"over_time"`
	RandTime      string `vici:"rand_time"`
	Pools         string `vici:"pools"`
	IfIDIn        string `vici:"if_id_in"`
	IfIDOut       string `vici:"if_id_out"`
	OCSP          string `vici:"ocsp"`
	Mediation     bool   `vici:"mediation"`
	MediatedBy    string `vici:"mediated_by"`
	MediationPeer string `vici:"mediation_peer"`
}

// LocalAuthConfig defines the local authentication round configuration.
type LocalAuthConfig struct {
	Round   int        `vici:"round"`
	Auth    string     `vici:"auth"`
	ID      string     `vici:"id"`
	EapID   string     `vici:"eap_id"`
	AAAID   string     `vici:"aaa_id"`
	XAuthID string     `vici:"xauth_id"`
	Certs   string     `vici:"certs"`
	Cert    CertConfig `vici:",inline"`
	Pubkeys string     `vici:"pubkeys"`
}

// CertConfig defines the certificate configuration.
type CertConfig struct {
	File   string `vici:"file"`
	Handle string `vici:"handle"`
	Slot   string `vici:"slot"`
	Module string `vici:"module"`
}

// RemoteAuthConfig defines the remote authentication round configuration.
type RemoteAuthConfig struct {
	Round      int          `vici:"round"`
	Auth       string       `vici:"auth"`
	ID         string       `vici:"id"`
	EapID      string       `vici:"eap_id"`
	Groups     string       `vici:"groups"`
	CertPolicy string       `vici:"cert_policy"`
	Certs      string       `vici:"certs"`
	Cert       CertConfig   `vici:",inline"`
	CACerts    string       `vici:"cacerts"`
	CACert     CACertConfig `vici:",inline"`
	CAID       string       `vici:"ca_id"`
	Pubkeys    string       `vici:"pubkeys"`
	Revocation string       `vici:"revocation"`
}

// CACertConfig defines the CA certificate configuration.
type CACertConfig struct {
	File   string `vici:"file"`
	Handle string `vici:"handle"`
	Slot   string `vici:"slot"`
	Module string `vici:"module"`
}

// ChildConfig defines the CHILD SA configuration subsection.
type ChildConfig struct {
	AHProposals    string `vici:"ah_proposals"`
	ESPProposals   string `vici:"esp_proposals"`
	SHA256_96      bool   `vici:"sha256_96"`
	LocalTS        string `vici:"local_ts"`
	RemoteTS       string `vici:"remote_ts"`
	RekeyTime      string `vici:"rekey_time"`
	LifeTime       string `vici:"life_time"`
	RandTime       string `vici:"rand_time"`
	RekeyBytes     string `vici:"rekey_bytes"`
	LifeBytes      string `vici:"life_bytes"`
	RandBytes      string `vici:"rand_bytes"`
	RekeyPackets   string `vici:"rekey_packets"`
	LifePackets    string `vici:"life_packets"`
	RandPackets    string `vici:"rand_packets"`
	Updown         string `vici:"updown"`
	Hostaccess     bool   `vici:"hostaccess"`
	Mode           string `vici:"mode"`
	Policies       bool   `vici:"policies"`
	PoliciesFwdOut bool   `vici:"policies_fwd_out"`
	DPDAction      string `vici:"dpd_action"`
	IPComp         bool   `vici:"ipcomp"`
	Inactivity     string `vici:"inactivity"`
	Reqid          int    `vici:"reqid"`
	Priority       int    `vici:"priority"`
	Interface      string `vici:"interface"`
	MarkIn         string `vici:"mark_in"`
	MarkInSA       bool   `vici:"mark_in_sa"`
	MarkOut        string `vici:"mark_out"`
	SetMarkIn      string `vici:"set_mark_in"`
	SetMarkOut     string `vici:"set_mark_out"`
	IfIDIn         string `vici:"if_id_in"`
	IfIDOut        string `vici:"if_id_out"`
	Label          string `vici:"label"`
	LabelMode      string `vici:"label_mode"`
	TFCPadding     int    `vici:"tfc_padding"`
	ReplayWindow   int    `vici:"replay_window"`
	HWOffload      string `vici:"hw_offload"`
	CopyDF         bool   `vici:"copy_df"`
	CopyECN        bool   `vici:"copy_ecn"`
	CopyDSCP       string `vici:"copy_dscp"`
	StartAction    string `vici:"start_action"`
	CloseAction    string `vici:"close_action"`
}

// PoolConfig defines named pools for virtual IPs and other configuration attributes.
type PoolConfig struct {
	Addrs string `vici:"addrs"`
	Attr  string `vici:"attr"`
}

// SecretsConfig defines secrets for authentication and private key decryption.
type SecretsConfig struct {
	EAPSecrets        []EAPSecret        `vici:",inline"`
	XAuthSecrets      []XAuthSecret      `vici:",inline"`
	NTLMSecrets       []NTLMSecret       `vici:",inline"`
	IKESecrets        []IKESecret        `vici:",inline"`
	PPKSecrets        []PPKSecret        `vici:",inline"`
	PrivateKeySecrets []PrivateKeySecret `vici:",inline"`
	RSASecrets        []RSASecret        `vici:",inline"`
	ECDSASecrets      []ECDSASecret      `vici:",inline"`
	PKCS8Secrets      []PKCS8Secret      `vici:",inline"`
	PKCS12Secrets     []PKCS12Secret     `vici:",inline"`
	TokenSecrets      []TokenSecret      `vici:",inline"`
}

// EAPSecret defines the EAP secret subsection.
type EAPSecret struct {
	Secret string `vici:"secret"`
	ID     string `vici:"id"`
}

// XAuthSecret defines the XAuth secret subsection.
type XAuthSecret struct {
	Secret string `vici:"secret"`
	ID     string `vici:"id"`
}

// NTLMSecret defines the NTLM secret subsection.
type NTLMSecret struct {
	Secret string `vici:"secret"`
	ID     string `vici:"id"`
}

// IKESecret defines the IKE preshared secret section.
type IKESecret struct {
	Secret string `vici:"secret"`
	ID     string `vici:"id"`
}

// PPKSecret defines the Postquantum Preshared Key (PPK) subsection.
type PPKSecret struct {
	Secret string `vici:"secret"`
	ID     string `vici:"id"`
}

// PrivateKeySecret defines the private key decryption passphrase.
type PrivateKeySecret struct {
	File   string `vici:"file"`
	Secret string `vici:"secret"`
}

// RSASecret defines the RSA private key decryption passphrase.
type RSASecret struct {
	File   string `vici:"file"`
	Secret string `vici:"secret"`
}

// ECDSASecret defines the ECDSA private key decryption passphrase.
type ECDSASecret struct {
	File   string `vici:"file"`
	Secret string `vici:"secret"`
}

// PKCS8Secret defines the PKCS#8 private key decryption passphrase.
type PKCS8Secret struct {
	File   string `vici:"file"`
	Secret string `vici:"secret"`
}

// PKCS12Secret defines the PKCS#12 decryption passphrase.
type PKCS12Secret struct {
	File   string `vici:"file"`
	Secret string `vici:"secret"`
}

// TokenSecret defines a private key stored on a token, smartcard, or TPM 2.0.
type TokenSecret struct {
	Handle string `vici:"handle"`
	Slot   string `vici:"slot"`
	Module string `vici:"module"`
	PIN    string `vici:"pin"`
}
