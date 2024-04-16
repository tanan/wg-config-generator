package handler

import (
	"os/exec"
	"strings"
)

func (h handler) CreatePrivateKey() (string, error) {
	privateKey, err := h.ExecCommand(exec.Command("wg", "genkey"), nil)
	if err != nil {
		return "", err
	}
	return strings.Trim(privateKey, "\n"), nil
}

func (h handler) CreatePreSharedKey() (string, error) {
	preSharedKey, err := h.ExecCommand(exec.Command("wg", "genpsk"), nil)
	if err != nil {
		return "", err
	}
	return strings.Trim(preSharedKey, "\n"), nil
}

func (h handler) CreatePublicKey(privateKey string) (string, error) {
	publicKey, err := h.ExecCommand(exec.Command("wg", "pubkey"), strings.NewReader(privateKey))
	if err != nil {
		return "", err
	}
	return strings.Trim(publicKey, "\n"), nil
}
