package handler

import (
	"encoding/json"
	"io"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/tanan/wg-config-generator/domain"
)

func (h handler) GetClientList() ([]domain.ClientConfig, error) {
	clientWorkDir := filepath.Join(h.Config.WorkDir, "clients")
	files, err := os.ReadDir(clientWorkDir)
	if err != nil {
		return nil, err
	}

	var clientList []domain.ClientConfig

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

func (h handler) readClient(dir string, entry fs.DirEntry) (domain.ClientConfig, error) {
	file, err := os.Open(filepath.Join(dir, entry.Name()))
	if err != nil {
		return domain.ClientConfig{}, err
	}

	data, err := io.ReadAll(file)
	if err != nil {
		return domain.ClientConfig{}, err
	}
	defer file.Close()

	var client domain.ClientConfig
	if err := json.Unmarshal(data, &client); err != nil {
		return domain.ClientConfig{}, err
	}

	// TODO: read private key file

	return client, nil
}
