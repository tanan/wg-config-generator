package config

import (
	"reflect"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	cfg := Config{
		WorkDir: "/etc/wireguard",
		Server: ServerConfig{
			Address:  "10.0.0.1",
			Endpoint: "10.0.0.1:51820",
			Port:     51820,
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
