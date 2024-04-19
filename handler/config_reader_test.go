package handler

import (
	"path/filepath"
	"reflect"
	"testing"

	"github.com/tanan/wg-config-generator/config"
	"github.com/tanan/wg-config-generator/model"
)

func Test_handler_GetClientList(t *testing.T) {
	tests := []struct {
		name    string
		want    []model.ClientConfig
		wantErr bool
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := handler{
				Config: config.Config{
					WorkDir: "testdata",
				},
			}
			got, err := h.GetClientList()
			if (err != nil) != tt.wantErr {
				t.Errorf("handler.GetClientList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("handler.GetClientList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_handler_readClient(t *testing.T) {

	client1 := model.ClientConfig{
		Name:         "client1",
		Address:      "10.10.10.10",
		PublicKey:    "publickey",
		PresharedKey: "presharedkey",
	}

	tests := []struct {
		name    string
		want    model.ClientConfig
		wantErr bool
	}{
		{name: "client1.json", want: client1, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := handler{
				Config: config.Config{
					WorkDir: "testdata",
				},
			}
			got, err := h.readClient(filepath.Join(h.getClientDir(), tt.name))
			if (err != nil) != tt.wantErr {
				t.Errorf("handler.readClient() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("handler.readClient() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_handler_readPrivateKey(t *testing.T) {
	tests := []struct {
		name    string
		fn      string
		want    string
		wantErr bool
	}{
		{name: "dummy.privatekey", fn: "dummy.privatekey", want: "privatekey", wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := handler{
				Config: config.Config{
					WorkDir: "testdata",
				},
			}
			got, err := h.readPrivateKey(filepath.Join("testdata", tt.fn))
			if (err != nil) != tt.wantErr {
				t.Errorf("handler.readPrivateKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("handler.readPrivateKey() = %v, want %v", got, tt.want)
			}
		})
	}
}
