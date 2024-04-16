package handler

import (
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/tanan/wg-config-generator/domain"
)

func (h handler) WriteServerConfig(server domain.ServerConfig) error {
	path := filepath.Join(h.Config.WorkDir, "wg0.conf")
	f, err := os.Create(path)
	if err != nil {
		slog.Error(fmt.Sprintf("failed to create file : %s", path))
		return err
	}

	var row []string
	row = append(row, "[Interface]")
	row = append(row, fmt.Sprintf("Address = %s", server.ServerInterface.Address))
	row = append(row, fmt.Sprintf("ListenPort = %s", strconv.Itoa(server.ServerInterface.ListenPort)))
	row = append(row, fmt.Sprintf("PrivateKey = %s", server.ServerInterface.PrivateKey))
	row = append(row, fmt.Sprintf("MTU = %s", strconv.Itoa(server.ServerInterface.MTU)))
	row = append(row, fmt.Sprintf("PostUp = %s", server.ServerInterface.PostUp))
	row = append(row, fmt.Sprintf("PostDown = %s", server.ServerInterface.PostDown))
	f.Write([]byte(strings.Join(row, "\n")))

	f.Write([]byte("\n\n"))

	for _, peer := range server.Peers {
		var row []string
		row = append(row, "[Peer]")
		row = append(row, fmt.Sprintf("PublicKey = %s", peer.PublicKey))
		row = append(row, fmt.Sprintf("PresharedKey = %s", peer.PresharedKey))
		row = append(row, fmt.Sprintf("AllowedIPs = %s", peer.AllowedIPs))
		f.Write([]byte(strings.Join(row, "\n")))
	}

	f.Write([]byte("\n"))

	return nil
}

func (h handler) WriteClientConfig(client domain.ClientConfig) error {
	path := filepath.Join(h.Config.WorkDir, fmt.Sprintf("%s.conf", client.Name))
	f, err := os.Create(path)
	if err != nil {
		slog.Error(fmt.Sprintf("failed to create file : %s", path))
		return err
	}

	var row []string
	row = append(row, "[Interface]")
	row = append(row, fmt.Sprintf("Address = %s", client.ClientInterface.Address))
	row = append(row, fmt.Sprintf("PrivateKey = %s", client.ClientInterface.PrivateKey))
	row = append(row, fmt.Sprintf("DNS = %s", client.ClientInterface.DNS))
	row = append(row, fmt.Sprintf("MTU = %s", strconv.Itoa(client.ClientInterface.MTU)))
	row = append(row, "")
	row = append(row, "[Peer]")
	row = append(row, fmt.Sprintf("PublicKey = %s", client.Peer.PublicKey))
	row = append(row, fmt.Sprintf("PresharedKey = %s", client.Peer.PresharedKey))
	row = append(row, fmt.Sprintf("AllowedIPs = %s", client.Peer.AllowedIPs))
	row = append(row, fmt.Sprintf("Endpoint = %s", client.Peer.Endpoint))

	f.Write([]byte(strings.Join(row, "\n")))
	f.Write([]byte("\n"))
	return nil
}

func (h handler) SendClientConfigByEmail(domain.ClientConfig) error {
	return nil
}
