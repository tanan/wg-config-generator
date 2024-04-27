package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/tanan/wg-config-generator/model"
	"github.com/tanan/wg-config-generator/util"
)

func (h handler) GetClientList() ([]model.ClientConfig, error) {
	clientWorkDir := h.getClientDir()
	files, err := os.ReadDir(clientWorkDir)
	if err != nil {
		return nil, fmt.Errorf("func ReadDir error: %w", err)
	}

	var clientList []model.ClientConfig

	for _, entry := range files {
		if !entry.IsDir() {
			client, err := h.readClient(filepath.Join(clientWorkDir, entry.Name()))
			if err != nil {
				return nil, fmt.Errorf("failed to read client: %w", err)
			}
			clientList = append(clientList, client)
		}
	}
	return clientList, nil
}

func (h handler) readClient(fn string) (model.ClientConfig, error) {
	f, err := os.Open(fn)
	if err != nil {
		return model.ClientConfig{}, util.NewFileError(util.FileNotFound, fn, err)
	}

	data, err := io.ReadAll(f)
	if err != nil {
		return model.ClientConfig{}, util.NewFileError(util.FileReadFailure, fn, err)
	}
	defer f.Close()

	var client model.ClientConfig
	if err := json.Unmarshal(data, &client); err != nil {
		return model.ClientConfig{}, util.NewFileError(util.FileUnmarshalFailure, fn, err)
	}

	return client, nil
}

func (h handler) readPrivateKey(fn string) (string, error) {
	f, err := os.Open(fn)
	if err != nil {
		return "", util.NewFileError(util.FileNotFound, fn, err)
	}
	defer f.Close()

	data, err := io.ReadAll(f)
	if err != nil {
		return "", util.NewFileError(util.FileReadFailure, fn, err)
	}
	return strings.Trim(string(data), " \t\n"), nil
}
