package engine

import (
	"fmt"
	"os"

	"github.com/goccy/go-yaml"
	"github.com/naivary/kubeplate/sdk/outputer"
)

type FuncConfig struct {
	Name    string `yaml:"name"`
	URL     string `yaml:"url"`
	Version string `yaml:"version"`
}

type Config struct {
	Funcs []FuncConfig `yaml:"funcs"`
}

type Engine interface {
	AddPlugin(path string) error

	Execute(out outputer.Outputer, data any) error
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
	fmt.Println(config)
	return config, nil
}
