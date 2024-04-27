package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/tanan/wg-config-generator/model"
	"github.com/tanan/wg-config-generator/util"
)

func (h handler) WriteServerConfig(server model.ServerConfig, peers []model.ClientConfig) error {
	workDir := h.getWorkDir()
	err := util.Makedir(workDir, 0700)
	if err != nil {
		return util.NewFileError(util.PermissionDenied, workDir, err)
	}
	confPath := filepath.Join(workDir, fmt.Sprintf("%s.conf", WGInterfaceName))
	f, err := util.CreateFile(confPath, 0600)
	if err != nil {
		return util.NewFileError(util.PermissionDenied, confPath, err)
	}

	if err := h.writeServerConfig(f, server, peers); err != nil {
		return util.NewFileError(util.FileWriteFailure, confPath, err)
	}

	return nil
}

func (h handler) WriteClientConfig(client model.ClientConfig, server model.ServerConfig) error {
	// write client config in secret dir since client profile includes a private key
	secretDir := h.getClientSecretDir()
	err := util.Makedir(secretDir, 0700)
	if err != nil {
		return util.NewFileError(util.PermissionDenied, secretDir, err)
	}
	confPath := filepath.Join(secretDir, fmt.Sprintf("%s.conf", client.Name))
	f, err := util.CreateFile(confPath, 0600)
	if err != nil {
		return util.NewFileError(util.PermissionDenied, confPath, err)
	}
	defer f.Close()

	if err := h.writeClientConfig(f, client, server); err != nil {
		return util.NewFileError(util.FileWriteFailure, confPath, err)
	}

	return nil
}

func (h handler) SaveClientSetting(cc model.ClientConfig) error {
	workDir := h.getWorkDir()
	if err := util.Makedir(workDir, 0700); err != nil {
		return util.NewFileError(util.PermissionDenied, workDir, err)
	}
	confPath := filepath.Join(workDir, fmt.Sprintf("%s.json", cc.Name))
	f, err := util.CreateFile(confPath, 0600)
	if err != nil {
		return util.NewFileError(util.PermissionDenied, confPath, err)
	}
	defer f.Close()

	if err := h.saveClientSetting(f, cc); err != nil {
		return util.NewFileError(util.FileWriteFailure, confPath, err)
	}
	return nil
}

func (h handler) writeServerConfig(w io.Writer, server model.ServerConfig, peers []model.ClientConfig) error {
	ew := util.NewErrorWriter(w)
	var row []string
	row = append(row, "[Interface]")
	row = append(row, fmt.Sprintf("Address = %s", server.Address))
	row = append(row, fmt.Sprintf("ListenPort = %s", strconv.Itoa(server.ListenPort)))
	row = append(row, fmt.Sprintf("PrivateKey = %s", server.PrivateKey))
	row = append(row, fmt.Sprintf("MTU = %s", strconv.Itoa(server.MTU)))
	row = append(row, fmt.Sprintf("PostUp = %s", server.PostUp))
	row = append(row, fmt.Sprintf("PostDown = %s", server.PostDown))
	ew.Write([]byte(strings.Join(row, "\n")))
	ew.Write([]byte("\n"))

	for _, peer := range peers {
		var row []string
		ew.Write([]byte("\n"))
		row = append(row, fmt.Sprintf("# %s", peer.Name))
		row = append(row, "[Peer]")
		row = append(row, fmt.Sprintf("PublicKey = %s", peer.PublicKey))
		row = append(row, fmt.Sprintf("PresharedKey = %s", peer.PresharedKey))
		row = append(row, fmt.Sprintf("AllowedIPs = %s/32", peer.Address))
		ew.Write([]byte(strings.Join(row, "\n")))
		ew.Write([]byte("\n"))
	}

	if ew.Err != nil {
		return ew.Err
	}

	return nil
}

func (h handler) saveClientSetting(w io.Writer, cc model.ClientConfig) error {
	ew := util.NewErrorWriter(w)
	data, _ := json.MarshalIndent(cc, "", "    ")
	ew.Write(data)
	if ew.Err != nil {
		return ew.Err
	}
	return nil
}

func (h handler) writeClientConfig(w io.Writer, client model.ClientConfig, server model.ServerConfig) error {
	ew := util.NewErrorWriter(w)
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

	ew.Write([]byte(strings.Join(row, "\n")))
	ew.Write([]byte("\n"))
	if ew.Err != nil {
		return ew.Err
	}
	return nil
}

// TODO
func (h handler) SendClientConfigByEmail(client model.ClientConfig, server model.ServerConfig) error {
	return nil
}
