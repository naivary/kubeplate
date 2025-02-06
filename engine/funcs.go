package engine

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/hashicorp/go-getter"
)

func get(url string, kind string) error {
	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	pwd, err := os.Getwd()
	if err != nil {
		return err
	}
	dst := filepath.Join(home, ".kubeplate", kind)
	fmt.Println(pwd)
	return getter.GetAny(dst, url, getter.ClientOption(func(c *getter.Client) error {
		c.Pwd = pwd
		return nil
	}))
}

func prefixFuncs(prefix string, funcs template.FuncMap) template.FuncMap {
	prefixedFuncs := template.FuncMap{}
	for name, fn := range funcs {
		if strings.HasPrefix(name, prefix) {
			continue
		}
		prefixedName := fmt.Sprintf("%s_%s", prefix, name)
		prefixedFuncs[prefixedName] = fn
	}
	return prefixedFuncs
}
