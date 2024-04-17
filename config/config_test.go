package config

import (
	"reflect"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	cfg := Config{
		WorkDir: "/etc/wireguard",
		Server: ServerConfig{
			Address:        "192.168.227.1",
			Endpoint:       "192.168.227.1:51820",
			Port:           51820,
			DNS:            "10.2.0.8",
			MTU:            1420,
			PrivateKeyFile: "/etc/wireguard/.serverkey",
			PublicKey:      "publickey",
			AllowedIPs:     "192.168.227.0/22",
			PostUp:         "iptables -A FORWARD -i %i -j ACCEPT; iptables -t nat -A POSTROUTING -o eth0 -j MASQUERADE",
			PostDown:       "iptables -D FORWARD -i %i -j ACCEPT; iptables -t nat -D POSTROUTING -o eth0 -j MASQUERADE",
		},
	}
	tests := []struct {
		name       string
		configFile string
		want       Config
		isEqual    bool
	}{
		{name: "success", configFile: "./testdata/config.match.yaml", want: cfg, isEqual: true},
		{name: "error", configFile: "./testdata/config.unmatch.yaml", want: cfg, isEqual: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			LoadConfig(tt.configFile)
			got := GetConfig()
			if reflect.DeepEqual(got, tt.want) != tt.isEqual {
				t.Errorf("LoadConfig() = %v, want %v, isEqual %v", got, tt.want, tt.isEqual)
			}
		})
	}
}
