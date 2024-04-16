package handler

import "github.com/tanan/wg-config-generator/domain"

func (h handler) WriteServerConfig(domain.ServerConfig) error {
	return nil
}

func (h handler) WriteClientConfig(domain.ClientConfig) error {
	return nil
}

func (h handler) SendClientConfigByEmail(domain.ClientConfig) error {
	return nil
}
