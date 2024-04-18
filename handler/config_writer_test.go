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

func readFiles(gotFn string, wantFn string) (got string, want string, err error) {
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
	server1 := model.ServerConfig{
		Address:    "192.168.227.1",
		ListenPort: 51820,
		PrivateKey: "PrivateKey",
		PublicKey:  "PublicKey",
		PostUp:     "iptables -A FORWARD -i %i -j ACCEPT; iptables -t nat -A POSTROUTING -o eth0 -j MASQUERADE",
		PostDown:   "iptables -D FORWARD -i %i -j ACCEPT; iptables -t nat -D POSTROUTING -o eth0 -j MASQUERADE",
		MTU:        1420,
		AllowedIPs: "192.168.227.0/24",
	}

	server2 := model.ServerConfig{
		Address:    "192.168.228.1",
		ListenPort: 51821,
		PrivateKey: "PrivateKey2",
		PublicKey:  "PublicKey2",
		PostUp:     "iptables -A FORWARD -i %i -j ACCEPT; iptables -t nat -A POSTROUTING -o eth1 -j MASQUERADE",
		PostDown:   "iptables -D FORWARD -i %i -j ACCEPT; iptables -t nat -D POSTROUTING -o eth1 -j MASQUERADE",
		MTU:        1400,
		AllowedIPs: "192.168.228.0/24",
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
		{name: "ok", args: args{server: server1, peers: peers}, wantFileName: "wg0.conf.1", wantErr: false},
		{name: "ng", args: args{server: server2, peers: peers}, wantFileName: "wg0.conf.2", wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := handler{
				Config: config.Config{WorkDir: t.TempDir()},
			}

			if err := h.WriteServerConfig(tt.args.server, tt.args.peers); (err != nil) != tt.wantErr {
				t.Errorf("handler.WriteServerConfig() error = %v, wantErr %v", err, tt.wantErr)
			}

			got, want, err := readFiles(filepath.Join(h.Config.WorkDir, fmt.Sprintf("%s.conf", WGInterfaceName)), filepath.Join("testdata", tt.wantFileName))
			if err != nil {
				t.Fatalf("readFile() error : %v", err)
			}

			if diff := cmp.Diff(want, got); diff != "" {
				t.Error(diff)
			}
		})
	}
}

func Test_handler_WriteClientConfig(t *testing.T) {
	server := model.ServerConfig{
		Endpoint:   "192.168.227.1:51820",
		PublicKey:  "PublicKey",
		MTU:        1420,
		DNS:        "10.2.0.8",
		AllowedIPs: "192.168.227.0/24",
	}

	client1 := model.ClientConfig{
		Name:         "client1",
		Address:      "10.1.0.2",
		PrivateKey:   "PrivateKey",
		PresharedKey: "PresharedKey",
	}

	client2 := model.ClientConfig{
		Name:         "client2",
		Address:      "10.1.0.3",
		PrivateKey:   "PrivateKey2",
		PresharedKey: "PresharedKey2",
	}

	type args struct {
		client model.ClientConfig
		server model.ServerConfig
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "client1", args: args{client: client1, server: server}, wantErr: false},
		{name: "client2", args: args{client: client2, server: server}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := handler{
				Config: config.Config{WorkDir: t.TempDir()},
			}
			if err := h.WriteClientConfig(tt.args.client, tt.args.server); (err != nil) != tt.wantErr {
				t.Errorf("handler.WriteClientConfig() error = %v, wantErr %v", err, tt.wantErr)
			}
			got, want, err := readFiles(filepath.Join(h.getClientSecretDir(), fmt.Sprintf("%s.conf", tt.args.client.Name)), filepath.Join("testdata", "clients", "secrets", fmt.Sprintf("%s.conf", tt.args.client.Name)))
			if err != nil {
				t.Fatalf("readFile() error : %v", err)
			}
			if diff := cmp.Diff(want, got); diff != "" {
				t.Error(diff)
			}
		})
	}
}
