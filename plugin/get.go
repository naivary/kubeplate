package plugin

import (
	"errors"
	"io/fs"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"github.com/hashicorp/go-getter"
)

func Get(url string, force bool) (string, bool, error) {
	const local = ""
	home, err := os.UserHomeDir()
	if err != nil {
		return "", false, err
	}
	pwd, err := os.Getwd()
	if err != nil {
		return "", false, err
	}
	opt := func(c *getter.Client) error {
		c.Pwd = pwd
		return nil
	}
	urlDir, err := dstDir(url)
	if err != nil {
		return "", false, err
	}
	dst := filepath.Join(home, ".kubeplate", urlDir)
	isExisting, err := isPluginExisting(dst)
	if err != nil {
		return "", false, err
	}
	if isExisting && !force && urlDir != local {
		return dst, false, nil
	}
	err = getter.GetAny(dst, url, opt)
	if err != nil {
		return "", false, err
	}
	return dst, true, nil
}

func isPluginExisting(dst string) (bool, error) {
	_, err := os.Stat(dst)
	if err == nil {
		return true, nil
	}
	if errors.Is(err, fs.ErrNotExist) {
		return false, nil
	}
	return false, err
}

func dstDir(getterURL string) (string, error) {
	pwd, err := os.Getwd()
	if err != nil {
		return "", err
	}
	detected, err := getter.Detect(getterURL, pwd, getter.Detectors)
	if err != nil {
		return "", err
	}
	httpsPrefixed := replacePrefix(detected)
	noSubDir := removeSubDir(httpsPrefixed)
	u, err := url.Parse(noSubDir)
	if err != nil {
		return "", err
	}
	version := u.Query().Get("ref")
	return url.JoinPath(u.Host, removeExt(u.Path), version)
}

func replacePrefix(url string) string {
	const https = "https://"
	const local = "local"
	switch {
	case strings.HasPrefix(url, "file://"):
		return ""
	case strings.HasPrefix(url, "s3::https://"):
		return strings.Replace(url, "s3::https://", https, 1)
	case strings.HasPrefix(url, "git::ssh://git@"):
		return strings.Replace(url, "git::ssh://git@", https, 1)
	case strings.HasPrefix(url, "git::http://"):
		return strings.Replace(url, "git::http://", https, 1)
	case strings.HasPrefix(url, "git::https://"):
		return strings.Replace(url, "git::https://", https, 1)
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

func removeSubDir(url string) string {
	splitted := strings.Split(url, "//")
	if len(splitted) < 3 {
		return url
	}
	splitted = splitted[:len(splitted)-1]
	return strings.Join(splitted, "//")
}
