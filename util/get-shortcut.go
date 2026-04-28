package util

import (
	"fmt"

	"github.com/aburg/kk/config"
)

func GetShortcut(cfg *config.Config, key string) (config.Shortcut, error) {
	shortcut, ok := cfg.Shortcuts[key]
	if !ok {
		return shortcut, fmt.Errorf("shortcut \"%v\" not found", key)
	}
	return shortcut, nil
}
