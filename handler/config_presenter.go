package handler

import (
	"fmt"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/tanan/wg-config-generator/domain"
	"github.com/tanan/wg-config-generator/utils"
)

func (h handler) WriteServerConfig(server domain.ServerConfig, peers []domain.ClientConfig) error {
	err := utils.Makedir(h.Config.WorkDir, 0700)
	if err != nil {
		return err
	}
	f, err := utils.CreateFile(filepath.Join(h.Config.WorkDir, fmt.Sprintf("%s.conf", WGInterfaceName)), 0600)
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

func (h handler) WriteClientConfig(client domain.ClientConfig, server domain.ServerConfig) error {
	// create a client profile to current path
	f, err := utils.CreateFile(fmt.Sprintf("%s.conf", client.Name), 0600)
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

func (h handler) WriteClientSecret(client domain.ClientConfig) error {
	err := utils.Makedir(filepath.Join(h.Config.WorkDir, SecretDir), 0700)
	if err != nil {
		return err
	}
	f, err := utils.CreateFile(filepath.Join(h.Config.WorkDir, SecretDir, fmt.Sprintf("%s.secret", client.Name)), 0600)
	if err != nil {
		return err
	}
	defer f.Close()

	f.WriteString(client.PrivateKey)
	return nil
}

// TODO
func (h handler) SendClientConfigByEmail(client domain.ClientConfig, server domain.ServerConfig) error {
	return nil
}
