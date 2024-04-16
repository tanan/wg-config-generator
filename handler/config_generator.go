package handler

import (
	"github.com/tanan/wg-config-generator/config"
	"github.com/tanan/wg-config-generator/domain"
)

// read client files in work dir
// create client list using client files
func (h handler) GetClientList() ([]domain.Client, error) {
	return nil, nil
}

// create private key and preshared key
// set address, dns, mtu via args
func (h handler) CreateClientConfig(name string, address string, cfg config.Config) (domain.ClientConfig, error) {

	return domain.ClientConfig{
		ClientInterface: domain.ClientInterface{
			Address:    "",
			PrivateKey: "",
			DNS:        "",
			MTU:        "",
		},
		Peer: domain.Server{
			ServerPublicKey: "",
			PresharedKey:    "",
			AllowedIPs:      "",
			Endpoint:        "",
		},
	}, nil
}

// create key
// set interface info using cfg
func (h handler) CreateServerConfig(peers []domain.Client, cfg config.Config) (domain.ServerConfig, error) {
	return domain.ServerConfig{
		ServerInterface: domain.ServerInterface{
			Address:          "",
			ListenPort:       "",
			ServerPrivateKey: "",
			MTU:              "",
			PostUp:           "",
			PostDown:         "",
		},
		Peers: peers,
	}, nil
}
