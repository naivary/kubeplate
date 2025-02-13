package plugin

import (
	"fmt"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/hashicorp/go-getter"
)

func Get(url string) (string, error) {
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
	urlDir, err := detectDir(url)
	if err != nil {
		return "", err
	}
	dst := filepath.Join(home, ".kubeplate", urlDir)
	err = getter.GetAny(dst, url, opt)
	return dst, err
}

func detectDir(getterURL string) (string, error) {
	pwd, err := os.Getwd()
	if err != nil {
		return "", err
	}
	detected, err := getter.Detect(getterURL, pwd, getter.Detectors)
	if err != nil {
		return "", err
	}
	u, err := url.Parse(replacePrefix(detected, pwd))
	if err != nil {
		return "", err
	}
	return url.JoinPath(u.Host, removeExt(u.Path))
}

func replacePrefix(url, pwd string) string {
	const https = "https://"
	absolutePath := fmt.Sprintf("file://%s", pwd)
	switch {
	case strings.HasPrefix(url, absolutePath):
		return path.Dir(strings.Replace(url, absolutePath, https, 1))
	case strings.HasPrefix(url, "file://"):
		return path.Dir(strings.Replace(url, "file://", https, 1))
	case strings.HasPrefix(url, "s3::https://"):
		return strings.Replace(url, "s3::https://", https, 1)
	case strings.HasPrefix(url, "git::ssh://git@"):
		return strings.Replace(url, "git::ssh://git@", https, 1)
	case strings.HasPrefix(url, "git::http://"):
		return strings.Replace(url, "git::http://", https, 1)
	default:
		return ""
	}
}

func removeExt(path string) string {
	ext := filepath.Ext(path)
	if ext == "" {
		return path
	}
	return strings.TrimSuffix(path, ext)
}
