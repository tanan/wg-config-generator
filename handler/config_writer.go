package handler

import (
	"encoding/json"
	"fmt"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/tanan/wg-config-generator/model"
	"github.com/tanan/wg-config-generator/util"
)

func (h handler) WriteServerConfig(server model.ServerConfig, peers []model.ClientConfig) error {
	err := util.Makedir(h.getWorkDir(), 0700)
	if err != nil {
		return err
	}
	f, err := util.CreateFile(filepath.Join(h.getWorkDir(), fmt.Sprintf("%s.conf", WGInterfaceName)), 0600)
	if err != nil {
		return err
	}

	var row []string
	row = append(row, "[Interface]")
	row = append(row, fmt.Sprintf("Address = %s", server.Address))
	row = append(row, fmt.Sprintf("ListenPort = %s", strconv.Itoa(server.ListenPort)))
	row = append(row, fmt.Sprintf("PrivateKey = %s", server.PrivateKey))
	row = append(row, fmt.Sprintf("MTU = %s", strconv.Itoa(server.MTU)))
	row = append(row, fmt.Sprintf("PostUp = %s", server.PostUp))
	row = append(row, fmt.Sprintf("PostDown = %s", server.PostDown))
	f.WriteString(strings.Join(row, "\n"))

	f.WriteString("\n")

	for _, peer := range peers {
		var row []string
		f.WriteString("\n")
		row = append(row, fmt.Sprintf("# %s", peer.Name))
		row = append(row, "[Peer]")
		row = append(row, fmt.Sprintf("PublicKey = %s", peer.PublicKey))
		row = append(row, fmt.Sprintf("PresharedKey = %s", peer.PresharedKey))
		row = append(row, fmt.Sprintf("AllowedIPs = %s/32", peer.Address))
		f.WriteString(strings.Join(row, "\n"))
		f.WriteString("\n")
	}

	return nil
}

func (h handler) SaveClientConfig(clientConfig model.ClientConfig) error {
	if err := util.Makedir(h.getClientDir(), 0700); err != nil {
		return err
	}
	f, err := util.CreateFile(filepath.Join(h.getClientDir(), fmt.Sprintf("%s.json", clientConfig.Name)), 0600)
	if err != nil {
		return err
	}
	defer f.Close()

	data, _ := json.MarshalIndent(clientConfig, "", "    ")
	f.Write(data)
	return nil
}

func (h handler) WriteClientConfig(client model.ClientConfig, server model.ServerConfig) error {
	// write client config in secret dir since client profile includes a private key
	err := util.Makedir(h.getClientSecretDir(), 0700)
	if err != nil {
		return err
	}
	f, err := util.CreateFile(filepath.Join(h.getClientSecretDir(), fmt.Sprintf("%s.conf", client.Name)), 0600)
	if err != nil {
		return err
	}
	defer f.Close()

	var row []string
	row = append(row, "[Interface]")
	row = append(row, fmt.Sprintf("Address = %s", client.Address))
	row = append(row, fmt.Sprintf("PrivateKey = %s", client.PrivateKey))
	row = append(row, fmt.Sprintf("DNS = %s", server.DNS))
	row = append(row, fmt.Sprintf("MTU = %s", strconv.Itoa(server.MTU)))
	row = append(row, "")
	row = append(row, "[Peer]")
	row = append(row, fmt.Sprintf("PublicKey = %s", server.PublicKey))
	row = append(row, fmt.Sprintf("PresharedKey = %s", client.PresharedKey))
	row = append(row, fmt.Sprintf("AllowedIPs = %s", server.AllowedIPs))
	row = append(row, fmt.Sprintf("Endpoint = %s", server.Endpoint))

	f.WriteString(strings.Join(row, "\n"))
	f.WriteString("\n")
	return nil
}

// TODO
func (h handler) SendClientConfigByEmail(client model.ClientConfig, server model.ServerConfig) error {
	return nil
}
