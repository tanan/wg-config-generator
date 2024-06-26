package config

import (
	"log/slog"
	"os"
	"sync"

	"github.com/spf13/viper"
)

var cfg Config

type Config struct {
	WorkDir string       `mapstructure:"work_dir"`
	Server  ServerConfig `mapstructure:"server"`
}

type ServerConfig struct {
	Address        string `mapstructure:"address"`
	Endpoint       string `mapstructure:"endpoint"`
	Port           int    `mapstructure:"port"`
	DNS            string `mapstructure:"dns"`
	MTU            int    `mapstructure:"mtu"`
	PrivateKeyFile string `mapstructure:"privatekeyfile"`
	PublicKey      string `mapstructure:"publickey"`
	AllowedIPs     string `mapstructure:"allowedips"`
	PostUp         string `mapstructure:"postup"`
	PostDown       string `mapstructure:"postdown"`
}

func GetConfig() Config {
	return cfg
}

func LoadConfig(configFile string) {
	var once sync.Once
	once.Do(func() {
		viper.SetConfigFile(configFile)
		viper.SetConfigType("yaml")
		viper.AutomaticEnv()
		if err := viper.ReadInConfig(); err != nil {
			slog.Error("failed to read config file", slog.String("config", configFile), slog.String("error", err.Error()))
			os.Exit(1)
		}

		if err := viper.Unmarshal(&cfg); err != nil {
			slog.Error("failed to unmarshal config", slog.String("config", configFile), slog.String("error", err.Error()))
			os.Exit(1)
		}
	})
}
