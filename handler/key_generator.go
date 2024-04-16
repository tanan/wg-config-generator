package handler

import (
	"os/exec"
	"strings"
)

func (h handler) CreatePrivateKey() string {
	privateKey, _ := h.ExecCommand(exec.Command("wg", "genkey"), nil)
	return strings.Trim(privateKey, "\n")
}

func (h handler) CreatePreSharedKey() string {
	preSharedKey, _ := h.ExecCommand(exec.Command("wg", "genpsk"), nil)
	return strings.Trim(preSharedKey, "\n")
}

func (h handler) CreatePublicKey(privateKey string) string {
	publicKey, _ := h.ExecCommand(exec.Command("wg", "pubkey"), strings.NewReader(privateKey))
	return strings.Trim(publicKey, "\n")
}
