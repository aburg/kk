package util

import (
	"fmt"

	"github.com/aburg/kk/config"
)

func ListShortcuts(cfg *config.Config) {
	for key, shortcut := range cfg.Shortcuts {
		fmt.Printf("%v: %v\n", key, shortcut.Description)
	}
}
