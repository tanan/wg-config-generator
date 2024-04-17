package handler

import (
	"github.com/tanan/wg-config-generator/domain"
)

func (h handler) CreateClientConfig(name string, address string) (domain.ClientConfig, error) {
	privateKey, err := h.CreatePrivateKey()
	if err != nil {
		return domain.ClientConfig{}, err
	}
	publicKey, err := h.CreatePublicKey(privateKey)
	if err != nil {
		return domain.ClientConfig{}, err
	}
	presharedKey, err := h.CreatePreSharedKey()
	if err != nil {
		return domain.ClientConfig{}, err
	}

	return domain.ClientConfig{
		Name:         name,
		Address:      address,
		PrivateKey:   privateKey,
		PublicKey:    publicKey,
		PresharedKey: presharedKey,
	}, nil
}

func (h handler) CreateServerConfig() (domain.ServerConfig, error) {
	privateKey, err := h.readPrivateKey(h.Config.Server.PrivateKeyFile)
	if err != nil {
		return domain.ServerConfig{}, nil
	}
	return domain.ServerConfig{
		Address:    h.Config.Server.Address,
		ListenPort: h.Config.Server.Port,
		Endpoint:   h.Config.Server.Endpoint,
		PrivateKey: privateKey,
		PublicKey:  h.Config.Server.PublicKey,
		PostUp:     h.Config.Server.PostUp,
		PostDown:   h.Config.Server.PostDown,
		DNS:        h.Config.Server.DNS,
		MTU:        h.Config.Server.MTU,
		AllowedIPs: h.Config.Server.AllowedIPs,
	}, nil
}
