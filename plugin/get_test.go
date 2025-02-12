package plugin

import (
	"testing"
)

func TestGet(t *testing.T) {
	tests := []struct {
		name string
		url  string
		err  bool
	}{
		{
			name: "plugin:local file",
			url:  "file::../examples/plugin/inputer/json",
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
			url:      "file://foo/bar/inputer_binary",
			expected: "bar",
		},
		{
			name:     "local file without prefix",
			url:      "./foo/bar/inputer_binary",
			expected: "bar",
		},
		{
			name:     "s3 bucket",
			url:      "bucket.s3.amazonaws.com/foo",
			expected: "bucket.s3.amazonaws.com/foo",
		},
		{
			name:     "git repo ssh addresses",
			url:      "git::ssh://git@example.com/foo/bar",
			expected: "example.com/foo/bar",
		},
		{
			name:     "git scp style",
			url:      "git@github.com:naivary/Bachelorarbeit.git",
			expected: "example.com/foo/bar",
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
			t.Log(dir)
		})
	}
}
