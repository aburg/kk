package aishit

import (
	"os"
	"path/filepath"
	"strings"
)

func ExpandHome(path string) (string, error) {
	if !strings.HasPrefix(path, "~/") {
		return path, nil
	}

	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	// Join home directory with everything after the "~/"
	return filepath.Join(home, path[2:]), nil
}
