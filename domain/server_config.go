package domain

type ServerConfig struct {
	ServerInterface ServerInterface
	Peers           []Client
}

type ServerInterface struct {
	Address          string
	ListenPort       int
	ServerPrivateKey string
	MTU              string
	PostUp           string
	PostDown         string
}

type Client struct {
	Name            string
	ClientPublicKey string
	PresharedKey    string
	AllowedIPs      []string
}
