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
	MTU        string
}

type Server struct {
	ServerPublicKey string
	PresharedKey    string
	AllowedIPs      string
	Endpoint        string
}
