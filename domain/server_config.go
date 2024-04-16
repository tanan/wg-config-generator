package domain

type ServerConfig struct {
	ServerInterface ServerInterface
	Peers           []Client
}

type ServerInterface struct {
	Address    string
	ListenPort int
	PrivateKey string
	MTU        int
	PostUp     string
	PostDown   string
}

type Client struct {
	Name         string
	PublicKey    string
	PresharedKey string
	AllowedIPs   []string
}
