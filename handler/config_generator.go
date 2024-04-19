package handler

import (
	"encoding/json"
	"fmt"
	"path/filepath"

	"github.com/tanan/wg-config-generator/model"
	"github.com/tanan/wg-config-generator/util"
)

func (h handler) CreateClientConfig(name string, address string) (model.ClientConfig, error) {
	privateKey, err := h.createPrivateKey()
	if err != nil {
		return model.ClientConfig{}, err
	}
	publicKey, err := h.createPublicKey(privateKey)
	if err != nil {
		return model.ClientConfig{}, err
	}
	presharedKey, err := h.createPreSharedKey()
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

	if err := h.saveClientConfig(clientConfig); err != nil {
		return model.ClientConfig{}, err
	}

	return clientConfig, nil
}

func (h handler) saveClientConfig(clientConfig model.ClientConfig) error {
	if err := util.Makedir(filepath.Join(h.Config.WorkDir, ClientDir), 0700); err != nil {
		return err
	}
	f, err := util.CreateFile(filepath.Join(h.Config.WorkDir, ClientDir, fmt.Sprintf("%s.json", clientConfig.Name)), 0600)
	if err != nil {
		return err
	}
	defer f.Close()

	data, _ := json.MarshalIndent(clientConfig, "", "    ")
	f.Write(data)
	return nil
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
