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
	clientWorkDir := filepath.Join(h.Config.WorkDir, "clients")
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
			DNS:        h.Config.Server.DNS,
			MTU:        h.Config.Server.MTU,
		},
		Peer: domain.Server{
			PublicKey:    h.Config.Server.PublicKey,
			PresharedKey: h.Config.Server.PresharedKey,
			AllowedIPs:   h.Config.Server.AllowedIPs,
			Endpoint:     h.Config.Server.Endpoint,
		},
	}, nil
}

func (h handler) CreateServerConfig(peers []domain.Client) (domain.ServerConfig, error) {
	privateKey, err := h.readPrivateKey(h.Config.Server.PrivateKeyFile)
	if err != nil {
		return domain.ServerConfig{}, nil
	}
	return domain.ServerConfig{
		ServerInterface: domain.ServerInterface{
			Address:    h.Config.Server.WireguardInterface.Address,
			ListenPort: h.Config.Server.Port,
			PrivateKey: privateKey,
			MTU:        h.Config.Server.MTU,
			PostUp:     h.Config.Server.PostUp,
			PostDown:   h.Config.Server.PostDown,
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
