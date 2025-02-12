package engine

import (
	"fmt"
	"os"
	"path/filepath"
	"plugin"
	"strings"
	"text/template"

	"github.com/hashicorp/go-getter"
	"github.com/naivary/kubeplate/sdk/outputer"
)

const kubeplate = "kubeplate"

type Engine interface {
	LoadFuncs(url string) error

	Execute(out outputer.Outputer, data any) error
}

var _ Engine = (*engine)(nil)

func New(configFile string) (Engine, error) {
	cfg, err := newConfig(configFile)
	if err != nil {
		return nil, err
	}

	home, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	e := &engine{
		funcs:    template.FuncMap{},
		config:   cfg,
		funcsDir: filepath.Join(home, kubeplate, "funcs"),
	}
	return e, nil
}

type engine struct {
	funcs template.FuncMap

	funcsDir string

	config *Config
}

func (e *engine) LoadFuncs(url string) error {
	const prefix = "file::"
	pwd, err := os.Getwd()
	if err != nil {
		return err
	}
	err = getter.GetAny(e.funcsDir, url, func(c *getter.Client) error {
		c.Pwd = pwd
		return nil
	})
	if err != nil {
		return err
	}
	path, _ := strings.CutPrefix(url, prefix)
	filename := filepath.Base(path)
	pluginPath := filepath.Join(e.funcsDir, filename)
	pl, err := plugin.Open(pluginPath)
	if err != nil {
		return err
	}
	symbol, err := pl.Lookup("Funcs")
	if err != nil {
		return err
	}
	funcs, isTemplateMap := symbol.(*template.FuncMap)
	if !isTemplateMap {
		return fmt.Errorf("`%s` does not contain an exported variable `Funcs`", pluginPath)
	}
	// TODO: name fehlt für prefixing
	return nil
}

func (e *engine) Execute(out outputer.Outputer, data any) error {
	return nil
}
