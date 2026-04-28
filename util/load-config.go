package util

import (
	"fmt"

	"github.com/aburg/kk/config"
	"github.com/spf13/viper"
)

func LoadConfig() (config.Config, error) {
	var c config.Config

	err := viper.Unmarshal(&c)
	if err != nil {
		return config.Config{}, fmt.Errorf("unable to decode into sctruct, %v", err)
	}
	return c, nil
}
