package engine

import (
	"os"

	"github.com/goccy/go-yaml"
)

type FuncConfig struct {
	Name    string `yaml:"name"`
	URL     string `yaml:"url"`
	Version string `yaml:"version"`
}

type Config struct {
	Funcs []FuncConfig `yaml:"funcs"`
}

func newConfig(path string) (*Config, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	config := &Config{}
	if err := yaml.Unmarshal(content, config); err != nil {
		return nil, err
	}
	return config, nil
}
