package plugin

import (
	"testing"
)

func TestGet(t *testing.T) {
	tests := []struct {
		name string
		url  string
		kind Kind
		err  bool
	}{
		{
			name: "plugin:local file",
			url:  "file::../examples/plugin/inputer/json",
			kind: Inputer,
		},
		{
			name: "plugin:git repo",
			url:  "http://github.com/hashicorp/go-getter",
			kind: Inputer,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			path, err := Get(tc.url, tc.kind)
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
			url:      "file::./foo/bar/inputer_binary",
			expected: "bar",
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
