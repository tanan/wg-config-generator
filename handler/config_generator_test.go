package handler

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/tanan/wg-config-generator/config"
	"github.com/tanan/wg-config-generator/model"
)

func Test_handler_CreateClientConfig(t *testing.T) {
	mockCommand := &MockCommand{
		CreatePrivateKeyFunc: func() (string, error) {
			return "privatekey", nil
		},
		CreatePreSharedKeyFunc: func() (string, error) {
			return "presharedkey", nil
		},
		CreatePublicKeyFunc: func(privateKey string) (string, error) {
			if privateKey == "privatekey" {
				return "publickey", nil
			}
			return "", fmt.Errorf("privatekey %s does not match.", privateKey)
		},
	}

	clientConfig := model.ClientConfig{
		Name:         "client1",
		Address:      "10.10.10.10",
		PrivateKey:   "privatekey",
		PublicKey:    "publickey",
		PresharedKey: "presharedkey",
	}

	type args struct {
		name    string
		address string
	}
	tests := []struct {
		name    string
		args    args
		want    model.ClientConfig
		wantErr bool
	}{
		{name: "success", args: args{name: "client1", address: "10.10.10.10"}, want: clientConfig, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := handler{
				Command: mockCommand,
			}
			got, err := h.CreateClientConfig(tt.args.name, tt.args.address)
			if (err != nil) != tt.wantErr {
				t.Errorf("handler.CreateClientConfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("handler.CreateClientConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_handler_CreateServerConfig(t *testing.T) {
	type fields struct {
		Command Command
		Config  config.Config
	}
	tests := []struct {
		name    string
		fields  fields
		want    model.ServerConfig
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
			got, err := h.CreateServerConfig()
			if (err != nil) != tt.wantErr {
				t.Errorf("handler.CreateServerConfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("handler.CreateServerConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}
