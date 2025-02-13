package plugin

import (
	"os"
	"path"
	"runtime"
	"testing"
)

func TestGet(t *testing.T) {
	_, filename, _, _ := runtime.Caller(0)
	dir := path.Join(path.Dir(filename), "..")
	err := os.Chdir(dir)
	if err != nil {
		t.Error(err)
	}
	tests := []struct {
		name string
		url  string
		err  bool
	}{
		{
			name: "plugin:local file",
			url:  "./examples/plugin/inputer/json",
		},
		{
			name: "plugin:git repo",
			url:  "git::http://github.com/hashicorp/go-getter.git?ref=v1.7.8",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			path, err := Get(tc.url)
			if err != nil && !tc.err {
				t.Error(err)
			}
			t.Logf("path: %s\n", path)
		})
	}
}

func TestDetectDir(t *testing.T) {
	tests := []struct {
		name     string
		url      string
		expected string
	}{
		{
			name:     "local file",
			url:      "foo/bar/inputer_binary",
			expected: "foo/bar",
		},
		{
			name:     "local file without prefix",
			url:      "./foo/bar/relative/inputer_binary",
			expected: "foo/bar/relative",
		},
		{
			name:     "s3 bucket",
			url:      "bucket.s3.amazonaws.com/foo",
			expected: "s3.amazonaws.com/bucket/foo",
		},
		{
			name:     "git repo ssh addresses",
			url:      "git::ssh://git@example.com/foo/bar",
			expected: "example.com/foo/bar",
		},
		{
			name:     "git scp style",
			url:      "git@github.com:naivary/Bachelorarbeit.git",
			expected: "github.com/naivary/Bachelorarbeit",
		},
		{
			name:     "git http",
			url:      "git::http://github.com/mitchellh/vagrant.git",
			expected: "github.com/mitchellh/vagrant",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			dir, err := detectDir(tc.url)
			if err != nil {
				t.Error(err)
			}
			if dir != tc.expected {
				t.Fatalf("expected: %s; got: %s\n", tc.expected, dir)
			}
		})
	}
}
