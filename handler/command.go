package handler

import (
	"io"
	"log/slog"
	"os/exec"
	"strings"
)

type Command interface {
	CreatePrivateKey() (string, error)
	CreatePreSharedKey() (string, error)
	CreatePublicKey(privateKey string) (string, error)
}

type command struct{}

func NewCommand() Command {
	return &command{}
}

func (h command) execCommand(cmd *exec.Cmd, stdin io.Reader) (string, error) {
	cmd.Stdin = stdin
	out, err := cmd.CombinedOutput()
	if err != nil {
		slog.Error("failed to exec command", slog.String("command", cmd.String()))
		return "", err
	}
	return string(out), err
}

func (c command) CreatePrivateKey() (string, error) {
	privateKey, err := c.execCommand(exec.Command("wg", "genkey"), nil)
	if err != nil {
		return "", err
	}
	return strings.Trim(privateKey, "\n"), nil
}

func (c command) CreatePreSharedKey() (string, error) {
	preSharedKey, err := c.execCommand(exec.Command("wg", "genpsk"), nil)
	if err != nil {
		return "", err
	}
	return strings.Trim(preSharedKey, "\n"), nil
}

func (c command) CreatePublicKey(privateKey string) (string, error) {
	publicKey, err := c.execCommand(exec.Command("wg", "pubkey"), strings.NewReader(privateKey))
	if err != nil {
		return "", err
	}
	return strings.Trim(publicKey, "\n"), nil
}
