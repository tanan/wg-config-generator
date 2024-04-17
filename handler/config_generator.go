package handler

import (
	"encoding/json"
	"fmt"
	"path/filepath"

	"github.com/tanan/wg-config-generator/domain"
	"github.com/tanan/wg-config-generator/utils"
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

	clientConfig := domain.ClientConfig{
		Name:         name,
		Address:      address,
		PrivateKey:   privateKey,
		PublicKey:    publicKey,
		PresharedKey: presharedKey,
	}

	if err := h.saveClientConfig(clientConfig); err != nil {
		return domain.ClientConfig{}, err
	}

	return clientConfig, nil
}

func (h handler) saveClientConfig(clientConfig domain.ClientConfig) error {
	if err := utils.Makedir(filepath.Join(h.Config.WorkDir, ClientDir), 0700); err != nil {
		return err
	}
	f, err := utils.CreateFile(filepath.Join(h.Config.WorkDir, ClientDir, fmt.Sprintf("%s.json", clientConfig.Name)), 0600)
	if err != nil {
		return err
	}
	defer f.Close()

	data, _ := json.MarshalIndent(clientConfig, "", "    ")
	f.Write(data)
	return nil
}

func (h handler) CreateServerConfig() (domain.ServerConfig, error) {
	privateKey, err := h.readPrivateKey(h.Config.Server.PrivateKeyFile)
	if err != nil {
		return domain.ServerConfig{}, err
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
