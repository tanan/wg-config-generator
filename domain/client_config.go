package domain

type ClientConfig struct {
	Name         string
	Address      string
	PrivateKey   string `json:"-"`
	PublicKey    string
	PresharedKey string
}
