package util

import (
	"fmt"
	"maps"
	"slices"

	"github.com/aburg/kk/config"
)

func ListShortcuts(cfg *config.Config) {
	keys := slices.Collect[string](maps.Keys(cfg.Shortcuts))

	slices.Sort(keys)

	for _, key := range keys {
		fmt.Printf("%v: %v\n", key, cfg.Shortcuts[key].Description)
	}
}
