package config

type Config struct {
	Model   []Model `yaml:"model"`
	Context string  `yaml:"context"`
}

type Model struct {
	Name string `yaml:"name"`
	Path string `yaml:"path"`
	Key  string `yaml:"key"`
	Type string `yaml:"type"`
}
