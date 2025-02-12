package plugin

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/hashicorp/go-getter"
)

func Get(url string) (string, error) {
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
	dst := filepath.Join(home, ".kubeplate")
	err = getter.GetAny(dst, url, opt)
	return dst, err
}

func detectDir(url string) (string, error) {
	// TODO: Bring every possible getter path to a http(s):// path 
	// just to use url.Parse and get the url.Path to use as a directory
	// structure
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

func cutPrefix(url string) (string, bool) {
	switch {
	case strings.HasPrefix(url, "https://") || strings.HasPrefix(url, "http://"):
		return url, true
	case strings.HasPrefix(url, "file://"):
		return strings.CutPrefix(url, "file://")
	case strings.HasPrefix(url, "git://"):
		return strings.CutPrefix(url, "git://")
	case strings.HasPrefix(url, "s3://"):
		return strings.CutPrefix(url, "file://")
	default:
		return "UKNOWN", false
	}
}
