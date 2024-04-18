package handler

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/tanan/wg-config-generator/config"
	"github.com/tanan/wg-config-generator/model"
)

func readFile(gotFn string, wantFn string) (got string, want string, err error) {
	g, err := os.ReadFile(gotFn)
	if err != nil {
		return "", "", err
	}
	w, err := os.ReadFile(wantFn)
	if err != nil {
		return "", "", err
	}
	return string(g), string(w), nil
}

func Test_handler_WriteServerConfig(t *testing.T) {
	server := model.ServerConfig{
		Address:    "192.168.227.1",
		ListenPort: 51820,
		Endpoint:   "192.168.227.1",
		PrivateKey: "PrivateKey",
		PublicKey:  "PublicKey",
		PostUp:     "iptables -A FORWARD -i %!i(MISSING) -j ACCEPT; iptables -t nat -A POSTROUTING -o eth0 -j MASQUERADE",
		PostDown:   "iptables -D FORWARD -i %!i(MISSING) -j ACCEPT; iptables -t nat -D POSTROUTING -o eth0 -j MASQUERADE",
		DNS:        "",
		MTU:        1420,
		AllowedIPs: "192.168.227.0/24",
	}

	peers := []model.ClientConfig{
		{
			Name:         "peer1",
			Address:      "192.168.227.2",
			PrivateKey:   "peer1PrivateKey",
			PublicKey:    "peer1PublicKey",
			PresharedKey: "peer1PresharedKey",
		},
	}
	type args struct {
		server model.ServerConfig
		peers  []model.ClientConfig
	}
	tests := []struct {
		name         string
		args         args
		wantFileName string
		wantErr      bool
	}{
		{name: "success", args: args{server: server, peers: peers}, wantFileName: "wg0.success.conf", wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := handler{
				Config: config.Config{WorkDir: t.TempDir()},
			}

			if err := h.WriteServerConfig(tt.args.server, tt.args.peers); (err != nil) != tt.wantErr {
				t.Errorf("handler.WriteServerConfig() error = %v, wantErr %v", err, tt.wantErr)
			}

			got, want, err := readFile(filepath.Join("testdata", tt.wantFileName), filepath.Join(h.Config.WorkDir, fmt.Sprintf("%s.conf", WGInterfaceName)))
			if err != nil {
				t.Fatalf("handler.CreateWGServerConfig() error : %v", err)
			}

			if diff := cmp.Diff(want, got); diff != "" {
				t.Errorf(diff)
			}
		})
	}
}

func Test_handler_WriteClientConfig(t *testing.T) {
	type fields struct {
		Command Command
		Config  config.Config
	}
	type args struct {
		client model.ClientConfig
		server model.ServerConfig
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := handler{
				Command: tt.fields.Command,
				Config:  tt.fields.Config,
			}
			if err := h.WriteClientConfig(tt.args.client, tt.args.server); (err != nil) != tt.wantErr {
				t.Errorf("handler.WriteClientConfig() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_handler_WriteClientSecret(t *testing.T) {
	type fields struct {
		Command Command
		Config  config.Config
	}
	type args struct {
		client model.ClientConfig
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := handler{
				Command: tt.fields.Command,
				Config:  tt.fields.Config,
			}
			if err := h.WriteClientSecret(tt.args.client); (err != nil) != tt.wantErr {
				t.Errorf("handler.WriteClientSecret() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_handler_SendClientConfigByEmail(t *testing.T) {
	type fields struct {
		Command Command
		Config  config.Config
	}
	type args struct {
		client model.ClientConfig
		server model.ServerConfig
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := handler{
				Command: tt.fields.Command,
				Config:  tt.fields.Config,
			}
			if err := h.SendClientConfigByEmail(tt.args.client, tt.args.server); (err != nil) != tt.wantErr {
				t.Errorf("handler.SendClientConfigByEmail() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
