package handler

import (
	"encoding/json"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/tanan/wg-config-generator/model"
)

func (h handler) GetClientList() ([]model.ClientConfig, error) {
	clientWorkDir := h.getClientDir()
	files, err := os.ReadDir(clientWorkDir)
	if err != nil {
		return nil, err
	}

	var clientList []model.ClientConfig

	for _, entry := range files {
		if !entry.IsDir() {
			client, err := h.readClient(filepath.Join(clientWorkDir, entry.Name()))
			if err != nil {
				return nil, err
			}
			clientList = append(clientList, client)
		}
	}
	return clientList, nil
}

func (h handler) readClient(fn string) (model.ClientConfig, error) {
	file, err := os.Open(fn)
	if err != nil {
		return model.ClientConfig{}, err
	}

	data, err := io.ReadAll(file)
	if err != nil {
		return model.ClientConfig{}, err
	}
	defer file.Close()

	var client model.ClientConfig
	if err := json.Unmarshal(data, &client); err != nil {
		return model.ClientConfig{}, err
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
