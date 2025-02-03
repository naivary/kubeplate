package funcs

import (
	"encoding/json"
	"os"
	"testing"

	"google.golang.org/protobuf/types/known/structpb"
)

func getTestJSONData(path string) (map[string]*structpb.Struct, error) {
	data := make(map[string]*structpb.Struct)
	content, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	anymap := make(map[string]any)
	if err := json.Unmarshal(content, &anymap); err != nil {
		return nil, err
	}
	str, err := structpb.NewStruct(anymap)
	if err != nil {
		return nil, err
	}
	data["test.json"] = str
	return data, nil
}

func TestGet(t *testing.T) {
	tests := []struct {
		name          string
		accessPath    string
		expectedValue any
		isError       bool
	}{
		{
			name:          "single string",
			accessPath:    "name",
			expectedValue: "kubeplate",
		},
		{
			name:          "does not exist",
			accessPath:    "doest-not-exist",
			expectedValue: "",
			isError:       true,
		},
		{
			name:          "key of nested object in array",
			accessPath:    "arr.0.key",
			expectedValue: "value",
		},
	}

	data, err := getTestJSONData("test.json")
	if err != nil {
		t.Error(err)
	}

	get := Get(data)

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			val, err := get("test.json", tc.accessPath)
			if err != nil && !tc.isError {
				t.Error(err)
			}

			if err != nil && tc.isError {
				t.Skip("error occured but was expected. Skipping...")
			}
			if val != tc.expectedValue {
				t.Errorf("expected: %v; got: %v\n", tc.expectedValue, val)
			}
		})
	}
}
