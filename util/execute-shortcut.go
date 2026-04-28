package util

import (
	"fmt"
	"os"

	"github.com/aburg/kk/aishit"
	"github.com/aburg/kk/config"
)

func ExecuteShortcut(key string, shortcut config.Shortcut, args []string) error {
	if shortcut.Chdir != "" {
		dir, err := aishit.ExpandHome(shortcut.Chdir)
		if err != nil {
			return fmt.Errorf("could not expand chdir: %v", err)
		}
		err = os.Chdir(dir)
		if err != nil {
			return fmt.Errorf("chdir failed: %v", err)
		}
		fmt.Println("chdir:", dir)
	}
	if shortcut.Command == "" {
		return fmt.Errorf("the shortcut \"%v\" is missing a command", key)
	}
	aishit.Execute(shortcut.Command, args)
	return nil
}
