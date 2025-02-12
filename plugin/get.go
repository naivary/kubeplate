package plugin

import (
	"os"
	"path/filepath"

	"github.com/hashicorp/go-getter"
)

func Get(url string, k Kind) (string, error) {
	// TODO: check if URL already exists localy or not
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	pwd, err := os.Getwd()
	if err != nil {
		return "", err
	}
	opt := func(c *getter.Client) error {
		c.Pwd = pwd
		return nil
	}
	dst := filepath.Join(home, ".kubeplate", k.String())
	err = getter.GetAny(dst, url, opt)
	return dst, err
}

func detectDir(url string) (string, error) {
	pwd, err := os.Getwd()
	if err != nil {
		return "", err
	}
	detected, err := getter.Detect(url, pwd, getter.Detectors)
	if err != nil {
		return "", err
	}
	return detected, nil
}
