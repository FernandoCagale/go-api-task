package config

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/spf13/viper"
)

type Config struct {
	Port         int    `envconfig:"PORT"`
	DatastoreURL string `envconfig:"DATASTORE_URL"`
}

func init() {
	viper.SetDefault("port", "3000")
	viper.SetDefault("datastoreURL", "postgresql://test:test@api-postgres:5432/test?sslmode=disable")
}

func LoadEnv() (*Config, error) {
	var instance Config
	if err := viper.Unmarshal(&instance); err != nil {
		return nil, err
	}

	err := envconfig.Process("", &instance)
	if err != nil {
		return &instance, err
	}

	return &instance, nil
}
