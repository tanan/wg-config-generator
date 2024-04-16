package handler

import (
	"errors"
	"io"
	"os/exec"
	"reflect"
	"strings"
	"testing"
)

type MockCommand struct {
	ExecCommandFunc func(cmd *exec.Cmd, stdin io.Reader) (string, error)
}

func (m *MockCommand) ExecCommand(cmd *exec.Cmd, stdin io.Reader) (string, error) {
	return m.ExecCommandFunc(cmd, stdin)
}

func isEqualCmd(cmd1, cmd2 *exec.Cmd) bool {
	if cmd1.Path != cmd2.Path || !reflect.DeepEqual(cmd1.Args, cmd2.Args) {
		return false
	}
	return true
}

func Test_handler_CreatePrivateKey(t *testing.T) {
	key := "privatekey"
	mockCmd := &MockCommand{
		ExecCommandFunc: func(cmd *exec.Cmd, stdin io.Reader) (string, error) {
			if isEqualCmd(exec.Command("wg", "genkey"), cmd) {
				return key, nil
			}
			return "", errors.New("command not found")
		},
	}

	tests := []struct {
		name    string
		want    string
		wantErr bool
	}{
		{name: "ok", want: key, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := handler{
				Command: mockCmd,
				WorkDir: "dummy",
			}
			got, err := h.CreatePrivateKey()
			if (err != nil) != tt.wantErr {
				t.Errorf("handler.CreatePrivateKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("handler.CreatePrivateKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_handler_CreatePreSharedKey(t *testing.T) {
	key := "presharedkey"
	mockCmd := &MockCommand{
		ExecCommandFunc: func(cmd *exec.Cmd, stdin io.Reader) (string, error) {
			if isEqualCmd(exec.Command("wg", "genpsk"), cmd) {
				return key, nil
			}
			return "", errors.New("command not found")
		},
	}
	tests := []struct {
		name    string
		want    string
		wantErr bool
	}{
		{name: "ok", want: key, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := handler{
				Command: mockCmd,
				WorkDir: "dummy",
			}
			got, err := h.CreatePreSharedKey()
			if (err != nil) != tt.wantErr {
				t.Errorf("handler.CreatePreSharedKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("handler.CreatePreSharedKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_handler_CreatePublicKey(t *testing.T) {
	pubKey := "publickey"
	privateKey := "privatekey"
	mockCmd := &MockCommand{
		ExecCommandFunc: func(cmd *exec.Cmd, stdin io.Reader) (string, error) {
			if isEqualCmd(exec.Command("wg", "pubkey"), cmd) && reflect.DeepEqual(strings.NewReader(privateKey), stdin) {
				return pubKey, nil
			}
			return "", errors.New("command not found")
		},
	}
	tests := []struct {
		name       string
		privateKey string
		want       string
		wantErr    bool
	}{
		{name: "ok", privateKey: privateKey, want: pubKey, wantErr: false},
		{name: "ng", privateKey: "dummy", want: "", wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := handler{
				Command: mockCmd,
				WorkDir: "dummy",
			}
			got, err := h.CreatePublicKey(tt.privateKey)
			if (err != nil) != tt.wantErr {
				t.Errorf("handler.CreatePublicKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("handler.CreatePublicKey() = %v, want %v", got, tt.want)
			}
		})
	}
}
