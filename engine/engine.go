package engine

import (
	"text/template"

	"github.com/naivary/kubeplate/sdk/outputer"
)

var _ Engine = (*engine)(nil)

func New(configFile string) (Engine, error) {
	cfg, err := newConfig(configFile)
	if err != nil {
		return nil, err
	}
	e := &engine{
		funcs:  &template.FuncMap{},
		config: cfg,
	}
	return e, nil
}

type engine struct {
	funcs *template.FuncMap

	config *Config
}

func (e *engine) AddPlugin(path string) error {
	return get(path, "funcs")
}

func (e *engine) Execute(out outputer.Outputer, data any) error {
	return nil
}
