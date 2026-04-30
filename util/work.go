package util

import "github.com/aburg/kk/config"

func Work(key string, args []string, cfg *config.Config) error {
	shortcut, err := GetShortcut(cfg, key)
	if err != nil {
		return err
	}
	err = ExecuteShortcut(key, args, shortcut)
	if err != nil {
		return err
	}

	if len(shortcut.Others) > 0 {
		for _, otherkey := range shortcut.Others {
			err = Work(otherkey, args, cfg)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
