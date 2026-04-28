package config

type Shortcut struct {
	Description string   `yaml:"description"`
	Command     string   `yaml:"command"`
	Chdir       string   `yaml:"chdir"`
	Arguments   []string `yaml:"arguments"`
}
