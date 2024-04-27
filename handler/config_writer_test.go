package handler

import (
	"bytes"
	"fmt"
	"path/filepath"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/tanan/wg-config-generator/model"
)

func Test_handler_writeServerConfig(t *testing.T) {
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
			// prepare testdata
			testFile := filepath.Join("testdata", tt.wantFileName)
			want, err := readFile(testFile)
			if err != nil {
				t.Errorf("failed to prepare testdata: %s, error: %v", testFile, err)
				return
			}

			// run test
			h := handler{}
			w := &bytes.Buffer{}
			if err := h.writeServerConfig(w, tt.args.server, tt.args.peers); (err != nil) != tt.wantErr {
				t.Errorf("handler.writeServerConfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			// compare
			got := w.String()
			if diff := cmp.Diff(want, got); diff != "" {
				t.Error(diff)
			}
		})
	}
}

func Test_handler_writeClientConfig(t *testing.T) {
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
			// prepare testdata
			testFile := filepath.Join("testdata", "clients", "secrets", fmt.Sprintf("%s.conf", tt.args.client.Name))
			want, err := readFile(testFile)
			if err != nil {
				t.Errorf("failed to prepare testdata: %s, error: %v", testFile, err)
				return
			}

			// run test
			h := handler{}
			w := &bytes.Buffer{}
			if err := h.writeClientConfig(w, tt.args.client, tt.args.server); (err != nil) != tt.wantErr {
				t.Errorf("handler.writeClientConfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			// compare
			got := w.String()
			if diff := cmp.Diff(want, got); diff != "" {
				t.Error(diff)
			}
		})
	}
}

func Test_handler_saveClientSetting(t *testing.T) {
	cc := model.ClientConfig{
		Name:         "client1",
		Address:      "10.10.10.10",
		PrivateKey:   "privatekey",
		PublicKey:    "publickey",
		PresharedKey: "presharedkey",
	}
	tests := []struct {
		name    string
		cc      model.ClientConfig
		wantErr bool
	}{
		{name: "ok", cc: cc, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// prepare testdata
			testFile := filepath.Join("testdata", "clients", fmt.Sprintf("%s.json", tt.cc.Name))
			want, err := readFile(testFile)
			if err != nil {
				t.Errorf("failed to prepare testdata: %s, error: %v", testFile, err)
				return
			}
			h := handler{}
			w := &bytes.Buffer{}
			if err := h.saveClientSetting(w, tt.cc); (err != nil) != tt.wantErr {
				t.Errorf("handler.saveClientSetting() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			// compare
			got := w.String()
			if diff := cmp.Diff(want, got); diff != "" {
				t.Error(diff)
			}
		})
	}
}
