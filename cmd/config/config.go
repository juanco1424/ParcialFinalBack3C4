package config

import (
	"errors"
	"os"
)

type Config struct {
	PublicConfig  PublicConfig
	PrivateConfig PrivateConfig
}

type PublicConfig struct {
	PublicKey string
}

type PrivateConfig struct {
	SecretKey string
}

var (
	envs = map[string]PublicConfig{
		"local": {PublicKey: "dh092023g4"},
	}
)

func NewConfig(env string) (Config, error) {

	publicConfig, exists := envs[env]
	if !exists {
		return Config{}, errors.New("env no existe")
	}

	secretKey := os.Getenv("SECRET_KEY")
	if secretKey == "" {
		return Config{}, errors.New("secret key no existe en env")
	}

	return Config{
		PublicConfig: publicConfig,
		PrivateConfig: PrivateConfig{
			SecretKey: secretKey,
		},
	}, nil

}
