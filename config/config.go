package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Server ServerConfig `yaml:"server"`
	Client ClientConfig `yaml:"client"`
	Secret SecretConfig `yaml:"secret"`
}

type ServerConfig struct {
	Endpoint           string             `yaml:"endpoint"`
	Port               int                `yaml:"port"`
	WireguardInterface WireguardInterface `yaml:"interface"`
}

type WireguardInterface struct {
	Name    string `yaml:"name"`
	Address string `yaml:"address"`
	Network string `yaml:"network"`
}

type ClientConfig struct {
	Path string `yaml:"path"`
}

type SecretConfig struct {
	Path string `yaml:"path"`
}

func LoadConfig(path string) (Config, error) {
	var c Config
	b, err := os.ReadFile(path)
	if err != nil {
		return Config{}, err
	}
	err = yaml.Unmarshal(b, &c)
	if err != nil {
		return Config{}, err
	}
	return c, nil
}
