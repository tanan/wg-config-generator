package model

type ServerConfig struct {
	Address    string
	ListenPort int
	Endpoint   string
	PrivateKey string
	PublicKey  string
	PostUp     string
	PostDown   string
	DNS        string
	MTU        int
	AllowedIPs string
}
