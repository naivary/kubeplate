package kubeplate

import "testing"

func TestLoadVars(t *testing.T) {
	tests := []struct {
		name        string
		inputerPath string
		varsPath    string
	}{
		{
			name:        "json inputer",
			inputerPath: "json",
			varsPath:    "./examples/vars/vars.json",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			_, err := LoadVars(tc.inputerPath, tc.varsPath)
			if err != nil {
				t.Error(err)
			}
		})
	}
}
