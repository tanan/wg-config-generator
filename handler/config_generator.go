package handler

import (
	"fmt"

	"github.com/tanan/wg-config-generator/model"
)

func (h handler) CreateClientConfig(name string, address string) (model.ClientConfig, error) {
	privateKey, err := h.Command.CreatePrivateKey()
	if err != nil {
		return model.ClientConfig{}, fmt.Errorf("failed to create private key: %w", err)
	}
	publicKey, err := h.Command.CreatePublicKey(privateKey)
	if err != nil {
		return model.ClientConfig{}, fmt.Errorf("failed to create public key: %w", err)
	}
	presharedKey, err := h.Command.CreatePreSharedKey()
	if err != nil {
		return model.ClientConfig{}, fmt.Errorf("failed to create preshared key: %w", err)
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
		return model.ServerConfig{}, fmt.Errorf("failed to read server private key: %w", err)
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
