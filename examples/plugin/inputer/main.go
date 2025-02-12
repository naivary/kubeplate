package main

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/hashicorp/go-getter"
	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/go-plugin"
	v1 "github.com/naivary/kubeplate/api/inputer/v1"
	"github.com/naivary/kubeplate/sdk/inputer"
	"google.golang.org/protobuf/types/known/structpb"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	logger := hclog.New(&hclog.LoggerOptions{
		Name:       "json_inputer",
		JSONFormat: true,
		Output:     os.Stderr,
	})
	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: inputer.Handshake,
		Plugins: map[string]plugin.Plugin{
			"inputer": &inputer.GRPCPlugin{Impl: &jsonInputer{l: logger}},
		},
		Logger:     logger,
		GRPCServer: plugin.DefaultGRPCServer,
	})
	return nil
}

var _ inputer.Inputer = (*jsonInputer)(nil)

type jsonInputer struct {
	l hclog.Logger
}

func (j *jsonInputer) Read(ctx context.Context, req *v1.ReadRequest) (*v1.ReadResponse, error) {
	const defaultTempDir = ""
	tmpDir, err := os.MkdirTemp(defaultTempDir, "getter")
	if err != nil {
		return nil, err
	}
	err = getter.GetAny(tmpDir, req.Url, getter.ClientOption(func(c *getter.Client) error {
		pwd, err := os.Getwd()
		if err != nil {
			return err
		}
		c.Pwd = pwd
		return nil
	}))
	if err != nil {
		return nil, err
	}
	path, _ := strings.CutPrefix(req.Url, "file::")
	filename := filepath.Base(path)
	data, err := os.ReadFile(filepath.Join(tmpDir, filename))
	if err != nil {
		return nil, err
	}
	anymap := make(map[string]any)
	if err := json.Unmarshal(data, &anymap); err != nil {
		return nil, err
	}
	str, err := structpb.NewValue(anymap)
	if err != nil {
		return nil, err
	}
	return &v1.ReadResponse{
		Data: map[string]*structpb.Struct{
			filename: str.GetStructValue(),
		},
	}, nil
}
