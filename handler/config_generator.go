package handler

import (
	"github.com/tanan/wg-config-generator/model"
)

func (h handler) CreateClientConfig(name string, address string) (model.ClientConfig, error) {
	privateKey, err := h.Command.CreatePrivateKey()
	if err != nil {
		return model.ClientConfig{}, err
	}
	publicKey, err := h.Command.CreatePublicKey(privateKey)
	if err != nil {
		return model.ClientConfig{}, err
	}
	presharedKey, err := h.Command.CreatePreSharedKey()
	if err != nil {
		return model.ClientConfig{}, err
	}

	clientConfig := model.ClientConfig{
		Name:         name,
		Address:      address,
		PrivateKey:   privateKey,
		PublicKey:    publicKey,
		PresharedKey: presharedKey,
	}

	return clientConfig, nil
}

func (h handler) CreateServerConfig() (model.ServerConfig, error) {
	privateKey, err := h.readPrivateKey(h.Config.Server.PrivateKeyFile)
	if err != nil {
		return model.ServerConfig{}, err
	}
	return model.ServerConfig{
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
