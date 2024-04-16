package domain

type ServerConfig struct {
	ServerInterface ServerInterface
	Peers           []Client
}

type ServerInterface struct {
	Address          string
	ListenPort       string
	ServerPrivateKey string
	MTU              string
	PostUp           string
	PostDown         string
}

type Client struct {
	ClientPublicKey string
	PresharedKey    string
	AllowedIPs      []string
}
