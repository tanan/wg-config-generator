package domain

type ClientConfig struct {
	Name            string
	ClientInterface ClientInterface
	Peer            Server
}

type ClientInterface struct {
	Address    string
	PrivateKey string
	DNS        string
	MTU        int
}

type Server struct {
	PublicKey    string
	PresharedKey string
	AllowedIPs   string
	Endpoint     string
}
