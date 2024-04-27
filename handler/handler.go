package handler

import (
	"path/filepath"

	"github.com/tanan/wg-config-generator/config"
	"github.com/tanan/wg-config-generator/model"
)

const (
	ClientDir       = "clients"
	SecretDir       = "secrets"
	WGInterfaceName = "wg0"
)

type Handler interface {
	GetClientList() ([]model.ClientConfig, error)
	CreateClientConfig(name string, address string) (model.ClientConfig, error)
	SaveClientSetting(client model.ClientConfig) error
	CreateServerConfig() (model.ServerConfig, error)
	WriteServerConfig(server model.ServerConfig, peers []model.ClientConfig) error
	WriteClientConfig(client model.ClientConfig, server model.ServerConfig) error
	SendClientConfigByEmail(client model.ClientConfig, server model.ServerConfig) error
}

type handler struct {
	Command
	Config config.Config
}

func NewHandler(cmd Command, cfg config.Config) Handler {
	return &handler{
		Command: cmd,
		Config:  cfg,
	}
}

func (h handler) getWorkDir() string {
	return h.Config.WorkDir
}

func (h handler) getClientDir() string {
	return filepath.Join(h.Config.WorkDir, ClientDir)
}

func (h handler) getClientSecretDir() string {
	return filepath.Join(h.Config.WorkDir, ClientDir, SecretDir)
}
