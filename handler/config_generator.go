package handler

import (
	"encoding/json"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/tanan/wg-config-generator/domain"
)

// read client files in work dir
// create client list using client files
func (h handler) GetClientList() ([]domain.Client, error) {
	clientWorkDir := filepath.Join(h.WorkDir, "clients")
	files, err := os.ReadDir(clientWorkDir)
	if err != nil {
		return nil, err
	}

	var clientList []domain.Client

	for _, entry := range files {
		if !entry.IsDir() {
			client, err := h.readClient(clientWorkDir, entry)
			if err != nil {
				return nil, err
			}
			clientList = append(clientList, client)
		}
	}
	return clientList, nil
}

func (h handler) CreateClientConfig(name string, address string) (domain.ClientConfig, error) {
	privateKey, err := h.CreatePrivateKey()
	if err != nil {
		return domain.ClientConfig{}, err
	}

	return domain.ClientConfig{
		Name: name,
		ClientInterface: domain.ClientInterface{
			Address:    address,
			PrivateKey: privateKey,
			DNS:        h.Cfg.Server.DNS,
			MTU:        h.Cfg.Server.MTU,
		},
		Peer: domain.Server{
			ServerPublicKey: h.Cfg.Server.PublicKey,
			PresharedKey:    h.Cfg.Server.PresharedKey,
			AllowedIPs:      h.Cfg.Server.AllowedIPs,
			Endpoint:        h.Cfg.Server.Endpoint,
		},
	}, nil
}

func (h handler) CreateServerConfig(peers []domain.Client) (domain.ServerConfig, error) {
	privateKey, err := h.readPrivateKey(h.Cfg.Server.PrivateKeyFile)
	if err != nil {
		return domain.ServerConfig{}, nil
	}
	return domain.ServerConfig{
		ServerInterface: domain.ServerInterface{
			Address:          h.Cfg.Server.WireguardInterface.Address,
			ListenPort:       h.Cfg.Server.Port,
			ServerPrivateKey: privateKey,
			MTU:              h.Cfg.Server.MTU,
			PostUp:           h.Cfg.Server.PostUp,
			PostDown:         h.Cfg.Server.PostDown,
		},
		Peers: peers,
	}, nil
}

func (h handler) readClient(dir string, entry fs.DirEntry) (domain.Client, error) {
	file, err := os.Open(filepath.Join(dir, entry.Name()))
	if err != nil {
		return domain.Client{}, err
	}

	data, err := io.ReadAll(file)
	if err != nil {
		return domain.Client{}, err
	}
	defer file.Close()

	var client domain.Client
	if err := json.Unmarshal(data, &client); err != nil {
		return domain.Client{}, err
	}

	return client, nil
}

func (h handler) readPrivateKey(fn string) (string, error) {
	f, err := os.Open(fn)
	if err != nil {
		return "", err
	}
	defer f.Close()

	data, err := io.ReadAll(f)
	if err != nil {
		return "", err
	}
	return strings.Trim(string(data), " \t\n"), nil
}
