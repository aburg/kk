package util

import (
	"fmt"

	"github.com/aburg/kk/config"
	"github.com/spf13/viper"
)

func LoadConfig() (config.Config, error) {
	var cfg config.Config

	err := viper.Unmarshal(&cfg)
	if err != nil {
		return config.Config{}, fmt.Errorf("unable to decode into sctruct, %v", err)
	}
	return cfg, nil
}
