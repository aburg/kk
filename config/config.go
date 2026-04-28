package config

type Config struct {
	Shortcuts map[string]Shortcut `yaml:"shortcuts"`
}
