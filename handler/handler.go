package handler

import (
	"io"
	"log/slog"
	"os/exec"

	"github.com/tanan/wg-config-generator/config"
	"github.com/tanan/wg-config-generator/domain"
)

type Handler interface {
	GetClientList() ([]domain.Client, error)
	CreateClientConfig(name string, address string) (domain.ClientConfig, error)
	CreateServerConfig(peers []domain.Client) (domain.ServerConfig, error)
	WriteServerConfig(domain.ServerConfig) error
	WriteClientConfig(domain.ClientConfig) error
	SendClientConfigByEmail(domain.ClientConfig) error
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
