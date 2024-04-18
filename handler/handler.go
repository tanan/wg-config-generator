package handler

import (
	"io"
	"log/slog"
	"os/exec"

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

type Command interface {
	ExecCommand(cmd *exec.Cmd, stdin io.Reader) (string, error)
}

type command struct{}

func NewCommand() Command {
	return &command{}
}

func (h command) ExecCommand(cmd *exec.Cmd, stdin io.Reader) (string, error) {
	cmd.Stdin = stdin
	out, err := cmd.CombinedOutput()
	if err != nil {
		slog.Error("failed to exec command", slog.String("command", cmd.String()))
		return "", err
	}
	return string(out), err
}
