package util

import (
	"fmt"
	"maps"
	"os"
	"slices"
	"text/tabwriter"

	"github.com/aburg/kk/config"
)

func ListShortcuts(cfg *config.Config) {
	keys := slices.Collect[string](maps.Keys(cfg.Shortcuts))

	slices.Sort(keys)

	w := tabwriter.NewWriter(os.Stdout, 1, 1, 3, ' ', 0)
	for _, key := range keys {
		fmt.Fprintf(w, "%v\t%v\n", key, cfg.Shortcuts[key].Description)
	}
	w.Flush()
}
